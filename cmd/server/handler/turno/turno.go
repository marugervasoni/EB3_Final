package turno

import (
	"fmt"
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


// Producto godoc
// @Summary Turnos
// @Description Create a new turno
// @Tags Turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos [post]
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

// Producto godoc
// @Summary Turnos
// @Description Get turno by id
// @Tags turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [get]
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

// Producto godoc
// @Summary Turnos
// @Description Update turno by id
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [put]
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

func (h *TurnoHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var tunoReq domain.Turno

		err = ctx.ShouldBindJSON(&tunoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		tunoPatched, err := h.Service.Patch(ctx, tunoReq, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": tunoPatched,
		})
	}
}

func (h *TurnoHandler) HandleDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = h.Service.Delete(ctx, id)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":fmt.Sprintf("turno con id %d eliminado", id),
		})
	}
	
}
