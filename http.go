package zzblog

import (
	"errors"
	httprouter "github.com/yang-zzhong/go-httprouter"
	"net"
	"net/http"
	"strings"
	"zzblog/log"
)

type zzblogRouter struct {
	*httprouter.Router
}

type ZzblogHttp struct {
	router *zzblogRouter
	zz     Zzblog
}

func (r *zzblogRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rw := r.Router.HandleRequest(w, req)
	log.Printf("http", "%s\t%s\t%v\t%d\t%s", req.Method, req.URL.Path, req.Proto, rw.StatusCode, req.RemoteAddr)
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
	h.router = &zzblogRouter{}
	h.router.Router = httprouter.NewRouter()
	h.router.Router.DocRoot = docroot
	if h.router.Router.DocRoot == "" {
		h.router.Router.Tries = []string{httprouter.Api}
	}
}

func (h *ZzblogHttp) registerTheme() {
	h.router.OnGet("/theme", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		w.Json(h.zz.Theme())
	})
}

func (h *ZzblogHttp) registerSitemap() {
	h.router.OnGet("/sitemap.txt", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		data := []string{GetConfig().Domain + "?lang=en", GetConfig().Domain + "?lang=zh-CN"}
		h.zz.Filter(func(group *LangGroup) *Blog {
			group.Each(func(blog *Blog) bool {
				data = append(data, GetConfig().Domain+"/"+blog.URLID+"?lang="+blog.Lang)
				return false
			})
			return nil
		})
		w.WriteString(strings.Join(data, "\n"))
	})
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
		blogs := h.zz.Filter(func(group *LangGroup) *Blog {
			return group.One(&OneFilter{lang, cate, tag})
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
		res := map[string]interface{}{
			"url_id":     blog.URLID,
			"title":      blog.Title,
			"tags":       blog.Tags,
			"category":   blog.Category,
			"overview":   blog.Overview,
			"lang":       blog.Lang,
			"created_at": blog.CreatedAt,
			"updated_at": blog.UpdatedAt,
			"image":      blog.Image,
		}
		if e == nil {
			res["content"] = string(b.Content)
			w.Json(res)
			return
		}
		w.WithStatusCode(404)
		w.String("not found")
	})
}

func (h *ZzblogHttp) registerForImage() {
	h.router.OnGet("/images/:hash", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		id := r.Bag.Get("hash").(string)
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
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		w.Json(h.zz.Cates(lang))
	})
}

func (h *ZzblogHttp) registerGetTags() {
	h.router.OnGet("/tags", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		w.Json(h.zz.Tags(lang))
	})
}

func (h *ZzblogHttp) registerAuthor() {
	h.router.OnGet("/author", func(w *httprouter.ResponseWriter, r *httprouter.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		author := h.zz.Author(lang)
		if author != nil {
			w.Json(author)
			return
		}
		w.WithStatusCode(404)
		w.String("not found")
	})
}

func (h *ZzblogHttp) Start(addr string) error {
	ms := []httprouter.Middleware{}
	if GetConfig().Renderer != "" {
		sr := NewServerRenderer(
			GetConfig().Bots,
			GetConfig().Renderer,
			GetConfig().RenderCacheDir,
			h.zz)
		if sr == nil {
			return errors.New("can not create render cache dir")
		}
		h.router.Router.BeforeEntryFile = func(w *httprouter.ResponseWriter, req *http.Request, _ string) bool {
			return sr.Before(w, &httprouter.Request{nil, req})
		}
	}
	if GetConfig().AllowCors {
		ms = append(ms, &AcrossDomain)
	}
	h.router.Group("/api", ms, func(_ *httprouter.Router) {
		h.registerGetCates()
		h.registerGetBlogs()
		h.registerGetBlog()
		h.registerGetTags()
		h.registerForImage()
		h.registerAuthor()
		h.registerTheme()
		h.registerSitemap()
	})
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h.router)
}
