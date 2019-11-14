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

func NewHttp(root string) *ZzblogHttp {
	zz := new(ZzblogHttp)
	zz.initRouter()
	zz.zz = NewFileZzblog(root)
	return zz
}

func (h *ZzblogHttp) initRouter() {
	h.router = httprouter.NewRouter()
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
		for _, blog := range blogs {
			log.Printf("%v\n", blog)
		}
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

func (h *ZzblogHttp) Start(addr string) error {
	h.router.Group("/api", nil, func(_ *httprouter.Router) {
		h.registerGetCates()
		h.registerGetBlogs()
		h.registerGetBlog()
		h.registerGetTags()
	})
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h.router)
}
