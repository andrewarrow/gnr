package app

import (
	"gnr/reddit"

	"codeberg.org/andrewarrow/roll"
	"github.com/andrewarrow/feedback/router"
)

func Seed(c *router.Context) {
	posts := reddit.GetPosts("gunsnroses")

	r := roll.NewRoll("posts")

	for _, post := range posts {
		id := post.GetString("id_reddit")
		c.Params = map[string]any{}
		c.Params["title"] = post.GetString("title")
		c.Params["url"] = post.GetString("url")
		c.Params["id_reddit"] = id
		c.ValidateAndInsert("post")

		delete(c.Params, "id_reddit")
		rq := roll.NewUpsert(id, c.Params)
		r.MakeNetworkPut(id, rq.Render())

	}
}
