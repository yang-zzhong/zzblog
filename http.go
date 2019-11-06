package zzblog

import (
	httprouter "github.com/yang-zzhong/go-httprouter"
)

type ZzblogHttp struct {
	router * httprouter.Router
	zz Zzblog
}

func NewZzblogHttp() *ZzblogHttp {
	zz := new(ZzblogHttp)
	zz.initRouter()
	zz.zz = NewFileZzblog("./test")
	return zz
}

func (h *ZzblogHttp) initRouter() {
	h.router = httprouter.NewRouter()
}

func (h *ZzblogHttp) registerGetBlogs() {
	h.router.OnGet("/blogs", func(w * httprouter.ResponseWriter, r * httprouter.Request) {
		h.zz.Filter(func(_ * Blog) bool {
			return true
		})
	})
}

func (h *ZzblogHttp) registerGetBlog() {
	h.router.OnGet("/blogs/:urlid", func(w * httprouter.ResponseWriter, r * httprouter.Request) {
		blog := h.zz.Get(r.Bag.Get("urlid").(string))
		if blog == nil {
			w.String("Not Found")
			return
		}
	})
}

func (h *ZzblogHttp) registerGetCates() {
	h.router.OnGet("/cates", func(w * httprouter.ResponseWriter, r * httprouter.Request) {
	})
}

func (h *ZzblogHttp) registerGetTags() {
	h.router.OnGet("/tags", func(w * httprouter.ResponseWriter, r * httprouter.Request) {
	})
}

func (h *ZzblogHttp)StartHttp(addr string) {
	h.router.Group("/api", nil, func (_ * httprouter.Router) {
		h.registerGetCates()
		h.registerGetBlogs()
		h.registerGetBlog()
		h.registerGetTags()
	})
}