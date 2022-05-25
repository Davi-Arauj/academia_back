package main

import (
	"database/sql"
	"log"
	"pdv/api"
	db "pdv/db/sqlc"
	"pdv/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Não foi possivel se conectar ao Banco de Dados:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Não foi possivel iniciar o Servidor:", err)
	}

}
