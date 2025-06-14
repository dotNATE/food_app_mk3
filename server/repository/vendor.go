package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type VendorRepository struct {
	DB *sql.DB
}

func NewVendorRepository(db *sql.DB) *VendorRepository {
	return &VendorRepository{DB: db}
}

type Vendor struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	AverageRating float64   `json:"average_rating"`
	CreatedAt     time.Time `json:"created_at"`
}

// GetAllVendors returns all rows in the vendors table
func (repo *VendorRepository) GetAllVendors() ([]Vendor, error) {
	rows, err := DB.Query("SELECT id, name, description, average_rating, created_at FROM vendors")
	if err != nil {
		return nil, fmt.Errorf("failed to query vendors: %w", err)
	}
	defer rows.Close()

	var vendors []Vendor
	for rows.Next() {
		var vendor Vendor

		err := rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description, &vendor.AverageRating, &vendor.CreatedAt)
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
func (repo *VendorRepository) InsertVendor(vendor Vendor) (*Vendor, error) {
	result, err := DB.Exec(
		"INSERT INTO vendors (name, description) VALUES (?, ?)",
		vendor.Name, vendor.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert vendor: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get inserted ID: %w", err)
	}

	var created Vendor
	err = DB.QueryRow(
		"SELECT id, name, description, created_at FROM vendors WHERE id = ?",
		id,
	).Scan(&created.ID, &created.Name, &created.Description, &created.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inserted vendor: %w", err)
	}

	return &created, nil
}

// VendorExists checks whether a vendor exists with the given vendor_id
func (r *VendorRepository) CheckVendorExists(vendor_id int64) (bool, error) {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM vendors WHERE id = ?)", vendor_id).Scan(&exists)
	return exists, err
}

func (r *VendorRepository) UpdateAverageRating(vendor_id int64) error {
	_, err := r.DB.Exec(`
		UPDATE vendors
		SET average_rating = (
			SELECT AVG(score) FROM ratings WHERE vendor_id = ?
		)
		WHERE id = ?`, vendor_id, vendor_id)
	return err
}
