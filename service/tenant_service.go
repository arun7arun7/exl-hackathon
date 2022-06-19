package service

import (
	"exl-server/entity"
	"exl-server/repository"
)

type TenantService interface {
	GetByAzureTenantID(tenantId string) (*entity.AzureTenant, error)
}

type TenantServiceImpl struct {
	tenantRepository repository.TenantRepository
}

func NewTenantServiceImpl(tp repository.TenantRepository) *TenantServiceImpl {
	return &TenantServiceImpl{
		tenantRepository: tp,
	}
}

func (tenantService *TenantServiceImpl) GetByAzureTenantID(tenantId string) (*entity.AzureTenant, error) {
	return tenantService.tenantRepository.FindByAzureTenantId(tenantId)
}