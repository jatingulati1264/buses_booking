package services

import (
    "bus_booking/models"
    "bus_booking/repositories"
)

type OrderService struct {
    orderRepository *repositories.OrderRepository
    busRepository   *repositories.BusRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository, busRepository *repositories.BusRepository) *OrderService {
    return &OrderService{orderRepository: orderRepository, busRepository: busRepository}
}

// Implement Order service methods
func (s *OrderService) CreateOrder(order *models.Order) error {
    return s.orderRepository.Create(order)
}

// Implement more methods as needed
