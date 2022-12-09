package database

import (
	"context"
	"gorm.io/gorm"
	"monica-admin/app/domain/entity"
	"monica-admin/app/domain/repository"
	"monica-admin/pkg/log"
)

type ServiceRepos struct {
	db *gorm.DB
}

func NewIServiceRepos(db *gorm.DB) repository.IServiceRepos {
	return &ServiceRepos{
		db: db,
	}
}

func (s *ServiceRepos) CreateService(ctx context.Context, record *entity.Service) error {
	if err := s.db.Create(record).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

// GetServiceByID 通过 ID 查询服务
func (s *ServiceRepos) GetServiceByID(ctx context.Context, sid int) (*entity.Service, error) {
	serv := new(entity.Service)
	if err := s.db.Where("id = ? AND deleted_at is null", sid).Find(serv).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return serv, nil
}

// GetServiceByName 通过名称查询服务
func (s *ServiceRepos) GetServiceByName(ctx context.Context, name string) (*entity.Service, error) {
	serv := new(entity.Service)
	if err := s.db.Where("name = ? AND deleted_at is null", name).Find(serv).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return serv, nil
}

func (s *ServiceRepos) UpdateService(ctx context.Context, record *entity.Service) error {
	if err := s.db.Where("id = ? AND deleted_at is null", record.ID).Updates(record).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (s *ServiceRepos) ListServiceByPage(ctx context.Context, kw, busi string, page, pageSize int) ([]*entity.Service, error) {
	sList := make([]*entity.Service, 0)
	if err := s.db.Where("name like ? AND  business like ? AND deleted_at is null", "%"+kw+"%", "%"+busi+"%").Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&sList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return sList, nil
}

func (s *ServiceRepos) CountServiceByCond(ctx context.Context, kw, busi string) (int64, error) {
	var count int64
	empty := new(entity.Service)
	if err := s.db.Table(empty.TableName()).Select("count(*)").Where("name like ? AND  business like ? AND deleted_at is null", "%"+kw+"%", "%"+busi+"%").Count(&count).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return count, nil
}
