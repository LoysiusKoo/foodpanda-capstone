package db

import (
	"context"
	"database/sql"
	"fmt"

	null "gopkg.in/guregu/null.v4"
)

// Store provides all function to execute db queries and transactions
type Store interface {
	Querier
	CreatePlaylistTx(ctx context.Context, arg CreatePlaylistTxParams) (Playlist, error)
}

// SQLStore provides all function to execute db queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx excutes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		tx.Rollback()
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

/*
Transaction to create
//1. Generate a playlist based on cuisine and number of restaurant item, and budget (per restaurant item)
		- create a playlist
		- insert
//2. create a playlist, add dishes to playlist
*/

type CreatePlaylistTxParams struct {
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	ImageUrl     null.String `json:"image_url"`
	IsPublic     bool        `json:"is_public"`
	DeliveryDay  null.String `json:"delivery_day"`
	Category     null.String `json:"category"`
	UserID       int64       `json:"user_id"`
	PlaylistID   int64       `json:"playlist_id"`
	DeliveryTime null.Time   `json:"delivery_time"`
	Status       null.String `json:"status"`
}

func (store *SQLStore) CreatePlaylistTx(ctx context.Context, arg CreatePlaylistTxParams) (Playlist, error) {

	var playlist Playlist
	var err error

	err = store.execTx(ctx, func(q *Queries) error {

		playlist, err = q.CreatePlaylist(ctx, CreatePlaylistParams{
			Name:        arg.Name,
			Description: arg.Description,
			ImageUrl:    arg.ImageUrl,
			IsPublic:    arg.IsPublic,
			DeliveryDay: arg.DeliveryDay,
			Category:    arg.Category,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateUser_Playlist(ctx, CreateUser_PlaylistParams{
			UserID:       arg.UserID,
			PlaylistID:   playlist.ID,
			DeliveryDay:  arg.DeliveryDay,
			DeliveryTime: arg.DeliveryTime,
			Status:       arg.Status,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return playlist, err
}
