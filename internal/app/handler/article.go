package handler

import (
	"net/http"
	"strconv"

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
	Icon     string `json:"icon"`
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

type StepInInput struct {
	ArticleId int    `json:"article_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Num       int    `json:"num"`
	Wifi      bool   `json:"wifi"`
	Question  bool   `json:"question"`
}

func (h *Handler) StepCreate(c *gin.Context) {
	var input StepInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.store.Article().CreateStep(models.Step{
		ArticleId: input.ArticleId,
		Title:     input.Title,
		Num:       input.Num,
		Wifi:      input.Wifi,
		Content:   input.Content,
		Question:  input.Question,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, id)
}

type ImageInInput struct {
	ImageName string `json:"image_name"`
	Image     string `json:"image"`
	StepId    int    `json:"step_id"`
}

func (h *Handler) ImagesCreate(c *gin.Context) {
	var input []ImageInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	for _, image := range input {
		err := h.store.Article().CreateImage(models.Image{
			ImageName: image.ImageName,
			Image:     image.Image,
			StepId:    image.StepId,
		})
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	newResponse(c, http.StatusOK, "OK")
}
