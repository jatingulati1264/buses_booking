package repositories

import (
    "database/sql"
    "bus_booking/models"
)

type RatingRepository struct {
    db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
    return &RatingRepository{db: db}
}

func (r *RatingRepository) Create(rating *models.Rating) error {
    _, err := r.db.Exec("INSERT INTO ratings (user_id, bus_id, score, review) VALUES (?, ?, ?, ?)", rating.UserID, rating.BusID, rating.Score, rating.Review)
    return err
}

func (r *RatingRepository) GetAll() ([]models.Rating, error) {
    rows, err := r.db.Query("SELECT id, user_id, bus_id, score, review FROM ratings")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ratings []models.Rating
    for rows.Next() {
        var rating models.Rating
        if err := rows.Scan(&rating.ID, &rating.UserID, &rating.BusID, &rating.Score, &rating.Review); err != nil {
            return nil, err
        }
        ratings = append(ratings, rating)
    }
    return ratings, nil
}

// Implement other CRUD methods as needed
