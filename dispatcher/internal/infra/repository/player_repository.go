package repository

import (
	"context"
	"database/sql"
	"dispatcher/internal/domain/entity"
	"dispatcher/internal/infra/db"
)

type PlayerRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewPlayerRepository(dbConn *sql.DB) *PlayerRepository {
	return &PlayerRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *PlayerRepository) Create(ctx context.Context, player *entity.Player) error {
	err := r.Queries.CreatePlayer(ctx, db.CreatePlayerParams{
		ID:    player.ID,
		Name:  player.Name,
		Price: player.Price,
	})
	return err
}

func (r *PlayerRepository) FindByID(ctx context.Context, id string) (*entity.Player, error) {
	player, err := r.Queries.FindPlayerById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Player{
		ID:    player.ID,
		Name:  player.Name,
		Price: player.Price,
	}, nil
}
