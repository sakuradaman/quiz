package infra

import (
	"context"
	"fmt"
	"log"

	"github.com/sakuradaman/quiz/pkg/ent"
	"github.com/sakuradaman/quiz/pkg/lib/config"
	"golang.org/x/xerrors"
)

func NewDBConnector(cfg *config.Config) (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:@tcp(%s)/%s?parseTime=true&loc=Local", cfg.DBUser, cfg.DBAddress, cfg.DBName)
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, xerrors.Errorf("db connection failed：%w", err)
	}
	return db, nil
}

// Run the auto migration tool.
func NewDBCreate(ctx context.Context, db *ent.Client) {
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
