package main

import (
	"embed"
	"fmt"
	"gnr/app"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]
	router.DB_FLAVOR = "sqlite"

	if arg == "seed" {
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		c := r.ToContext()
		app.Seed(c)
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "balancer" {

	} else if arg == "run" {
		handler := http.HandlerFunc(handleRequest)
		s := &http.Server{
			Addr:    ":8080",
			Handler: handler,
		}

		go s.ListenAndServe()

		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		r.Paths["core"] = app.Core
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + "3001")
	} else if arg == "help" {
	}
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	target := "http://127.0.0.1:3001"

	if strings.Contains(path, "kibana") {
		target = "http://127.0.0.1:5601"
	} else if strings.Contains(path, "esprefix") {
		path = strings.ReplaceAll(path, "/esprefix", "")
		target = "http://127.0.0.1:9200"
		return
	}

	proxyURL, err := url.Parse(target)
	fmt.Println(err)

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = proxyURL.Scheme
		req.URL.Host = proxyURL.Host
		req.URL.Path = path
		req.Host = proxyURL.Host
		req.URL.RawQuery = request.URL.RawQuery
		fmt.Println("req.URL.Path", path)
	}
	proxy.ServeHTTP(writer, request)
}
