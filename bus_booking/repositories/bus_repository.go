package repositories

import (
    "database/sql"
    "bus_booking/models"
)

type BusRepository struct {
    db *sql.DB
}

func NewBusRepository(db *sql.DB) *BusRepository {
    return &BusRepository{db: db}
}

func (r *BusRepository) Create(bus *models.Bus) error {
    _, err := r.db.Exec("INSERT INTO buses (name, capacity, departure, arrival) VALUES (?, ?, ?, ?)", bus.Name, bus.Capacity, bus.Departure, bus.Arrival)
    return err
}

func (r *BusRepository) GetAll() ([]models.Bus, error) {
    rows, err := r.db.Query("SELECT id, name, capacity, departure, arrival FROM buses")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var buses []models.Bus
    for rows.Next() {
        var bus models.Bus
        if err := rows.Scan(&bus.ID, &bus.Name, &bus.Capacity, &bus.Departure, &bus.Arrival); err != nil {
            return nil, err
        }
        buses = append(buses, bus)
    }
    return buses, nil
}

func (r *BusRepository) GetByID(id int) (*models.Bus, error) {
    bus := &models.Bus{}
    err := r.db.QueryRow("SELECT id, name, capacity, departure, arrival FROM buses WHERE id = ?", id).Scan(&bus.ID, &bus.Name, &bus.Capacity, &bus.Departure, &bus.Arrival)
    if err != nil {
        return nil, err
    }
    return bus, nil
}

func (r *BusRepository) Update(id int, bus *models.Bus) error {
    _, err := r.db.Exec("UPDATE buses SET name = ?, capacity = ?, departure = ?, arrival = ? WHERE id = ?", bus.Name, bus.Capacity, bus.Departure, bus.Arrival, id)
    return err
}

func (r *BusRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM buses WHERE id = ?", id)
    return err
}
