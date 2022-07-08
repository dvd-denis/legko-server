package handler

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/dvd-denis/legko-server/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSteps(c *gin.Context) {
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	steps, err := h.store.Article().GetSteps(article_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Sorting steps on num

	sort.Slice(steps, func(i, j int) bool {
		return steps[i].Num < steps[j].Num
	})

	newResponse(c, http.StatusOK, steps)
}

func (h *Handler) CreateStep(c *gin.Context) {
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

func (h *Handler) CreateImages(c *gin.Context) {
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
