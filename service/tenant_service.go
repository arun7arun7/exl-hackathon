package service

import (
	"exl-server/entity"
	"exl-server/repository"
)

type TenantService interface {
	GetByAzureTenantID(tenantId string) (*entity.AzureTenant, error)
	GetByAwsTenantID(tenantId string) (*entity.AwsTenant, error)
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

func (tenantService *TenantServiceImpl) GetByAwsTenantID(tenantId string) (*entity.AwsTenant, error) {
	return tenantService.tenantRepository.FindByAwsTenantId(tenantId)
}