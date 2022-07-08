package handler

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StepGetId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	steps, err := h.store.Article().GetSteps(id)
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
