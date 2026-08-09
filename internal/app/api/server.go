package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	posts_client "github.com/koliader/tellmi-sdk/clients/posts"
	users_client "github.com/koliader/tellmi-sdk/clients/users"
	"github.com/koliader/tellmi-gateway/internal/config"
	"github.com/koliader/tellmi-gateway/internal/lib/middleware"
	"github.com/koliader/tellmi-sdk/token"
)

var timeout = time.Second * 60

type Server struct {
	router      *gin.Engine
	config      config.Config
	usersClient users_client.Client
	postsClient posts_client.Client
	tokenMaker  token.Maker
}

func NewServer(config config.Config) (*Server, error) {
	usersClient := users_client.NewClient(config.UsersServiceAddress)
	postsClient := posts_client.NewClient(config.PostsServiceAddress)

	tokenMaker, err := token.NewJWTMaker(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("error to create token maker: %v", err)
	}

	server := Server{
		config:      config,
		usersClient: *usersClient,
		postsClient: *postsClient,
		tokenMaker:  tokenMaker,
	}
	server.setupRouter()
	return &server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.ContextWithFallback = true

	c := cors.New(cors.Config{
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization", "Refresh-Token"},
	})

	router.Use(c, otelgin.Middleware("tellmi-gateway"), requestLogger())

	m := middleware.NewMiddleware(s.tokenMaker)
	authRoutes := router.Group("/").Use(m.AuthMiddleware())
	adminRoutes := router.Group("/").Use(m.AdminMiddleware())

	// auth
	router.POST("/auth/register", s.register)
	router.POST("/auth/login", s.login)
	router.POST("/auth/refresh", s.refresh)

	// users
	router.GET("/users/:id", s.getUserById)
	adminRoutes.GET("/users", s.listUsers)
	authRoutes.PUT("/users", s.updateUser)

	// posts
	authRoutes.POST("/posts", s.createPost)
	router.GET("/posts", s.listPosts)
	router.GET("/posts/:id", s.getPostById)
	authRoutes.PUT("/posts", s.editPost)
	authRoutes.DELETE("/posts/:id", s.deletePost)

	// categories
	adminRoutes.POST("/categories", s.createCategory)
	router.GET("/categories", s.listCategories)
	adminRoutes.PUT("/categories", s.editCategory)

	// comments
	authRoutes.POST("/comments", s.createComment)
	router.GET("/comments/post/:id", s.listCommentsByPost)
	authRoutes.PUT("/comments", s.editComment)
	authRoutes.DELETE("/comments/:id", s.deleteComment)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
