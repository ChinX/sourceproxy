package main

import (
	"log"
	"net/http"

	"github.com/chinx/cobweb"
	_ "github.com/chinx/sourceproxy/modules/oauth/gitee"
	_ "github.com/chinx/sourceproxy/modules/oauth/github"
	_ "github.com/chinx/sourceproxy/modules/oauth/gitlab"
	"github.com/chinx/sourceproxy/router"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mux := cobweb.NewRouter()
	router.InitRouters(mux)
	http.ListenAndServe(":8080", mux)
}
