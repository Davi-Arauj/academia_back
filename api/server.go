package api

import (
	db "pdv/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Servidor atende solicitação HTTP para nossos serviços
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/v1/produtos", server.createProduto)
	router.GET("/api/v1/produtos/:cod_barras", server.getProduto)
	router.GET("/api/v1/produtos", server.listProduto)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
