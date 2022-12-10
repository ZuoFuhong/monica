package service

import (
	"context"
	"errors"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"monica/consts"
	"time"
)

type IRegisterService interface {
	Register(ctx context.Context, ns, sname string, insNode *entity.InstanceNode) error

	Renew(ctx context.Context, ns, sname, ip string) error

	Cancel(ctx context.Context, ns, sname, ip string) error
}

type RegisterService struct {
	repos repository.IServiceInstanceRepos
}

func NewRegisterService(insRepos repository.IServiceInstanceRepos) IRegisterService {
	return &RegisterService{
		repos: insRepos,
	}
}

func (s *RegisterService) Register(ctx context.Context, ns, sname string, insNode *entity.InstanceNode) error {
	ins, err := s.repos.GetInstanceByIP(ctx, ns, sname, insNode.IP)
	if err != nil {
		return err
	}
	now := time.Now()
	if ins.ID == 0 {
		ins := &entity.ServiceInstance{
			Namespace:   ns,
			ServiceName: sname,
			Isolate:     consts.Normal,
			IP:          insNode.IP,
			Port:        insNode.Port,
			Weight:      insNode.Weight,
			Metadata:    insNode.Metadata,
			RenewedAt:   &now,
		}
		return s.repos.CreateInstance(ctx, ins)
	}
	// 更新
	ins.RenewedAt = &now
	return s.repos.UpdateInstance(ctx, ins)
}

func (s *RegisterService) Renew(ctx context.Context, ns, sname, ip string) error {
	ins, err := s.repos.GetInstanceByIP(ctx, ns, sname, ip)
	if err != nil {
		return err
	}
	if ins.ID == 0 {
		return errors.New("not found service instance")
	}
	now := time.Now()
	ins.RenewedAt = &now
	return s.repos.UpdateInstance(ctx, ins)
}

func (s *RegisterService) Cancel(ctx context.Context, ns, sname, ip string) error {
	ins, err := s.repos.GetInstanceByIP(ctx, ns, sname, ip)
	if err != nil {
		return err
	}
	if ins.ID == 0 {
		return errors.New("not found service instance")
	}
	now := time.Now()
	ins.DeletedAt = &now
	return s.repos.UpdateInstance(ctx, ins)
}
