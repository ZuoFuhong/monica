package repository

import (
	"context"
	"monica/app/domain/entity"
)

type IServiceRepos interface {

	// CreateService 创建服务
	CreateService(ctx context.Context, record *entity.Service) error

	// GetServiceByNames 批量查询服务
	GetServiceByNames(ctx context.Context, snameList []string) ([]*entity.Service, error)

	// GetServiceByName 通过名称查询服务
	GetServiceByName(ctx context.Context, name string) (*entity.Service, error)

	// UpdateService 更新服务
	UpdateService(ctx context.Context, record *entity.Service) error

	// ListServiceByPage 查询服务列表
	ListServiceByPage(ctx context.Context, kw, busi string, page, pageSize int) ([]*entity.Service, error)

	// CountServiceByCond 分页查询-服务数量
	CountServiceByCond(ctx context.Context, kw, busi string) (int64, error)

	// BatchAddServiceNS 批量增加命名空间
	BatchAddServiceNS(ctx context.Context, nsList []*entity.ServiceNamespace) error

	// ListServiceNSByPage 分页查询-服务命名空间
	ListServiceNSByPage(ctx context.Context, kw, ip string, page, pageSize int) ([]*entity.ServiceNamespace, error)

	// CountServiceNSByCond 分页查询-服务命名空间数量
	CountServiceNSByCond(ctx context.Context, kw, ip string) (int64, error)

	// GetServiceNSByName 查询服务命名空间
	GetServiceNSByName(ctx context.Context, ns, sname string) (*entity.ServiceNamespace, error)
}
