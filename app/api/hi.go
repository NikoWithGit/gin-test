package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Hi(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "hi.html", gin.H{
		"content": "This is an index page...",
	})
}
