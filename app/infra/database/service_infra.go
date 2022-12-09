package database

import (
	"context"
	"gorm.io/gorm"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"monica/pkg/log"
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

// GetServiceByNames 批量查询服务
func (s *ServiceRepos) GetServiceByNames(ctx context.Context, nameList []string) ([]*entity.Service, error) {
	servList := make([]*entity.Service, 0)
	if len(nameList) == 0 {
		return servList, nil
	}
	if err := s.db.Where("name in (?) AND deleted_at is null", nameList).Find(&servList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return servList, nil
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

func (s *ServiceRepos) BatchAddServiceNS(ctx context.Context, nsList []*entity.ServiceNamespace) error {
	if len(nsList) == 0 {
		return nil
	}
	if err := s.db.CreateInBatches(nsList, 10).Error; err != nil {
		log.ErrorContextf(ctx, "call db.CreateInBatches failed, err: %v", err)
		return nil
	}
	return nil
}

func (s *ServiceRepos) ListServiceNSByPage(ctx context.Context, kw, ip string, page, pageSize int) ([]*entity.ServiceNamespace, error) {
	tx := s.db.Table("t_service_namespace").Where("service_name like ?", "%"+kw+"%")
	if ip != "" {
		tx = s.db.Table("t_service_namespace sn").Select("sn.*").Where("sn.service_name like ? AND sn.deleted_at is null AND ins.ip = ?", "%"+kw+"%", ip).
			Joins("INNER JOIN t_service_instance as ins ON ins.service_name = sn.service_name AND ins.namespace = sn.namespace")
	}
	nsList := make([]*entity.ServiceNamespace, 0)
	if err := tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&nsList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return nsList, nil
}

func (s *ServiceRepos) CountServiceNSByCond(ctx context.Context, kw, ip string) (int64, error) {
	tx := s.db.Table("t_service_namespace").Select("count(*)").Where("service_name like ?", "%"+kw+"%")
	if ip != "" {
		tx = s.db.Table("t_service_namespace sn").Select("count(*)").Where("sn.service_name like ? AND sn.deleted_at is null AND ins.ip = ?", "%"+kw+"%", ip).
			Joins("INNER JOIN t_service_instance as ins ON ins.service_name = sn.service_name AND ins.namespace = sn.namespace")
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return total, nil
}
