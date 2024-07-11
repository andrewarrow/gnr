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
		handler := http.HandlerFunc(handleRequest)
		s := &http.Server{
			Addr:    ":443",
			Handler: handler,
		}

		s.ListenAndServe()

	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		r.Paths["core"] = app.Core
		r.Paths["kibana"] = app.Kibana
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	query := c.Request.URL.RawQuery
	target := "http://127.0.0.1:5601"
	if query != "" {
		target += "?" + query
	}
	proxyURL, err := url.Parse(target)
	fmt.Println(err)

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	c.Request.Host = proxyURL.Host
	proxy.ServeHTTP(writer, request)
}
