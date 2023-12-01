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
