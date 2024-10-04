package models

type Rating struct {
    ID     int    `json:"id" db:"id"`
    UserID int    `json:"user_id" db:"user_id"`
    BusID  int    `json:"bus_id" db:"bus_id"`
    Score  int    `json:"score" db:"score"` // e.g. 1-5
    Review string `json:"review" db:"review"`
}
