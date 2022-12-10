package service

import (
	"context"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"monica/consts"
	"time"
)

type IDiscoveryService interface {
	Fetch(ctx context.Context, ns, sname string) ([]*entity.InstanceNode, error)

	Poll(ctx context.Context, ns, sname string) ([]*entity.InstanceNode, error)
}

type DiscoveryService struct {
	repos repository.IServiceInstanceRepos
}

func NewDiscoveryService(insRepos repository.IServiceInstanceRepos) IDiscoveryService {
	return &DiscoveryService{
		repos: insRepos,
	}
}

func (s *DiscoveryService) Fetch(ctx context.Context, ns, sname string) ([]*entity.InstanceNode, error) {
	insList, err := s.repos.GetAllInstance(ctx, ns, sname)
	if err != nil {
		return nil, err
	}
	nodeList := make([]*entity.InstanceNode, 0)
	now := time.Now()
	for _, ins := range insList {
		if ins.Isolate == consts.Isolation || ins.Weight == 0 || ins.RenewedAt == nil {
			continue
		}
		if ins.RenewedAt.Unix()+consts.SyncNSInterval < now.Unix() {
			continue
		}
		node := &entity.InstanceNode{
			IP:       ins.IP,
			Port:     ins.Port,
			Weight:   ins.Weight,
			Metadata: ins.Metadata,
		}
		nodeList = append(nodeList, node)
	}
	return nodeList, nil
}

func (s *DiscoveryService) Poll(ctx context.Context, ns, sname string) ([]*entity.InstanceNode, error) {
	panic("implement me")
}
