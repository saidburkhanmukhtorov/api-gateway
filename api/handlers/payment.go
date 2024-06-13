package handlers

import (
	"context"
	"net/http"

	pb "github.com/Project_Restaurant/api-gateway/genproto/payment" // Update this path to match your actual protobuf path

	"github.com/gin-gonic/gin"
)

// CreatePaymentHandler creates a new payment.
// @Summary Create new payment
// @Description Create new payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param payment body pb.CreatePaymentRequest true "Payment"
// @Success 201 {object} pb.CreatePaymentResponse "Payment created"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments [post]
func (h *HandlerStruct) CreatePaymentHandler(c *gin.Context) {
	var req pb.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Payment.CreatePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetPaymentHandler gets a payment by ID.
// @Summary Get payment by ID
// @Description Get payment by ID
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} pb.GetPaymentResponse "Payment found"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments/{id} [get]
func (h *HandlerStruct) GetPaymentHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetPaymentRequest{Id: id}

	resp, err := h.Payment.GetPayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdatePaymentHandler updates an existing payment.
// @Summary Update payment
// @Description Update payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Param payment body pb.UpdatePaymentRequest true "Payment"
// @Success 200 {object} pb.UpdatePaymentResponse "Payment updated"
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments/{id} [put]
func (h *HandlerStruct) UpdatePaymentHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Payment.Id = id
	resp, err := h.Payment.UpdatePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeletePaymentHandler deletes a payment.
// @Summary Delete payment
// @Description Delete payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 204 {object} string "Payment deleted"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments/{id} [delete]
func (h *HandlerStruct) DeletePaymentHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeletePaymentRequest{Id: id}

	_, err := h.Payment.DeletePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// PayForReservationHandler processes payment for a reservation.
// @Summary Pay for reservation
// @Description Pay for reservation
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param payment body pb.PayForReservationReq true "Payment details"
// @Success 200 {object} pb.PayForReservationRes "Payment confirmation"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments/reservation [post]
func (h *HandlerStruct) PayForReservationHandler(c *gin.Context) {
	var req pb.PayForReservationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Payment.PayForReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
