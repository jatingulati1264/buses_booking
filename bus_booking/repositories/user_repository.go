package repositories

import (
    "database/sql"
    "bus_booking/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
    _, err := r.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
    return err
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
    user := &models.User{}
    err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return user, nil
}
