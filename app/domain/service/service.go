package service

import (
	"context"
	"errors"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"time"
)

type IService interface {
	// SaveService 更新服务
	SaveService(ctx context.Context, record *entity.Service) error

	// ListServiceByPage 查询服务列表
	ListServiceByPage(ctx context.Context, kw, busi string, page, pageSize int) ([]*entity.Service, int64, error)

	// DeleteService 删除服务
	DeleteService(ctx context.Context, name string) error
}

type Service struct {
	repos repository.IServiceRepos
}

func NewService(repos repository.IServiceRepos) IService {
	return &Service{
		repos: repos,
	}
}

func (s *Service) SaveService(ctx context.Context, record *entity.Service) error {
	servDO, err := s.repos.GetServiceByName(ctx, record.Name)
	if err != nil {
		return err
	}
	// 新增服务
	if record.ID == 0 {
		if servDO.ID != 0 {
			return errors.New("服务名称已占用")
		}
		return s.repos.CreateService(ctx, record)
	}
	// 更新服务
	if servDO.ID == 0 {
		return errors.New("服务不存在")
	}
	if servDO.ID != record.ID {
		return errors.New("服务名称已占用")
	}
	return s.repos.UpdateService(ctx, record)
}

func (s *Service) ListServiceByPage(ctx context.Context, kw, busi string, page, pageSize int) ([]*entity.Service, int64, error) {
	list, err := s.repos.ListServiceByPage(ctx, kw, busi, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountServiceByCond(ctx, kw, busi)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// DeleteService 删除服务
func (s *Service) DeleteService(ctx context.Context, name string) error {
	serv, err := s.repos.GetServiceByName(ctx, name)
	if err != nil {
		return err
	}
	if serv.ID == 0 {
		return errors.New("服务不存在")
	}
	// 标记删除
	now := time.Now()
	serv.DeletedAt = &now
	return s.repos.UpdateService(ctx, serv)
}
