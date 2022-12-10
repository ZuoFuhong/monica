package service

import (
	"context"
	"errors"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"time"
)

type IServiceInstance interface {
	// SaveInstance 更新服务实例
	SaveInstance(ctx context.Context, ins *entity.ServiceInstance) error

	// ListInstanceByPage 查询服务实例列表
	ListInstanceByPage(ctx context.Context, ns, serviceName string, ip string, port, page, pageSize int) ([]*entity.ServiceInstance, int64, error)

	// DeleteInstance 删除服务实例
	DeleteInstance(ctx context.Context, ns, serviceName string, ip string) error
}

type Instance struct {
	servRepos repository.IServiceRepos
	repos     repository.IServiceInstanceRepos
}

func NewServiceInstance(servRepos repository.IServiceRepos, repos repository.IServiceInstanceRepos) IServiceInstance {
	return &Instance{
		servRepos: servRepos,
		repos:     repos,
	}
}

func (s *Instance) SaveInstance(ctx context.Context, ins *entity.ServiceInstance) error {
	insDO, err := s.repos.GetInstanceByIP(ctx, ins.Namespace, ins.ServiceName, ins.IP)
	if err != nil {
		return err
	}
	// 新增实例
	if ins.ID == 0 {
		if insDO.ID != 0 {
			return errors.New("实例IP已存在")
		}
		serDO, err := s.servRepos.GetServiceByName(ctx, ins.ServiceName)
		if err != nil {
			return err
		}
		if serDO.ID == 0 {
			return errors.New("无效的服务")
		}
		return s.repos.CreateInstance(ctx, ins)
	}
	// 更新实例
	if insDO.ID == 0 {
		return errors.New("实例不存在")
	}
	if ins.ID != insDO.ID {
		return errors.New("实例IP已存在")
	}
	// 更新字段
	insDO.Isolate = ins.Isolate
	insDO.Weight = ins.Weight
	insDO.Metadata = ins.Metadata
	return s.repos.UpdateInstance(ctx, insDO)
}

func (s *Instance) ListInstanceByPage(ctx context.Context, ns, serviceName string, ip string, port, page, pageSize int) ([]*entity.ServiceInstance, int64, error) {
	list, err := s.repos.ListInstanceByPage(ctx, ns, serviceName, ip, port, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountInstanceByCond(ctx, ns, serviceName, ip, port)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (s *Instance) DeleteInstance(ctx context.Context, ns, serviceName string, ip string) error {
	insDO, err := s.repos.GetInstanceByIP(ctx, ns, serviceName, ip)
	if err != nil {
		return err
	}
	if insDO.ID == 0 {
		return errors.New("服务实例不存在")
	}
	now := time.Now()
	insDO.DeletedAt = &now
	return s.repos.UpdateInstance(ctx, insDO)
}
