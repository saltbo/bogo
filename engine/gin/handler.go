package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/saltbo/bogo/engine/types"
	"net/http"
)

type Handler struct {
	Resource types.Resource
}

func (h *Handler) Get(c *gin.Context) {
	name := c.Param("name")
	resource, err := h.Resource.Get(name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *Handler) Find(c *gin.Context) {
	query := h.Resource.Query()
	if err := c.ShouldBind(&query); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resources, err := h.Resource.Find(query)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resources)
}

func (h *Handler) Create(c *gin.Context) {
	resource := h.Resource.Model()
	if err := c.ShouldBind(&resource); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.Resource.Create(resource); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *Handler) Put(c *gin.Context) {
	resource := h.Resource.Model()
	if err := c.ShouldBind(&resource); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.Resource.Update(resource); err != nil {
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
	resource, err := h.Resource.Get(name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.Resource.Delete(resource); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resource)
}
