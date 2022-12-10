package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"monica/app/domain/entity"
	"monica/app/domain/repository"
	"monica/consts"
	"time"
)

type IService interface {
	// SaveService 更新服务
	SaveService(ctx context.Context, record *entity.Service) error

	// ListServiceByPage 查询服务列表
	ListServiceByPage(ctx context.Context, kw, busi string, page, pageSize int) ([]*entity.Service, int64, error)

	// DeleteService 删除服务
	DeleteService(ctx context.Context, name string) error

	// ListServiceNSByPage 查询服务命名空间
	ListServiceNSByPage(ctx context.Context, kw, ip string, page, pageSize int) ([]*entity.ServiceNSVO, int64, error)

	// VerifyServiceToken 服务注册鉴权
	VerifyServiceToken(ctx context.Context, token, ns, serviceName string) error
}

type Service struct {
	repos    repository.IServiceRepos
	insRepos repository.IServiceInstanceRepos
}

func NewService(repos repository.IServiceRepos, insRepos repository.IServiceInstanceRepos) IService {
	return &Service{
		repos:    repos,
		insRepos: insRepos,
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
		if err := s.repos.CreateService(ctx, record); err != nil {
			return err
		}
		// 创建服务命名空间
		return s.syncGenerateServiceNS(ctx, record.ID, record.Name)
	}
	// 更新服务
	if servDO.ID == 0 {
		return errors.New("服务不存在")
	}
	if servDO.ID != record.ID {
		return errors.New("服务名称已占用")
	}
	// 不支持修改名称
	record.Name = ""
	return s.repos.UpdateService(ctx, record)
}

func (s *Service) syncGenerateServiceNS(ctx context.Context, sid int, sname string) error {
	nsList := []*entity.ServiceNamespace{
		{
			ServiceName: sname,
			Namespace:   consts.Production,
			Token:       uuid.New().String(),
		},
		{
			ServiceName: sname,
			Namespace:   consts.PreRelease,
			Token:       uuid.New().String(),
		},
		{
			ServiceName: sname,
			Namespace:   consts.Test,
			Token:       uuid.New().String(),
		},
		{
			ServiceName: sname,
			Namespace:   consts.Development,
			Token:       uuid.New().String(),
		},
	}
	return s.repos.BatchAddServiceNS(ctx, nsList)
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

// ListServiceNSByPage 查询服务命名空间
func (s *Service) ListServiceNSByPage(ctx context.Context, kw, ip string, page, pageSize int) ([]*entity.ServiceNSVO, int64, error) {
	nsList, err := s.repos.ListServiceNSByPage(ctx, kw, ip, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountServiceNSByCond(ctx, kw, ip)
	if err != nil {
		return nil, 0, err
	}
	snamelist := make([]string, 0)
	for _, ns := range nsList {
		snamelist = append(snamelist, ns.ServiceName)
	}
	// 查询业务名称
	sevList, err := s.repos.GetServiceByNames(ctx, snamelist)
	if err != nil {
		return nil, 0, err
	}
	busimap := make(map[string]string, 0)
	for _, sev := range sevList {
		busimap[sev.Name] = sev.Business
	}
	// 域转换
	nsvoList := make([]*entity.ServiceNSVO, 0)
	for _, sns := range nsList {
		nsvo := &entity.ServiceNSVO{
			Business:  "-",
			Name:      sns.ServiceName,
			Namespace: sns.Namespace,
			Token:     sns.Token,
			Ctime:     sns.CreatedAt.Unix(),
		}
		if bname, ok := busimap[sns.ServiceName]; ok {
			nsvo.Business = bname
		}
		nsvoList = append(nsvoList, nsvo)
	}
	return nsvoList, total, nil
}

func (s *Service) VerifyServiceToken(ctx context.Context, token, ns, serviceName string) error {
	sns, err := s.repos.GetServiceNSByName(ctx, ns, serviceName)
	if err != nil {
		return err
	}
	if sns.ID == 0 {
		return errors.New("not found service namespace")
	}
	if sns.Token != token {
		return errors.New("无效的 Token 令牌")
	}
	return nil
}
