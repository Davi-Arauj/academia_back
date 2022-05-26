package db

import (
	"database/sql"
	"log"
	"os"
	"pdv/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Não foi possivel se conectar ao Banco de Dados:", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
