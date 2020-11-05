package zzblog

import (
	"encoding/json"
	"errors"
	hr "github.com/yang-zzhong/go-httprouter"
	"github.com/yang-zzhong/logf"
	"html"
	"log"
	"net"
	"net/http"
	"strings"
)

type zzblogRouter struct {
	*hr.Router
}

type ZzblogHttp struct {
	router *zzblogRouter
	zz     Zzblog
}

func (r *zzblogRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res := r.Router.HandleRequest(w, req)
	res.Flush(req)
	logf.Printf("http: %s\t%s\t%v\t%s", req.Method, req.URL.Path, req.Proto, req.RemoteAddr)
}

func NewHttp(root string, docroot string) *ZzblogHttp {
	zz := new(ZzblogHttp)
	log.Printf("doc root: %s\n", docroot)
	log.Printf("blog root: %s\n", root)
	zz.initRouter(docroot)
	zz.zz = NewFileZzblog(root)
	if err := zz.zz.Init(); err != nil {
		panic(err)
	}
	return zz
}

func (h *ZzblogHttp) initRouter(docroot string) {
	h.router = &zzblogRouter{}
	h.router.Router = hr.NewRouter()
	h.router.Router.DocRoot = docroot
	if h.router.Router.DocRoot == "" {
		h.router.Router.Tries = []int{hr.API}
	}
}

func (h *ZzblogHttp) registerTheme() {
	h.router.OnGet("/theme", func(w *hr.Response, r *hr.Request) {
		body, _ := json.Marshal(h.zz.Theme())
		w.WithString(string(body))
	})
}

func (h *ZzblogHttp) registerSitemap() {
	h.router.OnGet("/sitemap.txt", func(w *hr.Response, r *hr.Request) {
		data := []string{GetConfig().Domain + "?lang=en", GetConfig().Domain + "?lang=zh-CN"}
		h.zz.Filter(func(group *LangGroup) *Blog {
			group.Each(func(blog *Blog) bool {
				url := GetConfig().Domain + "/" + blog.URLID + "?lang=" + blog.Lang
				if r.FormValue("format") == "xml" {
					url = html.EscapeString(url)
				}
				data = append(data, url)
				return false
			})
			return nil
		})
		if r.FormValue("format") == "xml" {
			output := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n" +
				"\t<url><loc>" + strings.Join(data, "</loc></url>\n\t<url><loc>") + "</loc></url>\n" +
				"</urlset>\n"
			w.WithString(output)
			return
		}
		w.WithString(strings.Join(data, "\n"))
	})
}

func (h *ZzblogHttp) registerGetBlogs() {
	h.router.OnGet("/blogs", func(w *hr.Response, r *hr.Request) {
		lang := r.FormValue("lang")
		page := 1
		pageSize := 10
		tag := r.FormValue("tag")
		cate := r.FormValue("cate")
		if p, e := r.FormInt("page"); e == nil {
			page = int(p)
		}
		if p, e := r.FormInt("page_size"); e == nil {
			pageSize = int(p)
		}
		blogs := h.zz.Filter(func(group *LangGroup) *Blog {
			return group.One(&OneFilter{lang, cate, tag})
		}).Sort(SC_TIME, ST_DESC).Sort(SC_TITLE, ST_ASC).Page(page, pageSize).Get()
		if b, e := json.Marshal(blogs); e != nil {
			w.InternalError(e)
		} else {
			w.WithString(string(b))
		}
	})
}

func (h *ZzblogHttp) registerGetBlog() {
	h.router.OnGet("/blogs/:urlid", func(w *hr.Response, r *hr.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		blog := h.zz.Get(r.Bag.Get("urlid").(string), lang)
		if blog == nil {
			w.WithStatus(404)
			w.WithString("Not Found")
			return
		}
		res := map[string]interface{}{
			"url_id":       blog.URLID,
			"title":        blog.Title,
			"tags":         blog.Tags,
			"category":     blog.Category,
			"overview":     blog.Overview,
			"lang":         blog.Lang,
			"published_at": blog.PublishedAt,
			"updated_at":   blog.UpdatedAt,
			"image":        blog.Image,
		}
		b, e := blog.Content()
		if e == nil {
			res["content"] = string(b)
			body, _ := json.Marshal(res)
			w.WithString(string(body))
			return
		}
		w.WithStatus(404)
		w.WithString("not found")
	})
}

func (h *ZzblogHttp) registerForImage() {
	h.router.OnGet("/images/:hash", func(w *hr.Response, r *hr.Request) {
		id := r.Bag.Get("hash").(string)
		img := h.zz.GetImage(id)
		if img == nil {
			img = h.zz.GetImageByFilename(id)
		}
		if img == nil {
			w.WithStatus(404)
			w.WithString("not found")
			return
		}
		w.WithHeader("Content-Type", img.MimeType())
		width, _ := r.FormUint("w")
		height, _ := r.FormUint("h")
		bs, err := img.Resize(uint(width), uint(uint(height)))
		if err != nil {
			w.WithStatus(500)
		}
		w.WithString(string(bs))
		w.WithHeader("Content-Type", img.MimeType())
		w.WithStatus(200)
	})
	h.router.OnPost("/images", func(w *hr.Response, r *hr.Request) {
		src, _, err := r.FormFile("image")
		if r.FormValue("token") != "" {
			w.WithStatus(403)
			return
		}
		if err != nil {
			w.WithStatus(500)
			return
		}
		err = h.zz.AddImage(src)
		if err != nil {
			w.WithStatus(500)
		}
	})
}

func (h *ZzblogHttp) registerGetCates() {
	h.router.OnGet("/cates", func(w *hr.Response, r *hr.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		body, _ := json.Marshal(h.zz.Cates(lang))
		w.WithString(string(body))
	})
}

func (h *ZzblogHttp) registerGetTags() {
	h.router.OnGet("/tags", func(w *hr.Response, r *hr.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		body, _ := json.Marshal(h.zz.Tags(lang))
		w.WithString(string(body))
	})
}

func (h *ZzblogHttp) registerAuthor() {
	h.router.OnGet("/author", func(w *hr.Response, r *hr.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = "en"
		}
		author := h.zz.Author(lang)
		if author != nil {
			body, _ := json.Marshal(author)
			w.WithString(string(body))
			return
		}
		w.WithStatus(404)
		w.WithString("not found")
	})
}

func (h *ZzblogHttp) Start(addr string) error {
	ms := []hr.Mw{}
	if GetConfig().Renderer != "" {
		sr := NewServerRenderer(
			GetConfig().Bots,
			GetConfig().Renderer,
			GetConfig().RenderCacheDir,
			h.zz)
		if sr == nil {
			return errors.New("can not create render cache dir")
		}
		h.router.Router.BeforeEntryFile = func(w *hr.Response, req *http.Request, _ string) bool {
			return sr.Before(w, &hr.Request{nil, req})
		}
	}
	if GetConfig().AllowCors {
		ms = append(ms, &AcrossDomain)
	}
	h.router.Group("/api", ms, func(_ *hr.Router) {
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
	log.Printf("listen on %s\n", addr)
	return http.Serve(l, h.router)
}
