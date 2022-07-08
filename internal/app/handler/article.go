package handler

import (
	"net/http"
	"strconv"

	_ "encoding/json"

	"github.com/dvd-denis/legko-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetArticles(c *gin.Context) {
	group_id, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var articles []models.Article

	str := c.Query("search")
	if str != "" {
		articles, err = h.store.Article().SeatchArticle(group_id, str)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		articles, err = h.store.Article().GetArticle(group_id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	newResponse(c, http.StatusOK, articles)
}

func (h *Handler) CreateArticle(c *gin.Context) {
	var input models.Article
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.store.Article().CreateArticle(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, id)
}
