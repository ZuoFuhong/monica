package database

import (
	"context"
	"gorm.io/gorm"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"monica/pkg/log"
)

type ServiceInstanceRepos struct {
	db *gorm.DB
}

func NewIServiceInstanceRepos(db *gorm.DB) repository.IServiceInstanceRepos {
	return &ServiceInstanceRepos{
		db: db,
	}
}

func (s *ServiceInstanceRepos) CreateInstance(ctx context.Context, record *entity.ServiceInstance) error {
	if err := s.db.Create(record).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (s *ServiceInstanceRepos) GetInstanceByIP(ctx context.Context, ns string, serviceId int, ip string) (*entity.ServiceInstance, error) {
	ins := new(entity.ServiceInstance)
	if err := s.db.Where("namespace = ? AND service_id = ? AND ip = ? AND deleted_at is null", ns, serviceId, ip).Find(ins).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return ins, nil
}

func (s *ServiceInstanceRepos) UpdateInstance(ctx context.Context, ins *entity.ServiceInstance) error {
	if err := s.db.Where("id = ? AND deleted_at is null", ins.ID).Updates(ins).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}

func (s *ServiceInstanceRepos) ListInstanceByPage(ctx context.Context, ns string, serviceId int, ip string, port int, page, pageSize int) ([]*entity.ServiceInstance, error) {
	insList := make([]*entity.ServiceInstance, 0)
	tx := s.db.Where("namespace = ? AND service_id = ? AND deleted_at is null", ns, serviceId)
	if ip != "" {
		tx = tx.Where("ip = ?", ip)
	}
	if port != 0 {
		tx = tx.Where("port = ?", port)
	}
	if err := tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&insList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return insList, nil
}

func (s *ServiceInstanceRepos) CountInstanceByCond(ctx context.Context, ns string, serviceId int, ip string, port int) (int64, error) {
	var count int64
	empty := new(entity.ServiceInstance)
	tx := s.db.Table(empty.TableName()).Select("count(*)").Where("namespace = ? AND service_id = ? AND deleted_at is null", ns, serviceId)
	if ip != "" {
		tx = tx.Where("ip = ?", ip)
	}
	if port != 0 {
		tx = tx.Where("port = ?", port)
	}
	if err := tx.Count(&count).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return count, nil
}
