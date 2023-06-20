package api

func (s *Server) RegisterRoutes() {
	s.Gin.GET("/hi", s.Hi)
	s.Gin.GET("/books/:isbn", s.HandleGetByISBN)
	s.Gin.POST("/books", s.HandleAddBook)
}
