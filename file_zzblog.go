package zzblog

import (
	"os"
	"io"
	"log"
	"path"
	"io/ioutil"
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

 func TraversingDir(dirpath string, handle func (pathfile string)) error {
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

 // Zzblog
 type FileZzblog struct {
	root string
	meta map[string]*Blog
	cates []string
	tags []string
 }

 func NewFileZzblog(root string) *FileZzblog {
	zz := new(FileZzblog)
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
	zz.cates = []string{}
	zz.tags = []string{}
	return zz.build()
 }

 func (zz *FileZzblog) build() error {
	return TraversingDir(zz.root, func(file string) {
		if path.Ext(file) != ".md" {
			return
		}
		f, e := os.Open(file)
		if e != nil {
			return
		}
		if blog, e := zz.AddByReader(f); e == nil {
			blog.File = file
		}
	})
 }

 func (zz *FileZzblog) Has(id string) bool {
	_, ok := zz.meta[id]
	return ok
 }

 func (zz *FileZzblog) Get(id string) *Blog {
	 if blog, ok := zz.meta[id]; !ok {
		 return nil
	 } else {
		 return blog
	 }
 }

 func (zz *FileZzblog) AddByReader(r io.Reader) (blog *Blog, err error) {
	var p *ParsedBlog
	if p, err = ParseBlog(r); err != nil {
		return
	}
	blog = new(Blog)
	blog.URLID = p.URLID
	blog.Title = p.Title
	blog.Tags = p.Tags
	blog.Category = p.Category
	blog.Overview = p.Overview
	blog.Images = p.Images
	blog.Lang = p.Lang
	blog.Langs = p.Langs
	zz.Add(blog)
	return
 }

 func (zz *FileZzblog) Add(blog *Blog) error {
	if zz.Has(blog.URLID) {
		blog = zz.Get(blog.URLID)
	} else {
		zz.meta[blog.URLID] = blog
	}
	zz.updateTag(blog.Tags)
	zz.updateCate(blog.Category)
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

 func (zz *FileZzblog) updateCate(cate string) {
	found := false
	for _, t := range zz.cates {
		if t == cate {
			found = true
			break
		}
	}
	if !found {
		zz.cates = append(zz.cates, cate)
	}
 }

 func (zz *FileZzblog) updateTag(tags []string) {
	for _, tag := range tags {
		found := false
		for _, t := range zz.tags {
			if t == tag {
				found = true
				break
			}
		}
		if !found {
			zz.tags = append(zz.tags, tag)
		}
	}
 }