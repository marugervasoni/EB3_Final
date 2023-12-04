package odontologo

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jum8/EBE3_Final.git/internal/domain"
	"github.com/jum8/EBE3_Final.git/internal/odontologo"
	"github.com/jum8/EBE3_Final.git/pkg/web"
)

type Controller struct {
	service odontologo.Service
}

func NewControllerOdontologo(service odontologo.Repository) *Controller {
	return &Controller{
		service: service,
	}
}

// Odontologo godoc
// @Summary Get all odontologos
// @Description Get all odontologos
// @Tags Odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := c.service.GetAll(ctx)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, odontologos)
	}
}

// Odontologo godoc
// @Summary Get odontologo by id
// @Description Get odontologo by id
// @Tags Odontologo
// @Param id path int true "id del odontologo"
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/{id} [get]
func (c *Controller) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Invalid ID")
			return
		}

		odontologoById, err := c.service.GetById(ctx, id)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, odontologoById)
	}
}

// Odontologo godoc
// @Summary Create a new odontologo
// @Description Create a new odontologo
// @Tags Odontologo
// @Param token header string true "auth token"
// @Param payload body domain.Odontologo true "Odontologo"
// @Accept json
// @Produce json
// @Success 201 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var odontologoReq domain.Odontologo

		err := ctx.Bind(&odontologoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		odontologoCreated, err := c.service.Create(ctx, odontologoReq)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusCreated, odontologoCreated)
	}
}

// Odontologo godoc
// @Summary Complete odontologo update by id
// @Description Update all odontologo fields by id
// @Tags Odontologo
// @Param token header string true "auth token"
// @Param id path int true "id del odontologo"
// @Param payload body domain.Odontologo true "Odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/{id} [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Invalid ID")
			return
		}

		var odontologoReq domain.Odontologo

		err = ctx.Bind(&odontologoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		odontologoUpdated, err := c.service.Update(ctx, id, odontologoReq)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, odontologoUpdated)
	}
}

// Odontologo godoc
// @Summary Delete odontologo by id
// @Description Delete odontologo by id
// @Tags Odontologo
// @Param token header string true "auth token"
// @Param id path int true "id del odontologo"
// @Produce json
// @Success 200 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/{id} [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Invalid ID")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, fmt.Sprintf("odontologo con id %d eliminado", id))
	}
}

// Odontologo godoc
// @Summary Complete or partial odontologo update by id
// @Description Update all or some odontologo fields by id
// @Tags Odontologo
// @Param token header string true "auth token"
// @Param id path int true "id del odontologo"
// @Param payload body domain.Odontologo true "Odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/{id} [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Invalid ID")
			return
		}

		var odontologoReq domain.Odontologo

		err = ctx.Bind(&odontologoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		odontologoPatched, err := c.service.Patch(ctx, id, odontologoReq)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, odontologoPatched)
	}
}

func errorHandler(ctx *gin.Context, err error) {
	if errors.Is(err, odontologo.ErrInvalidAttributes) {
		web.Error(ctx, http.StatusBadRequest, "%s", "atributos de odontologo incorrectos")
	} else if errors.Is(err, odontologo.ErrDuplicateMatricula) {
		web.Error(ctx, http.StatusBadRequest, "%s", "ya existe un odontologo con la matricula ingresada")
	} else if errors.Is(err, odontologo.ErrNotFound) {
		web.Error(ctx, http.StatusNotFound, "%s", "odontologo no encontrado")
	} else {
		web.InternalServerError(ctx)
	}
}
