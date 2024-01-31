package handler

import (
	"github.com/easywalk/go-restful/easywalk"
	"github.com/easywalk/go-restful/easywalk/service"
	"github.com/gin-gonic/gin"
)

type GenericHandlerInterface[T any] interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	Read(c *gin.Context)
	FindAll(c *gin.Context)
}

func NewHandler[T easywalk.SimplyEntityInterface](group *gin.RouterGroup, svc service.SimplyServiceInterface[T]) GenericHandlerInterface[T] {
	handlers := &SimplyHandler[T]{Svc: svc}

	group.POST("", handlers.Create)
	group.PATCH(":id", handlers.Update)
	group.DELETE(":id", handlers.Delete)

	group.GET(":id", handlers.Read)
	group.GET("", handlers.FindAll)

	return handlers
}
