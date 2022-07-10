package service

import (
	"exl-server/entity"
	"exl-server/repository"
)

type TenantService interface {
	GetByAzureOrgID(orgId string) (*entity.AzureTenant, error)
	GetByAwsOrgID(orgId string) (*entity.AwsTenant, error)
}

type TenantServiceImpl struct {
	tenantRepository repository.TenantRepository
}

func NewTenantServiceImpl(tp repository.TenantRepository) *TenantServiceImpl {
	return &TenantServiceImpl{
		tenantRepository: tp,
	}
}

func (tenantService *TenantServiceImpl) GetByAzureOrgID(orgId string) (*entity.AzureTenant, error) {
	return tenantService.tenantRepository.FindByAzureOrgId(orgId)
}

func (tenantService *TenantServiceImpl) GetByAwsOrgID(orgId string) (*entity.AwsTenant, error) {
	return tenantService.tenantRepository.FindByAwsOrgId(orgId)
}