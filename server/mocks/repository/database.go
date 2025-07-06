package repo_mocks

import "gorm.io/gorm"

type MockGormDBWrapper struct {
	DB *gorm.DB
}

func (m *MockGormDBWrapper) WithTransaction(fn func(tx *gorm.DB) error) error {
	return fn(&gorm.DB{})
}

func (m *MockGormDBWrapper) GetDB() *gorm.DB {
	return &gorm.DB{}
}
