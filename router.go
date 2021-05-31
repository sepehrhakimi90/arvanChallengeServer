package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sepehrhakimi90/arvanChallengeServer/controller"
)

func NewGinRouter(ruleController controller.RuleController) *gin.Engine {
	route := gin.Default()

	rulesRoute := route.Group("/rules")
	rulesRoute.POST("/create", ruleController.CreateRule)
	rulesRoute.GET("/", ruleController.ActiveRules)
	rulesRoute.GET("/delete", ruleController.DeleteRule)
	rulesRoute.GET("/:domain", ruleController.DomainActiveRules)


	return route
}
