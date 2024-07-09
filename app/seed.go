package app

import "github.com/andrewarrow/feedback/router"

func Seed(c *router.Context) {
	posts := reddit.GetPosts("gunsnroses")
	c.Parans = map[string]any{}
	for _, post := range posts {
		c.Params["title"] = post.GetString("title")
	}
	c.ValidateAndInsert("post")
}
