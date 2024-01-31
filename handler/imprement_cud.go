package handler

import (
	"database/sql"
	"github.com/easywalk/go-restful"
	"github.com/easywalk/go-restful/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type SimplyHandler[T easywalk.SimplyEntityInterface] struct {
	Svc service.SimplyServiceInterface[T]
}

func (h SimplyHandler[T]) Create(c *gin.Context) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.Svc.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, created)
}

func (h SimplyHandler[T]) Update(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Svc.UpdateByID(uuidID, req)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h SimplyHandler[T]) Delete(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Svc.DeleteByID(uuidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == 0 {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"result": "success"})
}

func (h SimplyHandler[T]) Read(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Svc.ReadByID(uuidID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == nil {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h SimplyHandler[T]) FindAll(c *gin.Context) {
	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entities, err := h.Svc.FindAll(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if entities == nil || len(entities) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusOK, entities)
}
