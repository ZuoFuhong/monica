package entity

import "time"

// ServiceNamespace 服务命名空间
type ServiceNamespace struct {
	ID          int        `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ServiceName string     `gorm:"column:service_name;NOT NULL"`                         // 服务名称
	Namespace   string     `gorm:"column:namespace;NOT NULL"`                            // 命名空间
	Token       string     `gorm:"column:token;NOT NULL"`                                // 鉴权凭证
	CreatedAt   *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdatedAt   *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (m *ServiceNamespace) TableName() string {
	return "t_service_namespace"
}
