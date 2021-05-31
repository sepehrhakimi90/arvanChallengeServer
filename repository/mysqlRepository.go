package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/sepehrhakimi90/arvanChallengeServer/entity"
	"github.com/sepehrhakimi90/arvanChallengeServer/utils"
)

type mysqlRuleRepo struct {
	db *gorm.DB
}

func NewMysqlRuleRepository(db *gorm.DB) RuleRepository{
	repo := mysqlRuleRepo{db: db}
	return &repo
}

func (m *mysqlRuleRepo) Save(rule *entity.Rule) (*entity.Rule, error) {
	result := m.db.Create(&rule)
	if result.Error != nil {
		utils.LogError("mysqlRuleRepository", "Save", result.Error)
		return nil, result.Error
	}
	return rule, nil
}

func (m *mysqlRuleRepo) GetActiveRules() ([]entity.Rule, error) {
	rules := make([]entity.Rule, 0)
	result := m.db.Where("end_time > ?", time.Now().Unix()).Find(&rules)
	if result.Error != nil {
		utils.LogError("mysqlRuleRepository", "GetExpiredRules", result.Error)
		return nil, result.Error
	}
	return rules, nil
}

func (m *mysqlRuleRepo) GetDomainActiveRules(domain string) ([]entity.Rule, error) {
	rules := make([]entity.Rule, 0)
	result := m.db.Where("end_time > ? AND domain = ?", time.Now().Unix(), domain).Find(&rules)
	if result.Error != nil {
		utils.LogError("mysqlRuleRepository", "GetExpiredRules", result.Error)
		return nil, result.Error
	}
	return rules, nil
}

func (m *mysqlRuleRepo) DeleteById(id int) error {
	rule := &entity.Rule{ID: id}
	result := m.db.Delete(rule)
	if result.Error != nil {
		utils.LogError("mysqlRuleRepository", "DeleteById", result.Error)
		return result.Error
	}
	return nil
}