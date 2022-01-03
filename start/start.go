package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"saas/ent/migrate"
	"saas/internal/database"
)

func main() {
	client,_ := database.Open()
	defer client.Close()
	err := client.Debug().Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
