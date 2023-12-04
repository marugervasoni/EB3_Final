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
// @Summary Create a new turno
// @Description Create a new turno
// @Tags Turno
// @Param token header string true "auth token"
// @Accept json
// @Produce json
// @Success 201 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos [post]
func (h *TurnoHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turnoReq domain.Turno
		err := ctx.Bind(&turnoReq) 
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turnoCreated, err := h.Service.Create(ctx, turnoReq)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.Success(ctx, http.StatusCreated, turnoCreated)
	}
}

// Producto godoc
// @Summary Get turno by id
// @Description Get turno by id
// @Tags Turno
// @Param id path int true "id del turno"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [get]
func (h *TurnoHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "ID inválido")
			return
		}

		turnoById, err := h.Service.GetById(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "%s", "turno no encontrado")
			return
		}

		web.Success(ctx, http.StatusOK, turnoById)
	}
}

// Producto godoc
// @Summary Complete turno update by id
// @Description Update all turno fields by id
// @Tags Turno
// @Param token header string true "auth token"
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [put]
func (h *TurnoHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "ID inválidoD")
			return
		}

		var turnoReq domain.Turno
		err = ctx.Bind(&turnoReq); 
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turnoUpdated, err := h.Service.Update(ctx, id, turnoReq)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError,"%s", "internal server error")
			return
		}
		web.Success(ctx, http.StatusOK, turnoUpdated)
	}
}

// Turno godoc
// @Summary Complete or partial turno update by id
// @Description Update all or some turno fields by id
// @Tags Turno
// @Param id path int true "id del turno"
// @Param token header string true "auth token"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [patch]
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
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		tunoPatched, err := h.Service.Patch(ctx, tunoReq, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, tunoPatched)
	}
}

// Turno godoc
// @Summary Delete turno by id
// @Description Delete turno by id
// @Tags Turno
// @Param id path int true "id del turno"
// @Param token header string true "auth token"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [delete]
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

		web.Success(ctx, http.StatusOK, fmt.Sprintf("turno con id %d eliminado", id))
		
	}	
}

// Turno godoc
// @Summary Get turno by DNI
// @Description Get turno by DNI
// @Tags Turno
// @Param dni query int true "dni del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos [get]
func (h *TurnoHandler) HandlerGetByDNI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dni, err := strconv.Atoi(ctx.Query("dni"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Invalid dni")
			return
		}

		turnos, err := h.Service.GetByDNI(ctx, dni)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, turnos)
	}
}