package handler

import (
	"net/http"
	"strconv"

	"github.com/dvd-denis/legko-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetGroups(c *gin.Context) {
	model := c.Query("model")
	if model == "" {
		model = "default"
	}

	articles, err := h.store.Article().GetGroups(model)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, articles)
}

func (h *Handler) CreateGroup(c *gin.Context) {
	var input models.Group
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.store.Article().CreateGroup(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, id)
}

func (h *Handler) DeleteGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.store.Article().DeleteGroup(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, "OK")
}
