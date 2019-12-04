package zzblog

import (
	"encoding/json"
	"errors"
	"gopkg.in/djherbis/times.v1"
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
	meta   map[string]*Blog
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
	zz.meta = make(map[string]*Blog)
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
			zz.addBlog(file)
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
	info, e := f.Stat()
	if e != nil {
		return e
	}
	blog.UpdatedAt = info.ModTime()
	t, err := times.Stat(file)
	if err != nil {
		return err
	}
	if t.HasBirthTime() {
		blog.CreatedAt = t.BirthTime()
	}
	return nil
}

func (zz *FileZzblog) Has(id, lang string) bool {
	_, ok := zz.meta[zz.id(id, lang)]
	return ok
}

func (zz *FileZzblog) Get(id string, lang string) *Blog {
	blog, ok := zz.meta[zz.id(id, lang)]
	if ok {
		return blog
	}
	blog, ok = zz.meta[zz.id(id, "en")]
	if ok {
		return blog
	}
	for _, lang := range zz.langs {
		blog, ok := zz.meta[zz.id(id, lang)]
		if ok {
			return blog
		}
	}
	return nil
}

func (zz *FileZzblog) id(id, lang string) string {
	return id + "+" + lang
}

func (zz *FileZzblog) AddByReader(r io.Reader) (blog *Blog, err error) {
	var p *ParsedBlog
	if p = ParseBlog(r); p == nil {
		err = errors.New("read content error")
		return
	}
	blog = new(Blog)
	blog.URLID = p.URLID
	blog.Title = p.Title
	blog.Tags = p.Tags
	blog.Category = p.Category
	blog.Overview = p.Overview
	blog.Lang = p.Lang
	blog.Image = p.Image
	zz.Add(blog)
	return
}

func (zz *FileZzblog) Add(blog *Blog) error {
	if blog.URLID == "" {
		return errors.New("has no urlid field")
	}
	zz.meta[zz.id(blog.URLID, blog.Lang)] = blog
	zz.updateTag(blog.Lang, blog.Tags)
	zz.updateCate(blog.Lang, blog.Category)
	zz.updateLangs(blog.Lang)
	return nil
}

func (zz *FileZzblog) Filter(filter func(*Blog) bool) BlogSet {
	blogs := []*Blog{}
	for _, blog := range zz.meta {
		if filter(blog) {
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
		if cate.Lang == lang {
			cates = append(cates, cate.Val)
		}
	}
	return cates
}

func (zz *FileZzblog) Tags(lang string) []string {
	tags := []string{}
	for _, t := range zz.tags {
		if t.Lang == lang {
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
			log.Print("%v\n", err)
			return nil
		}
		file, err = os.Open(path.Join(zz.root, "author.json"))
		if err != nil {
			log.Print("%v\n", err)
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
