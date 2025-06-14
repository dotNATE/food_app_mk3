package repository

import (
	"database/sql"
	"fmt"
)

type VendorRepository struct {
	DB *sql.DB
}

func NewVendorRepository(db *sql.DB) *VendorRepository {
	return &VendorRepository{DB: db}
}

type Vendor struct {
	ID          int     `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// GetAllVendors returns all rows in the vendors table
func (repo *VendorRepository) GetAllVendors() ([]Vendor, error) {
	rows, err := DB.Query("SELECT id, name, description FROM vendors")
	if err != nil {
		return nil, fmt.Errorf("failed to query vendors: %w", err)
	}
	defer rows.Close()

	var vendors []Vendor
	for rows.Next() {
		var vendor Vendor

		err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vendor row: %w", err)
		}

		vendors = append(vendors, vendor)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return vendors, nil
}

// InsertVendor inserts one new row only to the vendors table
func (repo *VendorRepository) InsertVendor(vendor Vendor) (int64, error) {
	result, err := DB.Exec(
		"INSERT INTO vendors (name, description) VALUES (?, ?)",
		vendor.Name, vendor.Description,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert vendor: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get inserted ID: %w", err)
	}

	return id, nil
}

// VendorExists checks whether a vendor exists with the given vendor_id
func (r *VendorRepository) CheckVendorExists(vendor_id int) (bool, error) {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM vendors WHERE id = ?)", vendor_id).Scan(&exists)
	return exists, err
}
