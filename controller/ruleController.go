package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sepehrhakimi90/arvanChallengeServer/entity"
	"github.com/sepehrhakimi90/arvanChallengeServer/service"
	"github.com/sepehrhakimi90/arvanChallengeServer/utils"
)

type RuleController interface {
	CreateRule(c *gin.Context)
	ActiveRules(c *gin.Context)
	DomainActiveRules(c *gin.Context)
	DeleteRule(c *gin.Context)
}

type ruleController struct {
	ruleService service.RuleService
	publisher service.Publisher
}

func NewController(ruleService service.RuleService, publisher service.Publisher) RuleController{
	return &ruleController{
		ruleService: ruleService,
		publisher: publisher,
	}
}

func (r *ruleController) CreateRule(c *gin.Context) {
	rule := &entity.Rule{}
	if err := c.ShouldBindJSON(rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		utils.LogError("ruleController", "CreateRule", err)
		return
	}

	if rule.StartTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start time cannot be in the past"})
		return
	}

	err := r.publisher.Publish(rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something goes wrong in our end"})
		utils.LogError("ruleController", "CreateRule", err)
		return
	}

	rule, err = r.ruleService.Create(rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something goes wrong in our end"})
		utils.LogError("ruleController", "CreateRule", err)
		return
	}
	c.JSON(http.StatusOK, rule)
}

func (r *ruleController) ActiveRules(c *gin.Context) {
	rules, err := r.ruleService.FindActiveRules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something goes wrong in our end"})
		utils.LogError("ruleController", "ActiveRules", err)
		return
	}
	c.JSON(http.StatusOK, rules)
}

func (r *ruleController) DomainActiveRules(c *gin.Context) {
	domain := c.Param("domain")
	rules, err := r.ruleService.FindDomainActiveRules(domain)
	if err != nil {
		utils.LogError("ruleController", "DomainActiveRules", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something goes wrong in our end"})
		return
	}
	c.JSON(http.StatusOK, rules)
}

func (r *ruleController) DeleteRule(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		utils.LogError("ruleController", "DeleteRule", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
		return
	}
	err = r.ruleService.Delete(int(id))
	if err != nil {
		utils.LogError("ruleController", "DeleteRule", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something goes wrong in our end"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
