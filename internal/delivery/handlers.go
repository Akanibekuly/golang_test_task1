package delivery

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Akanibekuly/golang_test_task1.git/internal"
	"github.com/Akanibekuly/golang_test_task1.git/internal/models"
	"github.com/Akanibekuly/golang_test_task1.git/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cs internal.CityService
}

func NewHandlers(db *sql.DB) *Handler {
	cs := service.NewCityService(db)
	return &Handler{cs: cs}
}

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.cs.GetCities()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseWithData{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseWithData{
		Status: "OK",
		Data:   cities,
	})
}
func (h *Handler) GetCity(c *gin.Context) {
	idStr, exists := c.Params.Get("id")
	if !exists {
		log.Println("id doesn't exists in request")
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "bad request: id is empty",
		})
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("id: err")
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "invalid input: id should be integer",
		})
		return
	}

	city, err := h.cs.GetCity(id)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, &models.ResponseWithData{
				Status:  "error",
				Message: fmt.Sprintf("city with id %d not found", id),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &models.ResponseWithData{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &models.ResponseWithData{
		Status: "OK",
		Data:   city,
	})
}

func (h *Handler) CreateCity(c *gin.Context) {
	var city models.City
	err := c.ShouldBindJSON(&city)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "bad reques: could't  serialize json into city",
		})
		return
	}

	id, err := h.cs.CreateCity(&city)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &models.ResponseWithData{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, &models.ResponseWithData{
		Status:  "OK",
		Message: fmt.Sprintf("city with id %d successfully created", id),
	})
}

func (h *Handler) DeleteCity(c *gin.Context) {
	idStr, exists := c.Params.Get("id")
	if !exists {
		log.Println("id doesn't exists in request")
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "bad request: id is empty",
		})
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("id: err")
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "invalid input: id should be integer",
		})
		return
	}
	err = h.cs.DeleteCity(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &models.ResponseWithData{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &models.ResponseWithData{
		Status:  "OK",
		Message: fmt.Sprintf("city with id %d succesfully deleted", id),
	})
}

func (h *Handler) UpdateCity(c *gin.Context) {
	idStr, exists := c.Params.Get("id")
	if !exists {
		log.Println("id doesn't exists in request")
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "bad request: id is empty",
		})
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("id: err")
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "invalid input: id should be integer",
		})
		return
	}

	var city models.City
	err = c.ShouldBindJSON(&city)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.ResponseWithData{
			Status:  "error",
			Message: "bad reques: could't  serialize json into city",
		})
		return
	}

	err = h.cs.UpdateCity(id, &city)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &models.ResponseWithData{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &models.ResponseWithData{
		Status:  "OK",
		Message: fmt.Sprintf("city with id %d succesfully deleted", id),
	})
}
