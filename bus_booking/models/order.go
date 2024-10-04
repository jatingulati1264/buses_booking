package models

type Order struct {
    ID     int    `json:"id" db:"id"`
    UserID int    `json:"user_id" db:"user_id"`
    BusID  int    `json:"bus_id" db:"bus_id"`
    Status string `json:"status" db:"status"` // e.g. "pending", "confirmed"
}
