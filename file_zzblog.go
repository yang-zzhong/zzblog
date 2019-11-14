package zzblog

import (
	"errors"
	"io"
	"io/ioutil"
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
		if err := os.MkdirAll(dir, 0744); err != nil {
			return err
		}
		return nil
	}
	return err
}

func TraversingDir(dirpath string, handle func(pathfile string)) error {
	dir, e := ioutil.ReadDir(dirpath)
	if e != nil {
		return e
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
	tags   []StringWithLang
	images map[string]*Image
}

func NewFileZzblog(root string) *FileZzblog {
	zz := new(FileZzblog)
	zz.images = make(map[string]*Image)
	zz.root = root
	if err := ensureDirExist(zz.blogDir()); err != nil {
		panic(err)
	}
	if err := zz.Rebuild(); err != nil {
		panic(err)
	}
	return zz
}

func (zz *FileZzblog) Rebuild() error {
	zz.meta = make(map[string]*Blog)
	zz.cates = []StringWithLang{}
	zz.tags = []StringWithLang{}
	return zz.build()
}

func (zz *FileZzblog) build() error {
	return TraversingDir(zz.root, func(file string) {
		ext := path.Ext(file)
		if ext == ".md" {
			zz.addBlog(file)
		} else if ext == ".png" || ext == ".jpeg" || ext == ".jpg" || ext == ".gif" {
			zz.addImage(file)
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
	pathfile := path.Join(zz.root, "images", img.Id+getImageExt(img.Format))
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
	if e == nil {
		return e
	}
	blog.UpdatedAt = info.ModTime()
	// u := info.Sys()
	// if u != nil {
	// 	stat := u.(*syscall.Stat_t)
	// 	log.Printf("%v\n", stat.Mtim)
	// 	// blog.CreatedAt = time.Unix(int64(ts.Sec), int64(ts.Nsec))
	// }
	return nil
}

func (zz *FileZzblog) Has(id, lang string) bool {
	_, ok := zz.meta[zz.id(id, lang)]
	return ok
}

func (zz *FileZzblog) Get(id string, lang string) *Blog {
	if blog, ok := zz.meta[zz.id(id, lang)]; !ok {
		return nil
	} else {
		return blog
	}
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
	zz.Add(blog)
	return
}

func (zz *FileZzblog) Add(blog *Blog) error {
	zz.meta[zz.id(blog.URLID, blog.Lang)] = blog
	zz.updateTag(blog.Lang, blog.Tags)
	zz.updateCate(blog.Lang, blog.Category)
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
