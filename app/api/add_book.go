package api

import (
	"log"
	"net/http"

	"gin-test/model"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
)

func (s *Server) HandleAddBook(ctx *gin.Context) {
	var book model.Book

	err := ctx.BindJSON(&book)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book.Id = ulid.Now()

	r := s.DB.Create(&book)

	if r.Error != nil {
		log.Fatal(r.Error)
	}

	s.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}
