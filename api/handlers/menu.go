package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Project_Restaurant/api-gateway/genproto/menu" // Update this path to match your actual protobuf path

	"github.com/gin-gonic/gin"
)

// CreateMenuHandler creates a new menu item.
// @Summary Create new menu item
// @Description Create new menu item
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param menu body pb.CreateMenuRequest true "Menu item"
// @Success 201 {object} pb.CreateMenuResponse "Menu item created"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menus [post]
func (h *HandlerStruct) CreateMenuHandler(c *gin.Context) {
	var req pb.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Menu.CreateMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetMenuHandler gets a menu item by ID.
// @Summary Get menu item by ID
// @Description Get menu item by ID
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Menu item ID"
// @Success 200 {object} pb.GetMenuResponse "Menu item found"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menus/{id} [get]
func (h *HandlerStruct) GetMenuHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetMenuRequest{Id: id}

	resp, err := h.Menu.GetMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateMenuHandler updates an existing menu item.
// @Summary Update menu item
// @Description Update menu item
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Menu item ID"
// @Param menu body pb.UpdateMenuRequest true "Menu item"
// @Success 200 {object} pb.UpdateMenuResponse "Menu item updated"
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menus/{id} [put]
func (h *HandlerStruct) UpdateMenuHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Menu.Id = id
	resp, err := h.Menu.UpdateMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteMenuHandler deletes a menu item.
// @Summary Delete menu item
// @Description Delete menu item
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Menu item ID"
// @Success 204 {object} string "Menu item deleted"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menus/{id} [delete]
func (h *HandlerStruct) DeleteMenuHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteMenuRequest{Id: id}

	_, err := h.Menu.DeleteMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllMenusHandler gets a list of all menus for a restaurant.
// @Summary Get all menus for a restaurant
// @Description Get all menus for a restaurant, optionally filtered by name, description, and price range.
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param restaurant_id query string true "Restaurant ID"
// @Param name query string false "Menu item name"
// @Param description query string false "Menu item description"
// @Param minPrice query string false "Minimum price"
// @Param maxPrice query string false "Maximum price"
// @Success 200 {object} pb.GetAllMenusResponse "List of menu items"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menus [get]
func (h *HandlerStruct) GetAllMenusHandler(c *gin.Context) {
	restaurantID := c.Query("restaurant_id")
	name := c.Query("name")
	description := c.Query("description")
	minPriceSt := c.Query("minPrice")
	maxPriceSt := c.Query("maxPrice")
	var minPF, maxPF float64
	if minPriceSt != "" {
		var err error
		minPF, err = strconv.ParseFloat(minPriceSt, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minPrice format"})
			return
		}
	}
	if maxPriceSt != "" {
		var err error
		maxPF, err = strconv.ParseFloat(maxPriceSt, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maxPrice format"})
			return
		}
	}

	req := pb.GetAllMenusRequest{
		RestaurantId: restaurantID,
		Name:         name,
		Description:  description,
		MinPrice:     minPF,
		MaxPrice:     maxPF,
	}

	resp, err := h.Menu.GetAllMenus(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
