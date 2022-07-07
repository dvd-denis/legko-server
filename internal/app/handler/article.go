package handler

import (
	"net/http"
	"strconv"

	_ "encoding/json"

	"github.com/dvd-denis/legko-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Articles(c *gin.Context) {
	model := c.Query("model")
	if model == "" {
		model = "default"
	}

	articles, err := h.store.Article().All(model, false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, articles)
}

func (h *Handler) Questions(c *gin.Context) {
	model := c.Query("model")
	if model == "" {
		model = "default"
	}

	questions, err := h.store.Article().All(model, true)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, questions)
}

func (h *Handler) ArticleCreate(c *gin.Context) {
	var input models.Article
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()+"1")
		return
	}

	id, err := h.store.Article().CreateArticle(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, id)
}

func (h *Handler) ArticleDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.store.Article().DeleteArticle(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, "OK")
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
	for _, image := range input {
		err := h.store.Article().CreateImage(image)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	newResponse(c, http.StatusOK, "OK")
}
