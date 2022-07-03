package handler

import (
	"net/http"

	_ "encoding/json"

	"github.com/dvd-denis/legko-server/internal/app/models"
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

type ArticleInInput struct {
	Title    string `json:"title"`
	IconName string `json:"icon_name"`
	Icon     []byte `json:"icon"`
	Url      string `json:"url"`
	Color    string `json:"color"`
}

func (h *Handler) ArticleCreate(c *gin.Context) {
	var input ArticleInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()+"1")
		return
	}

	id, err := h.store.Article().CreateArticle(models.Article{
		Title:    input.Title,
		IconName: input.IconName,
		Icon:     input.Icon,
		Url:      input.Url,
		Color:    input.Color,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, id)
}

func (h *Handler) StepCreate(c *gin.Context) {
	var input models.Step

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.store.Article().CreateStep(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, id)
}

func (h *Handler) ImagesCreate(c *gin.Context) {
	var input []models.Image

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.store.Article().CreateImages(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, "OK")
}
