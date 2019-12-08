package zzblog

import (
	"crypto/md5"
	"encoding/hex"
	httprouter "github.com/yang-zzhong/go-httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

type ServerRenderer struct {
	bots     []string
	server   string
	cacheDir string
}

func NewServerRenderer(bots []string, server string, cachedir string) *ServerRenderer {
	_, err := os.Stat(cachedir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil
		}
		if err = os.MkdirAll(cachedir, 0755); err != nil {
			return nil
		}
	}
	return &ServerRenderer{bots, server, cachedir}
}

func (sr *ServerRenderer) Before(w *httprouter.ResponseWriter, req *httprouter.Request) bool {
	log.Printf("user-agent %s\n", req.Header.Get("User-Agent"))
	for _, bot := range sr.bots {
		reg := regexp.MustCompile("(?i)" + bot)
		if reg.MatchString(req.Header.Get("User-Agent")) {
			log.Printf("begin renderer\n")
			return !sr.Render(w, req.Request)
		}
	}
	return true
}

func (sr *ServerRenderer) After(w *httprouter.ResponseWriter, _ *httprouter.Request) bool {
	return true
}

func (sr *ServerRenderer) Render(w *httprouter.ResponseWriter, req *http.Request) bool {
	target := sr.url(req)
	lang := req.FormValue("lang")
	if lang == "" {
		lang = "en"
	}
	sum := md5.Sum([]byte(target))
	h := hex.EncodeToString(sum[:])
	cachefile := path.Join(sr.cacheDir, h+lang+".html")
	file, err := os.Open(cachefile)
	defer file.Close()
	write := func(r io.Reader, w *httprouter.ResponseWriter) {
		if data, err := ioutil.ReadAll(r); err != nil {
			w.WithStatusCode(500)
			w.Write([]byte("read body of server render error"))
		} else {
			w.Write(data)
		}
	}
	if err == nil {
		write(file, w)
		return true
	}
	if !os.IsNotExist(err) {
		return false
	}
	file, err = os.OpenFile(cachefile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return false
	}
	u := sr.server + "/render/" + url.QueryEscape(target)
	log.Printf("url: %s\n", u)
	res, err := http.Get(u)
	if err != nil {
		log.Printf("render error: %v\n", err)
		return false
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("render status code %d\n", res.StatusCode)
		return false
	}
	io.Copy(file, res.Body)
	file.Seek(0, os.SEEK_SET)
	write(file, w)
	return true
}

func (sr *ServerRenderer) url(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
}
