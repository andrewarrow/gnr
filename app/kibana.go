package app

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/andrewarrow/feedback/router"
)

func Kibana(c *router.Context, second, third string) {
	fmt.Println("1111")
	query := c.Request.URL.RawQuery
	target := "http://127.0.0.1:5601/kibana"
	if query != "" {
		target += "?" + query
	}
	proxyURL, err := url.Parse(target)
	fmt.Println(err)

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	c.Request.Host = proxyURL.Host
	fmt.Println("1111222")
	proxy.ServeHTTP(c.Writer, c.Request)
}
