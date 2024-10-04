package models

type Bus struct {
    ID        int    `json:"id" db:"id"`
    Name      string `json:"name" db:"name"`
    Capacity  int    `json:"capacity" db:"capacity"`
    Departure string `json:"departure" db:"departure"`
    Arrival   string `json:"arrival" db:"arrival"`
}
