package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"balance": 1000})
}

func GetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"transactions": []string{}})
}

func GenerateMonthlyStatement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"statement": []string{}})
}

func PerformDeposit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func PerformWithdrawal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
