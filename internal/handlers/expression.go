package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ItzB1ack/CalculatorYL2/internal/models"
)

var expressions = make(map[string]models.Expression)

func CalculateExpression(c *gin.Context) {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid data"})
		return
	}

	expressionID := "some-unique-id"
	expressions[expressionID] = models.Expression{
		ID:     expressionID,
		Status: "pending",
	}

	c.JSON(http.StatusCreated, gin.H{"id": expressionID})
}
