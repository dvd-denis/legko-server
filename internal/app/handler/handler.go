package handler

import (
	"github.com/dvd-denis/legko-server/internal/app/store"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store *store.Store
}

func New(st *store.Store) *Handler {
	return &Handler{
		store: st,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	router.GET("/articles", h.Articles)
	router.GET("/steps/:id", h.StepGetId)
	router.GET("/questions", h.Questions)
	router.GET("/article/delete/:id", h.ArticleDelete)
	router.POST("/article", h.ArticleCreate)
	router.POST("/step", h.StepCreate)
	router.POST("/images", h.ImagesCreate)

	return router
}
