package service

import (
	"main/dto"
	"main/models"
	"main/repository"
)

type VendorService struct {
	VendorRepo repository.VendorRepository
}

func NewVendorService(vendorRepo repository.VendorRepository) *VendorService {
	return &VendorService{
		VendorRepo: vendorRepo,
	}
}

func (service *VendorService) GetAllVendors() ([]*dto.Vendor, error) {
	vendors, err := service.VendorRepo.GetAllVendors()
	if err != nil {
		return nil, err
	}

	result := make([]*dto.Vendor, len(vendors))
	for index, vendor_model := range vendors {
		result[index] = dto.ConvertVendorModelToDto(vendor_model)
	}

	return result, nil
}

func (service *VendorService) CreateNewVendor(vendor_request *dto.AddVendorRequest, created_by_user_id int64) (*dto.Vendor, error) {
	vendor, err := service.VendorRepo.InsertVendor(models.Vendor{
		Name:        vendor_request.Name,
		Description: vendor_request.Description,
		CreatedBy:   created_by_user_id,
	})
	if err != nil {
		return nil, err
	}

	return dto.ConvertVendorModelToDto(vendor), nil
}

func (service *VendorService) GetVendorById(vendor_id int64) (*dto.Vendor, error) {
	vendor, err := service.VendorRepo.GetVendorById(vendor_id)
	if err != nil {
		return nil, err
	}

	return dto.ConvertVendorModelToDto(vendor), nil
}

func (service *VendorService) CheckVendorExists(vendor_id int64) (bool, error) {
	vendor_exists, err := service.VendorRepo.CheckVendorExists(vendor_id)
	if err != nil {
		return false, err
	}

	return vendor_exists, nil
}

func (service *VendorService) UpdateAverageRating(vendor_id int64) error {
	err := service.VendorRepo.UpdateAverageRating(vendor_id)
	if err != nil {
		return err
	}

	return nil
}
