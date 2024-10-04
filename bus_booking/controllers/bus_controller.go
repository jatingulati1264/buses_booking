package controllers

import (
    "net/http"
    "bus_booking/models"
    "bus_booking/services"
    "github.com/gin-gonic/gin"
)

type BusController struct {
    service services.BusService
}

func NewBusController(service services.BusService) *BusController {
    return &BusController{service: service}
}

func (ctrl *BusController) CreateBus(c *gin.Context) {
    var bus models.Bus
    if err := c.ShouldBindJSON(&bus); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.service.CreateBus(&bus); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, bus)
}

func (ctrl *BusController) GetAllBuses(c *gin.Context) {
    buses, err := ctrl.service.GetAllBuses()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, buses)
}

func (ctrl *BusController) GetBusByID(c *gin.Context) {
    id := c.Param("id")
    bus, err := ctrl.service.GetBusByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bus not found"})
        return
    }

    c.JSON(http.StatusOK, bus)
}

func (ctrl *BusController) UpdateBus(c *gin.Context) {
    id := c.Param("id")
    var bus models.Bus
    if err := c.ShouldBindJSON(&bus); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.service.UpdateBus(id, &bus); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, bus)
}

func (ctrl *BusController) DeleteBus(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.service.DeleteBus(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}
