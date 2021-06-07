package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sepehrhakimi90/arvanChallengeServer/controller"
)

func NewGinRouter(ruleController controller.RuleController) *gin.Engine {
	route := gin.Default()

	route.Static("/css", "./templates/css")
	route.Static("/js", "./templates/js")
	route.LoadHTMLGlob("templates/*.html")

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	rulesRoute := route.Group("/rules")
	rulesRoute.POST("/create", ruleController.CreateRule)
	rulesRoute.GET("/", ruleController.ActiveRules)
	rulesRoute.GET("/delete", ruleController.DeleteRule)
	rulesRoute.GET("/:domain", ruleController.DomainActiveRules)


	return route
}
