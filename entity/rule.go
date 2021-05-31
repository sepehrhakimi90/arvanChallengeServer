package entity

import (
	"time"

	"gorm.io/gorm"
)

type RuleData struct {
	Domain    string    `gorm:"not null" json:"domain" form:"required"`
	CreatedAt time.Time
	StartTime time.Time `gorm:"not null" json:"start_time" form:"required"`
	Suspect   string    `gorm:"not null" json:"suspect" form:"required ip4_addr"`
	TTL       int       `gorm:"not null" json:"ttl" form:"required gte=1"`
}

type Rule struct {
	ID        int	`json:"id" gorm:"primarykey;autoIncrement"`
	RuleData
	EndTime   int64     `json:"-" gorm:"not null"`
}

func (r *Rule) BeforeSave(tx *gorm.DB) (err error) {
	r.EndTime = getEndTime(r.StartTime, r.TTL)
	return
}

func (r *Rule) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = 0
	return
}

func getEndTime(startTime time.Time, ttl int) int64{
	return startTime.Add(time.Duration(ttl) * time.Second).Unix()
}