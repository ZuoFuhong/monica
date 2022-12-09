package repository

import (
	"context"
	"monica-admin/app/domain/entity"
)

type IServiceRepos interface {

	// CreateService 创建服务
	CreateService(ctx context.Context, record *entity.Service) error

	// GetServiceByID 通过 ID 查询服务
	GetServiceByID(ctx context.Context, sid int) (*entity.Service, error)

	// GetServiceByName 通过名称查询服务
	GetServiceByName(ctx context.Context, name string) (*entity.Service, error)

	// UpdateService 更新服务
	UpdateService(ctx context.Context, record *entity.Service) error

	// ListServiceByPage 查询服务列表
	ListServiceByPage(ctx context.Context, kw, busi string, page, pageSize int) ([]*entity.Service, error)

	// CountServiceByCond 分页查询-服务数量
	CountServiceByCond(ctx context.Context, kw, busi string) (int64, error)
}
