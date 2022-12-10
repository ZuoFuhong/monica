package entity

import (
	"time"
)

// ServiceInstance 服务实例表
type ServiceInstance struct {
	ID          int        `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Namespace   string     `gorm:"column:namespace;NOT NULL"`                            // 命名空间
	ServiceName string     `gorm:"column:service_name;NOT NULL"`                         // 服务名
	Isolate     int        `gorm:"column:isolate;default:1;NOT NULL"`                    // 隔离状态：1-不隔离 2-隔离
	IP          string     `gorm:"column:ip;NOT NULL"`                                   // 实例IP
	Port        int        `gorm:"column:port;NOT NULL"`                                 // 端口
	Weight      int        `gorm:"column:weight;default:0;NOT NULL"`                     // 权重
	Metadata    string     `gorm:"column:metadata;NOT NULL"`                             // 元数据
	RenewedAt   *time.Time `gorm:"column:renewed_at"`                                    // 健康更新时间
	CreatedAt   *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdatedAt   *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (m *ServiceInstance) TableName() string {
	return "t_service_instance"
}
