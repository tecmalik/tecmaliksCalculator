package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func postEvaluate(expression *gin.Context) string {
	expression.IndentedJSON(http.StatusOK, expression)
	return ""
}
