package turno

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jum8/EBE3_Final.git/internal/domain"
	"github.com/jum8/EBE3_Final.git/internal/turno"
	"github.com/jum8/EBE3_Final.git/pkg/web"
)

type TurnoHandler struct {
	Service turno.Service
}

func NewTurnoHandler(service turno.Service) *TurnoHandler {
	return &TurnoHandler{
		Service: service,
	}
}

func (h *TurnoHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/turno", h.HandlerCreate())
	router.GET("/turno/:id", h.HandlerGetById())
	router.PUT("/turno/:id", h.HandlerUpdate())
}

func (h *TurnoHandler) HandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno
		if err := c.Bind(&turno); err != nil {
			web.Error(c, http.StatusBadRequest, "Bad request")
			return
		}

		newTurno, err := h.Service.Create(c, turno)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, gin.H{"data": newTurno})
	}
}

func (h *TurnoHandler) HandlerGetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		turno, err := h.Service.GetById(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": turno})
	}
}

func (h *TurnoHandler) HandlerUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		var turno domain.Turno
		if err := c.Bind(&turno); err != nil {
			web.Error(c, http.StatusBadRequest, "Bad request binding")
			return
		}

		updatedTurno, err := h.Service.Update(c, id, turno)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": updatedTurno})
	}
}
