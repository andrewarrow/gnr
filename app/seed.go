package app

import (
	"gnr/reddit"

	"github.com/andrewarrow/feedback/router"
)

func Seed(c *router.Context) {
	posts := reddit.GetPosts("gunsnroses")
	c.Params = map[string]any{}
	for _, post := range posts {
		c.Params["title"] = post.GetString("title")
	}
	c.ValidateAndInsert("post")
}
