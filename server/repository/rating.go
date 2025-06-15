package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type RatingRepository struct {
	DB *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{DB: db}
}

type Rating struct {
	ID        int64     `json:"id"`
	Score     int64     `json:"score"`
	Review    string    `json:"review"`
	VendorId  int64     `json:"vendor_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (repo *RatingRepository) InsertRating(rating Rating) (*Rating, error) {
	result, err := DB.Exec(
		"INSERT INTO ratings (score, review, vendor_id) VALUES (?, ?, ?)",
		rating.Score, rating.Review, rating.VendorId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert rating: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get inserted ID: %w", err)
	}

	var created Rating
	err = DB.QueryRow(
		"SELECT id, score, review, vendor_id, created_at FROM ratings WHERE id = ?",
		id,
	).Scan(&created.ID, &created.Score, &created.Review, &created.VendorId, &created.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inserted vendor: %w", err)
	}

	return &created, nil
}

func (repo *RatingRepository) GetRatingById(rating_id int64, vendor_id int64) (*Rating, error) {
	var rating Rating
	err := DB.QueryRow(
		"SELECT id, score, review, vendor_id, created_at FROM ratings WHERE id = ? AND vendor_id = ?",
		rating_id,
		vendor_id,
	).Scan(&rating.ID, &rating.Score, &rating.Review, &rating.VendorId, &rating.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rating: %w", err)
	}

	return &rating, nil
}

func (repo *RatingRepository) GetRatingsByVendorId(vendor_id int64) ([]Rating, error) {
	rows, err := DB.Query("SELECT id, vendor_id, score, review, created_at FROM ratings WHERE vendor_id = ?", vendor_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []Rating
	for rows.Next() {
		var r Rating
		err := rows.Scan(&r.ID, &r.VendorId, &r.Score, &r.Review, &r.CreatedAt)
		if err != nil {
			return nil, err
		}

		ratings = append(ratings, r)
	}

	return ratings, nil
}
