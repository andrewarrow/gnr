package app

import (
	"gnr/reddit"

	"github.com/andrewarrow/feedback/router"
)

func Seed(c *router.Context) {
	posts := reddit.GetPosts("gunsnroses")
	for _, post := range posts {
		c.Params = map[string]any{}
		c.Params["title"] = post.GetString("title")
		c.Params["url"] = post.GetString("url")
		c.Params["id_reddit"] = post.GetString("id_reddit")
		c.ValidateAndInsert("post")
	}
}
