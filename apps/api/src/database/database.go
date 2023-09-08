package database

import (
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/tithanayut/learnhub-ql/src/config"
	"github.com/tithanayut/learnhub-ql/src/generated/ent"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabaseClient(config config.DatabaseConfig) *ent.Client {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.DatabaseName,
		config.Password,
	)
	client, err := ent.Open(dialect.Postgres, dataSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	return client
}
