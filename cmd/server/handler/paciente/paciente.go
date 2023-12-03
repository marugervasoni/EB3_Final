package paciente

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jum8/EBE3_Final.git/internal/domain"
	"github.com/jum8/EBE3_Final.git/internal/paciente"
	"github.com/jum8/EBE3_Final.git/pkg/web"
)

type PacienteHandler struct {
	Service paciente.Service
}

func NewPacienteHandler(service paciente.Service) *PacienteHandler {
	return &PacienteHandler{
		Service: service,
	}
}



func (h *PacienteHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/pacientes", h.HandlerGetAll())
	router.GET("/paciente/:id", h.HandlerGetById())
	router.POST("/paciente", h.HandlerCreate())
	router.PUT("/paciente/:id", h.HandlerUpdate())
	router.DELETE("/paciente/:id", h.HandlerDelete())
	router.PATCH("/paciente/:id", h.HandlerPatch())
}



// Paciente godoc
// @Summary Paciente
// @Description get all pacientes
// @Tags Paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes [get]
func (h *PacienteHandler) HandlerGetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := h.Service.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": pacientes})
	}
}


// Paciente godoc
// @Summary Get paciente by id
// @Description Get paciente by id
// @Tags Paciente
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [get]
func (h *PacienteHandler) HandlerGetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		paciente, err := h.Service.GetById(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": paciente})
	}
}


// Paciente godoc
// @Summary Create a new paciente
// @Description Create a new paciente
// @Tags Paciente
// @Param token header string true "auth token"
// @Accept json
// @Produce json
// @Success 201 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /paciente [post]
func (h *PacienteHandler) HandlerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente
		if err := c.Bind(&paciente); err != nil {
			web.Error(c, http.StatusBadRequest, "Bad request")
			return
		}

		newPaciente, err := h.Service.Create(c, paciente)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, gin.H{"data": newPaciente})
	}
}


// Paciente godoc
// @Summary Complete paciente update by id
// @Description Update all paciente fields by id
// @Tags Paciente
// @Param token header string true "auth token"
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [put]
func (h *PacienteHandler) HandlerUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		var paciente domain.Paciente
		if err := c.Bind(&paciente); err != nil {
			web.Error(c, http.StatusBadRequest, "Bad request binding")
			return
		}

		updatedPaciente, err := h.Service.Update(c, id, paciente)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": updatedPaciente})
	}
}


// Paciente godoc
// @Summary Delete paciente by id
// @Description Delete paciente by id
// @Tags Paciente
// @Param token header string true "auth token"
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [delete]
func (h *PacienteHandler) HandlerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		err = h.Service.Delete(c, id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"message": "Patient deleted successfully"})
	}
}


// Paciente godoc
// @Summary Complete or partial paciente update by id
// @Description Update all or some paciente fields by id
// @Tags Paciente
// @Param token header string true "auth token"
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [patch]
func (h *PacienteHandler) HandlerPatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		var paciente domain.Paciente
		if err := c.Bind(&paciente); err != nil {
			web.Error(c, http.StatusBadRequest, "Bad request binding")
			return
		}

		patchedPaciente, err := h.Service.Patch(c, id, paciente)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": patchedPaciente})
	}
}
