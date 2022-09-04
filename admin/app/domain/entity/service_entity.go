package entity

import "time"

type Service struct {
	Id        int64      `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Business  string     `gorm:"column:business" json:"business"`
	Owners    string     `gorm:"column:owners" json:"owners"`
	Remark    string     `gorm:"column:remark" json:"remark"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

func (t *Service) TableName() string {
	return "t_service"
}
