package api

import (
	"log"
	"net/http"

	"gin-test/model"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetByISBN(ctx *gin.Context) {
	var book model.Book

	log.Println("CONTACT")
	isbn := ctx.Param("isbn")

	ret := s.DB.First(&book, "isbn = ?", isbn)

	if ret.RowsAffected == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ret.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, ret.Error)
		return
	}

	ctx.JSON(http.StatusOK, book)
}
