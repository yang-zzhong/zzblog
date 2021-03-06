package zzblog

import (
	"encoding/json"
	"errors"
	"github.com/yang-zzhong/md"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

/*

 dir
 data/ ---
 		|
		|----- blogs/ ----+----- hello-world.md
		|				  |----- my-second.md
		                  |----- my-tset-blog.md

*/
func ensureDirExist(dir string) error {
	stat, err := os.Stat(dir)
	if err == nil && stat.IsDir() {
		return nil
	}
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		return nil
	}
	return err
}

func TraversingDir(dirpath string, handle func(pathfile string)) error {
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return err
	}
	for _, file := range dir {
		pathfile := path.Join(dirpath, file.Name())
		if !file.IsDir() {
			handle(pathfile)
			continue
		}
		if err := TraversingDir(pathfile, handle); err != nil {
			return err
		}
	}
	return nil
}

type StringWithLang struct {
	Lang string
	Val  string
}

// Zzblog
type FileZzblog struct {
	root   string
	meta   map[string]*LangGroup
	cates  []StringWithLang
	langs  []string
	tags   []StringWithLang
	images map[string]*Image
}

func NewFileZzblog(root string) *FileZzblog {
	zz := new(FileZzblog)
	zz.root = root
	return zz
}

func (zz *FileZzblog) Init() error {
	zz.meta = make(map[string]*LangGroup)
	zz.cates = []StringWithLang{}
	zz.tags = []StringWithLang{}
	zz.images = make(map[string]*Image)
	return zz.build()
}

func (zz *FileZzblog) build() error {
	if err := ensureDirExist(zz.blogDir()); err != nil {
		return err
	}
	err := TraversingDir(zz.blogDir(), func(file string) {
		ext := path.Ext(file)
		if ext == ".md" {
			if err := zz.addBlog(file); err != nil {
				log.Printf("add %s file error: %s", file, err.Error())
			}
		}
	})
	if err != nil {
		return err
	}
	if err := ensureDirExist(zz.imageDir()); err != nil {
		return err
	}
	return TraversingDir(zz.imageDir(), func(file string) {
		ext := path.Ext(file)
		if ext == ".png" || ext == ".bmp" || ext == ".jpeg" || ext == ".jpg" || ext == ".gif" {
			err := zz.addImage(file)
			if err != nil {
				log.Printf("add image error: %s\n", err.Error())
			}
		}
	})
}

func (zz *FileZzblog) addImage(file string) error {
	img, err := getImage(file)
	if err != nil {
		return err
	}
	if _, ok := zz.images[img.Id]; !ok {
		zz.images[img.Id] = img
	}

	return nil
}

func (zz *FileZzblog) AddImage(r ImageReader) error {
	img, err := getImageFromFile(r)
	if err != nil {
		return err
	}
	if _, ok := zz.images[img.Id]; ok {
		return nil
	}
	pathfile := path.Join(zz.imageDir(), img.Id+getImageExt(img.Format))
	f, err := os.OpenFile(pathfile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

func (zz *FileZzblog) GetImage(id string) *Image {
	img, ok := zz.getImage(id)
	if !ok {
		return nil
	}
	return img
}

func (zz *FileZzblog) GetImageByFilename(filename string) *Image {
	for _, img := range zz.images {
		pf := path.Join(zz.imageDir(), filename)
		if pf == img.Pathfile {
			return img
		}
	}
	return nil
}

func (zz *FileZzblog) getImage(id string) (img *Image, ok bool) {
	img, ok = zz.images[id]
	return
}

func (zz *FileZzblog) addBlog(file string) error {
	f, e := os.Open(file)
	if e != nil {
		return e
	}
	defer f.Close()
	blog, e := zz.AddByReader(f)
	if e != nil {
		return e
	}
	blog.SetFile(file)
	return nil
}

func (zz *FileZzblog) Has(id, lang string) bool {
	if _, ok := zz.meta[id]; ok {
		return zz.meta[id].Select(lang).Lang == lang
	}
	return false
}

func (zz *FileZzblog) Get(id string, lang string) *Blog {
	group, ok := zz.meta[id]
	if ok {
		return group.Select(lang)
	}
	return nil
}

func (zz *FileZzblog) AddByReader(r io.Reader) (blog *Blog, err error) {
	var head md.MdHead
	if head, err = md.ParseHead(r); err != nil {
		return
	}
	blog = new(Blog)
	blog.URLID = head.Urlid
	blog.Title = head.Title
	blog.Tags = head.Tags
	blog.Category = head.Cate
	blog.Overview = head.Overview
	blog.Lang = head.Lang
	blog.Image = head.Image
	blog.PublishedAt = head.PublishedAt
	blog.UpdatedAt = head.UpdatedAt
	err = zz.Add(blog)
	return
}

func (zz *FileZzblog) Add(blog *Blog) error {
	if blog.URLID == "" {
		return errors.New("has no urlid field")
	}
	zz.updateTag(blog.Lang, blog.Tags)
	zz.updateCate(blog.Lang, blog.Category)
	zz.updateLangs(blog.Lang)
	if _, ok := zz.meta[blog.URLID]; ok {
		zz.meta[blog.URLID].Add(blog)
		return nil
	}
	zz.meta[blog.URLID] = NewLangGroup("zh-CN", blog)
	return nil
}

func (zz *FileZzblog) Filter(filter func(*LangGroup) *Blog) BlogSet {
	blogs := []*Blog{}
	for _, group := range zz.meta {
		if blog := filter(group); blog != nil {
			blogs = append(blogs, blog)
		}
	}
	return NewMBlogSet(blogs)
}

func (zz *FileZzblog) blogDir() string {
	return path.Join(zz.root, "blogs")
}

func (zz *FileZzblog) imageDir() string {
	return path.Join(zz.root, "images")
}

func (zz *FileZzblog) updateCate(lang string, cate string) {
	found := false
	for _, t := range zz.cates {
		if t.Val == cate && t.Lang == lang {
			found = true
			break
		}
	}
	if !found {
		var c StringWithLang
		c.Lang = lang
		c.Val = cate
		zz.cates = append(zz.cates, c)
	}
}

func (zz *FileZzblog) updateLangs(lang string) {
	for _, l := range zz.langs {
		if l == lang {
			return
		}
	}
	zz.langs = append(zz.langs, lang)
}

func (zz *FileZzblog) updateTag(lang string, tags []string) {
	for _, tag := range tags {
		found := false
		for _, t := range zz.tags {
			if t.Val == tag && t.Lang == lang {
				found = true
				break
			}
		}
		if !found {
			var t StringWithLang
			t.Lang = lang
			t.Val = tag
			zz.tags = append(zz.tags, t)
		}
	}
}

func (zz *FileZzblog) Cates(lang string) []string {
	cates := []string{}
	for _, cate := range zz.cates {
		if lang == "" || lang != "" && cate.Lang == lang {
			cates = append(cates, cate.Val)
		}
	}
	return cates
}

func (zz *FileZzblog) Tags(lang string) []string {
	tags := []string{}
	for _, t := range zz.tags {
		if lang == "" || lang != "" && t.Lang == lang {
			tags = append(tags, t.Val)
		}
	}
	return tags
}

func (zz *FileZzblog) Author(lang string) *Author {
	var author Author
	file, err := os.Open(path.Join(zz.root, "author-"+lang+".json"))
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("%v\n", err)
			return nil
		}
		file, err = os.Open(path.Join(zz.root, "author.json"))
		if err != nil {
			log.Printf("%v\n", err)
			return nil
		}
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&author); err != nil {
		log.Printf("%v\n", err)
		return nil
	}

	return &author
}

func (zz *FileZzblog) Theme() []map[string]string {
	ret := []map[string]string{}
	file, err := os.Open(path.Join(zz.root, "theme.json"))
	if err != nil {
		return nil
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&ret); err != nil {
		log.Printf("%v\n", err)
	}

	return ret
}
