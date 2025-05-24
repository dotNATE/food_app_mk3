package repository

import (
	"fmt"
)

type Vendor struct {
	ID   int
	Name string
}

// GetAllVendors returns all rows in the vendors table
func GetAllVendors() ([]Vendor, error) {
	rows, err := DB.Query("SELECT id, name FROM vendors")
	if err != nil {
		return nil, fmt.Errorf("failed to query vendors: %w", err)
	}
	defer rows.Close()

	var vendors []Vendor
	for rows.Next() {
		var vendor Vendor

		err := rows.Scan(&vendor.ID, &vendor.Name)
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
