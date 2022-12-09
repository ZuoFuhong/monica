package entity

import "time"

// Service 服务表
type Service struct {
	ID        int        `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name      string     `gorm:"column:name;NOT NULL"`                                 // 服务名称
	Business  string     `gorm:"column:business;NOT NULL"`                             // 业务名称
	Owners    string     `gorm:"column:owners;NOT NULL"`                               // 负责人
	Remark    string     `gorm:"column:remark;NOT NULL"`                               // 备注
	CreatedAt *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (m *Service) TableName() string {
	return "t_service"
}
