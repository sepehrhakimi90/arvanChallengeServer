package service

import (
	"github.com/sepehrhakimi90/arvanChallengeServer/entity"
	"github.com/sepehrhakimi90/arvanChallengeServer/repository"
)

type RuleService interface {
	Create(*entity.Rule) (*entity.Rule, error)
	FindActiveRules() ([]entity.Rule, error)
	FindDomainActiveRules(string) ([]entity.Rule, error)
	Delete(int) error
}

type ruleService struct {
	ruleRepo repository.RuleRepository
}

func NewRuleService(ruleRepo repository.RuleRepository) RuleService {
	return &ruleService{ruleRepo:ruleRepo}
}

func (r *ruleService) Create(rule *entity.Rule) (*entity.Rule, error) {
	return r.ruleRepo.Save(rule)
}

func (r *ruleService) FindActiveRules() ([]entity.Rule, error) {
	return r.ruleRepo.GetActiveRules()
}

func (r *ruleService) FindDomainActiveRules(domain string) ([]entity.Rule, error) {
	return r.ruleRepo.GetDomainActiveRules(domain)
}

func (r *ruleService) Delete(id int) error {
	return r.ruleRepo.DeleteById(id)
}



