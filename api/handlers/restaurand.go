package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Project_Restaurant/api-gateway/genproto/restaurant"

	"github.com/gin-gonic/gin"
)

// CreateRestaurantHandler creates a new restaurant.
// @Summary Create new restaurant
// @Description Create new restaurant
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param restaurant body pb.CreateRestaurantRequest true "Restaurant"
// @Success 201 {object} pb.CreateRestaurantResponse "Restaurant created"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants [post]
func (h *HandlerStruct) CreateRestaurantHandler(c *gin.Context) {
	var req pb.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Restaurant.CreateRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetRestaurantHandler gets a restaurant by ID.
// @Summary Get restaurant by ID
// @Description Get restaurant by ID
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Restaurant ID"
// @Success 200 {object} pb.GetRestaurantResponse "Restaurant found"
// @Failure 404 {object} string "Restaurant not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants/{id} [get]
func (h *HandlerStruct) GetRestaurantHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetRestaurantRequest{Id: id}

	resp, err := h.Restaurant.GetRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateRestaurantHandler updates an existing restaurant.
// @Summary Update restaurant
// @Description Update restaurant
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Restaurant ID"
// @Param restaurant body pb.UpdateRestaurantRequest true "Restaurant"
// @Success 200 {object} pb.UpdateRestaurantResponse "Restaurant updated"
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Restaurant not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants/{id} [put]
func (h *HandlerStruct) UpdateRestaurantHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Restaurant.Id = id
	resp, err := h.Restaurant.UpdateRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteRestaurantHandler deletes a restaurant.
// @Summary Delete restaurant
// @Description Delete restaurant
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Restaurant ID"
// @Success 204 {object} string "Restaurant deleted"
// @Failure 404 {object} string "Restaurant not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants/{id} [delete]
func (h *HandlerStruct) DeleteRestaurantHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteRestaurantRequest{Id: id}

	_, err := h.Restaurant.DeleteRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListRestaurantsHandler gets a list of restaurants.
// @Summary Get list of restaurants
// @Description Get list of restaurants, optionally filtered by name and address, and paginated.
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param name query string false "Restaurant name"
// @Param address query string false "Restaurant address"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} pb.ListRestaurantsResponse "List of restaurants"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants [get]
func (h *HandlerStruct) ListRestaurantsHandler(c *gin.Context) {
	name := c.Query("name")
	address := c.Query("address")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	var pageInt, limitInt int64
	var err error
	if pageInt, err = strconv.ParseInt(page, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page format"})
		return
	}
	if limitInt, err = strconv.ParseInt(limit, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit format"})
		return
	}

	req := pb.ListRestaurantsRequest{
		Name:    name,
		Address: address,
		Page:    int32(pageInt),
		Limit:   int32(limitInt),
	}

	resp, err := h.Restaurant.ListRestaurants(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
