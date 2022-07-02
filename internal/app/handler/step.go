package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetId(c *gin.Context) {
	steps, err := h.store.Article().GetSteps(c.GetInt("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, steps)
}
