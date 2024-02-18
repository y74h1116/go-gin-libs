package main

import (
	"net/http"

	"github.com/y74h1116/go-gin-libs/page"
	"github.com/gin-gonic/gin"
)

func (home *Home) index(ctx *gin.Context) {
	viewParam := home.pageHome.CreateViewParam()
	ctx.HTML(http.StatusOK, "home/index", viewParam)
}

type Home struct {
	pageHome *page.Page
}

func NewHome() *Home {
	pageHome := page.NewPage("ホーム", map[string]map[string]string{})
	return &Home{pageHome}
}
