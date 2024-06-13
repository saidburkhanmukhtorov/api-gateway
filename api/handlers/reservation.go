package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Project_Restaurant/api-gateway/genproto/reservation"
	"github.com/gin-gonic/gin"
)

// CreateReservationHandler creates a new reservation.
// @Summary Create new reservation
// @Description Create new reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param reservation body reservation.CreateReservationRequest true "Reservation"
// @Success 201 {object} reservation.CreateReservationResponse "Reservation created"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations [post]
func (h *HandlerStruct) CreateReservationHandler(c *gin.Context) {
	var req reservation.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Reservation.CreateReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetReservationHandler gets a reservation by ID.
// @Summary Get reservation by ID
// @Description Get reservation by ID
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Success 200 {object} reservation.GetReservationResponse "Reservation found"
// @Failure 404 {object} string "Reservation not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id} [get]
func (h *HandlerStruct) GetReservationHandler(c *gin.Context) {
	id := c.Param("id")
	req := reservation.GetReservationRequest{Id: id}

	resp, err := h.Reservation.GetReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateReservationHandler updates an existing reservation.
// @Summary Update reservation
// @Description Update reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Param reservation body reservation.UpdateReservationRequest true "Reservation"
// @Success 200 {object} reservation.UpdateReservationResponse "Reservation updated"
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Reservation not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id} [put]
func (h *HandlerStruct) UpdateReservationHandler(c *gin.Context) {
	id := c.Param("id")
	var req reservation.UpdateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Reservation.Id = id
	resp, err := h.Reservation.UpdateReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteReservationHandler deletes a reservation.
// @Summary Delete reservation
// @Description Delete reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Success 204 {object} string "Reservation deleted"
// @Failure 404 {object} string "Reservation not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id} [delete]
func (h *HandlerStruct) DeleteReservationHandler(c *gin.Context) {
	id := c.Param("id")
	req := reservation.DeleteReservationRequest{Id: id}

	_, err := h.Reservation.DeleteReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListReservationsHandler gets a list of reservations.
// @Summary Get list of reservations
// @Description Get list of reservations
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param user_id query string false "User ID"
// @Param restaurant_id query string false "Restaurant ID"
// @Param status query string false "Reservation status"
// @Param start_time query string false "Start time"
// @Param end_time query string false "End time"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} reservation.ListReservationsResponse "List of reservations"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations [get]
func (h *HandlerStruct) ListReservationsHandler(c *gin.Context) {
	req := reservation.ListReservationsRequest{
		UserId:       c.Query("user_id"),
		RestaurantId: c.Query("restaurant_id"),
		Status:       c.Query("status"),
		StartTime:    c.Query("start_time"),
		EndTime:      c.Query("end_time"),
	}

	resp, err := h.Reservation.ListReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CheckAvailabilityHandler checks reservation availability.
// @Summary Check reservation availability
// @Description Check reservation availability
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param restaurant_id query string true "Restaurant ID"
// @Param reservation_time query string true "Reservation time"
// @Success 200 {object} reservation.CheckAvailabilityResponse "Availability status"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/check [get]
func (h *HandlerStruct) CheckAvailabilityHandler(c *gin.Context) {
	req := reservation.CheckAvailabilityRequest{
		RestaurantId:    c.Query("restaurant_id"),
		ReservationTime: c.Query("reservation_time"),
	}

	resp, err := h.Reservation.CheckAvailability(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetFoodListHandler gets a list of available foods for a reservation.
// @Summary Get list of foods for reservation
// @Description Get list of foods for a reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param restaurant_id query string true "Restaurant ID"
// @Success 200 {object} reservation.OrderFoodListRes "List of foods"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id}/foodlist [get]
func (h *HandlerStruct) GetFoodListHandler(c *gin.Context) {
	id := c.Query("restaurant_id")
	log.Println(id)
	req := reservation.OrderFoodListReq{RestaurantId: id}

	resp, err := h.Reservation.FoodList(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// OrderFoodHandler orders food for a reservation.
// @Summary Order food for reservation
// @Description Order food for reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Param order body reservation.OrderFoodReq true "Order details"
// @Success 200 {object} reservation.OrderFoodRes "Order confirmation"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id}/order [post]
func (h *HandlerStruct) OrderFoodHandler(c *gin.Context) {
	id := c.Param("id")
	var req reservation.OrderFoodReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ReservationId = id
	resp, err := h.Reservation.OrderFood(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// OrderBillHandler processes the bill for a reservation.
// @Summary Process bill for reservation
// @Description Process bill for reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Success 200 {object} reservation.OrderBillRes "Bill details"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id}/bill [post]
func (h *HandlerStruct) OrderBillHandler(c *gin.Context) {
	var req reservation.OrderBillReq
	req.ReservationId = c.Param("id")
	log.Println(req)
	resp, err := h.Reservation.OrderBill(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
