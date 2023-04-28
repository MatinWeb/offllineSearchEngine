package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"searchEngine/internals/searchEngine"
	"time"
)

type Server struct {
	GinEngine    *gin.Engine
	SearchEngine searchEngine.ISearchEngine
	jwtHandler   *JWTHandler
}

type SearchResponse struct {
	Search string `json:"search" binding:"required"`
}

func (s *Server) MetricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("start")
		c.Next()
		fmt.Printf("Next handdler duration : %v", time.Now().Sub(start))
	}
}

func (s *Server) SetRouts() {
	s.GinEngine.POST("/signing", s.jwtHandler.SignInHandler())

	apiGroup := s.GinEngine.Group("/user")
	apiGroup.Use(s.jwtHandler.AuthMiddleware())
	apiGroup.Use(s.MetricMiddleware())
	s.GinEngine.POST("/search", s.SearchHandler())
}

func (s *Server) SearchHandler() gin.HandlerFunc {
	//todo search need error handler
	return func(c *gin.Context) {
		var res SearchResponse
		if err := c.ShouldBindJSON(&res); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}
		search := s.SearchEngine.Search(res.Search)
		c.JSON(http.StatusOK, gin.H{
			"message": search,
		})
	}
}

func NewServer(i searchEngine.ISearchEngine) Server {
	server := Server{
		GinEngine:    gin.Default(),
		SearchEngine: i,
	}
	server.SetRouts()
	return server
}
