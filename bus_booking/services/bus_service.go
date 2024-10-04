package services

import (
    "bus_booking/models"
    "bus_booking/repositories"
    "strconv"
)

type BusService struct {
    busRepository *repositories.BusRepository
}

func NewBusService(busRepository *repositories.BusRepository) *BusService {
    return &BusService{busRepository: busRepository}
}

func (s *BusService) CreateBus(bus *models.Bus) error {
    return s.busRepository.Create(bus)
}

func (s *BusService) GetAllBuses() ([]models.Bus, error) {
    return s.busRepository.GetAll()
}

func (s *BusService) GetBusByID(id string) (*models.Bus, error) {
    idInt, err := strconv.Atoi(id)
    if err != nil {
        return nil, err // handle conversion error
    }
    return s.busRepository.GetByID(idInt)
}

func (s *BusService) UpdateBus(id string, bus *models.Bus) error {
    idInt, err := strconv.Atoi(id)
    if err != nil {
        return err // handle conversion error
    }
    return s.busRepository.Update(idInt, bus)
}

func (s *BusService) DeleteBus(id string) error {
    idInt, err := strconv.Atoi(id)
    if err != nil {
        return err // handle conversion error
    }
    return s.busRepository.Delete(idInt)
}
