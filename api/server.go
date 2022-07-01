package api

import (
	db "pdv/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Servidor atende solicitação HTTP para nossos serviços
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/v1/products", server.CreateProduto)
	router.GET("/api/v1/products/:codigo_barras", server.GetProduto)
	router.GET("/api/v1/products", server.ListProduto)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {

	return server.router.Run(address)
}
