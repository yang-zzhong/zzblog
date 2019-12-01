package zzblog

import (
	httprouter "github.com/yang-zzhong/go-httprouter"
	"log"
	"net"
	"net/http"
)

type ZzblogHttp struct {
	router *httprouter.Router
	zz     Zzblog
}

func NewHttp(root string, docroot string) *ZzblogHttp {
	zz := new(ZzblogHttp)
	zz.initRouter(docroot)
	zz.zz = NewFileZzblog(root)
	if err := zz.zz.Init(); err != nil {
		panic(err)
	}
	return zz
}

func (h *ZzblogHttp) initRouter(docroot string) {
	h.router = httprouter.NewRouter()
	h.router.DocRoot = docroot
	if h.router.DocRoot == "" {
		h.router.Tries = []string{httprouter.Api}
	}
}

func (h *ZzblogHttp) registerGetBlogs() {
	h.router.OnGet("/blogs", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		lang := r.FormValue("lang")
		page := 1
		pageSize := 10
		tag := r.FormValue("tag")
		cate := r.FormValue("cate")
		if r.FormInt("page") != 0 {
			page = int(r.FormInt("page"))
		}
		if r.FormInt("page_size") != 0 {
			pageSize = int(r.FormInt("page_size"))
		}
		blogs := h.zz.Filter(func(b *Blog) bool {
			tm := true
			tc := true
			lc := true
			if tag != "" {
				tm = false
				for _, t := range b.Tags {
					if tag == t {
						tm = true
						break
					}
				}
			}
			if cate != "" {
				tc = cate == b.Category
			}
			if lang != "" {
				lc = lang == b.Lang
			}
			return tm && tc && lc
		}).Page(page, pageSize).Get()
		w.Json(blogs)
	})
}

func (h *ZzblogHttp) registerGetBlog() {
	h.router.OnGet("/blogs/:urlid", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		blog := h.zz.Get(r.Bag.Get("urlid").(string), lang)
		if blog == nil {
			w.WithStatusCode(404)
			w.String("Not Found")
			return
		}
		b, e := blog.Detail()
		if e == nil {
			w.Json(b)
			return
		}
		w.WithStatusCode(404)
		w.String("not found")
	})
}

func (h *ZzblogHttp) registerForImage() {
	h.router.OnGet("/images/:hash", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		id := r.Bag.Get("hash").(string)
		log.Printf("id: %s\n", id)
		img := h.zz.GetImage(id)
		if img == nil {
			img = h.zz.GetImageByFilename(id)
		}
		if img == nil {
			w.WithStatusCode(404)
			w.String("not found")
			return
		}
		w.WithHeader("Content-Type", img.MimeType())
		bs, err := img.Resize(uint(r.FormInt("w")), uint(r.FormInt("h")))
		if err != nil {
			w.WithStatusCode(500)
		}
		w.Write(bs)
		w.WithHeader("Content-Type", img.MimeType())
		w.WithStatusCode(200)
	})
	h.router.OnPost("/images", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		src, _, err := r.FormFile("image")
		if r.FormValue("token") != "" {
			w.WithStatusCode(403)
			return
		}
		if err != nil {
			w.WithStatusCode(500)
			return
		}
		err = h.zz.AddImage(src)
		if err != nil {
			w.WithStatusCode(500)
		}
	})
}

func (h *ZzblogHttp) registerGetCates() {
	h.router.OnGet("/cates", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		w.Json(h.zz.Cates("en"))
	})
}

func (h *ZzblogHttp) registerGetTags() {
	h.router.OnGet("/tags", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		w.Json(h.zz.Tags("en"))
	})
}

func (h *ZzblogHttp) registerAuthor() {
	h.router.OnGet("/author", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		author := h.zz.Author()
		if author != nil {
			w.Json(h.zz.Author())
			return
		}
		w.WithStatusCode(404)
		w.String("not found")
	})
}

func (h *ZzblogHttp) Start(addr string) error {
	ms := []httprouter.Middleware{
		&AcrossDomain,
	}
	h.router.Group("/api", ms, func(_ *httprouter.Router) {
		h.registerGetCates()
		h.registerGetBlogs()
		h.registerGetBlog()
		h.registerGetTags()
		h.registerForImage()
		h.registerAuthor()
	})
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h.router)
}
