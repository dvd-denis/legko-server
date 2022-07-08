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
	router.Use(CheckKey())

	router.GET("/groups", h.GetGroups)
	router.GET("/articles/:group_id", h.GetArticles)
	router.GET("/steps/:article_id", h.GetSteps)
	router.POST("/group/delete/:id", h.DeleteGroup)
	router.POST("/group", h.CreateGroup)
	router.POST("/article", h.CreateArticle)
	router.POST("/step", h.CreateStep)
	router.POST("/images", h.CreateImages)

	return router
}
