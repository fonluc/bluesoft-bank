package main

import (
	"github.com/fonluc/bluesoft-bank/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rotas para a API
	r.GET("/balance/:id", services.GetBalance)
	r.GET("/transactions/:id", services.GetTransactions)
	r.GET("/statement/:id", services.GenerateMonthlyStatement)
	r.POST("/deposit", services.PerformDeposit)
	r.POST("/withdraw", services.PerformWithdrawal)

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
