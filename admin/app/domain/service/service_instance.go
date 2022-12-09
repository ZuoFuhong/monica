package service

import (
	"context"
	"errors"
	"monica-admin/app/domain/entity"
	"monica-admin/app/domain/repository"
	"time"
)

type IServiceInstance interface {
	// SaveInstance 更新服务实例
	SaveInstance(ctx context.Context, ins *entity.ServiceInstance) error

	// ListInstanceByPage 查询服务实例列表
	ListInstanceByPage(ctx context.Context, ns string, serviceId int, ip string, port, page, pageSize int) ([]*entity.ServiceInstance, int64, error)

	// DeleteInstance 删除服务实例
	DeleteInstance(ctx context.Context, ns string, serviceID int, ip string) error
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
	insDO, err := s.repos.GetInstanceByIP(ctx, ins.Namespace, ins.ServiceID, ins.IP)
	if err != nil {
		return err
	}
	// 新增实例
	if ins.ID == 0 {
		if insDO.ID != 0 {
			return errors.New("请勿添加重复的实例IP")
		}
		serDO, err := s.servRepos.GetServiceByID(ctx, ins.ServiceID)
		if err != nil {
			return err
		}
		if serDO.ID == 0 {
			return errors.New("无效的 ServiceID")
		}
		return s.repos.CreateInstance(ctx, ins)
	}
	// 更新实例
	if insDO.ID == 0 {
		return errors.New("实例不存在")
	}
	if insDO.ID != ins.ID {
		return errors.New("请勿添加重复的实例IP")
	}
	return s.repos.UpdateInstance(ctx, ins)
}

func (s *Instance) ListInstanceByPage(ctx context.Context, ns string, serviceId int, ip string, port, page, pageSize int) ([]*entity.ServiceInstance, int64, error) {
	list, err := s.repos.ListInstanceByPage(ctx, ns, serviceId, ip, port, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountInstanceByCond(ctx, ns, serviceId, ip, port)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (s *Instance) DeleteInstance(ctx context.Context, ns string, serviceID int, ip string) error {
	insDO, err := s.repos.GetInstanceByIP(ctx, ns, serviceID, ip)
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
