package repository

import (
	"context"
	"monica/app/domain/entity"
)

type IServiceInstanceRepos interface {
	// CreateInstance 创建服务实例
	CreateInstance(ctx context.Context, record *entity.ServiceInstance) error

	// GetInstanceByIP 通过 IP 查询实例
	GetInstanceByIP(ctx context.Context, ns, serviceName string, ip string) (*entity.ServiceInstance, error)

	// GetAllInstance 查询所有服务实例
	GetAllInstance(ctx context.Context, ns, sname string) ([]*entity.ServiceInstance, error)

	// UpdateInstance 更新服务实例
	UpdateInstance(ctx context.Context, ins *entity.ServiceInstance) error

	// ListInstanceByPage 查询服务实例列表
	ListInstanceByPage(ctx context.Context, ns, serviceName string, ip string, port int, page, pageSize int) ([]*entity.ServiceInstance, error)

	// CountInstanceByCond 分页查询-服务实例数量
	CountInstanceByCond(ctx context.Context, ns, serviceName string, ip string, port int) (int64, error)
}
