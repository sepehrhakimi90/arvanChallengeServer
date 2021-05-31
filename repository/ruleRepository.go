package repository

import (
	"github.com/sepehrhakimi90/arvanChallengeServer/entity"
)

type RuleRepository interface {
	Save(*entity.Rule) (*entity.Rule, error)
	GetActiveRules() ([]entity.Rule, error)
	GetDomainActiveRules(string) ([]entity.Rule, error)
	DeleteById(int) error
}