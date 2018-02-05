package router

import (
	"github.com/chinx/sourceproxy/handler"
	"github.com/chinx/cobweb"
)

func InitRouters(mux *cobweb.Router)  {
	mux.Group("/sourceproxy/v1/:provider", func() {
		mux.Group("/oauth2", func() {
			mux.Post("/", handler.OAuthURL)
			mux.Post("/token",handler.AccessToken)
		})

		mux.Get("/user", handler.SourceUser)

		mux.Group("/repos", func() {
			mux.Get("/", handler.SourceRepos)
			mux.Group("/:owner/:repo", func() {
				mux.Get("/branches", handler.SourceRepoBranches)
				mux.Get("/commits", handler.SourceRepoCommits)
			})
		})
	})
}