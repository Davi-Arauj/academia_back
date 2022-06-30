package api

import (
	db "pdv/db/sqlc"

	"github.com/gin-contrib/cors"
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
	router.GET("/api/v1/products/:cod_barras", server.GetProduto)
	router.GET("/api/v1/products", server.ListProduto)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {

	server.router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:4200"},
		AllowMethods:  []string{"PUT", "PATCH", "GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	return server.router.Run(address)
}
