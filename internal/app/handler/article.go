package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ArticleAll(c *gin.Context) {
	articles, err := h.store.Article().All()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, articles)
}
