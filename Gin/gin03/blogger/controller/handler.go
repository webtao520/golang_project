package controller

import (
	_"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandle (c *gin.Context){
c.HTML(http.StatusOK,"views/index.html",gin.H{})			
}