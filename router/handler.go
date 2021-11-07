package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/saltbo/bogo/router/types"
)

type Handler struct {
	Resource types.Resource
}

func (h *Handler) Get(c *gin.Context) {
	name := c.Param("name")
	resource, err := h.Resource.Get(c, name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *Handler) Find(c *gin.Context) {
	query := h.Resource.Query()
	if err := c.ShouldBindQuery(query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resources, total, err := h.Resource.Find(c, query)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(total)
	c.JSON(http.StatusOK, resources)
}

func (h *Handler) Create(c *gin.Context) {
	resource := h.Resource.Model()
	if err := c.ShouldBind(resource); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.Resource.Create(c, resource); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *Handler) Put(c *gin.Context) {
	resource := h.Resource.Model()
	if err := c.ShouldBind(resource); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.Resource.Update(c, resource); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *Handler) Patch(c types.Context) {
	panic("implement me")
}

func (h *Handler) Delete(c *gin.Context) {
	name := c.Param("name")
	resource, err := h.Resource.Get(c, name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.Resource.Delete(c, resource); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}
