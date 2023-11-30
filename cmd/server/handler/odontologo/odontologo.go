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

func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := c.service.GetAll(ctx)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologos,
		})
	}
}

func (c *Controller) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		odontologoById, err := c.service.GetById(ctx, id)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologoById,
		})
	}
}

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

		web.Success(ctx, http.StatusCreated, gin.H{
			"data": odontologoCreated,
		})
	}
}

func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
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

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologoUpdated,
		})
	}
}

func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"message": fmt.Sprintf("odontologo con id %d eliminado", id),
		})
	}
}

func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
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

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologoPatched,
		})
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
