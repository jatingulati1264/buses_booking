package repositories

import (
    "database/sql"
    "bus_booking/models"
)

type OrderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
    _, err := r.db.Exec("INSERT INTO orders (user_id, bus_id, status) VALUES (?, ?, ?)", order.UserID, order.BusID, order.Status)
    return err
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
    rows, err := r.db.Query("SELECT id, user_id, bus_id, status FROM orders")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.ID, &order.UserID, &order.BusID, &order.Status); err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }
    return orders, nil
}

// Implement other CRUD methods as needed
