package repository

import (
	"database/sql"

	e "github.com/eltoncasacio/vantracking/internal/domain/partner_category/entity"
)

type partnerCategoryRepository struct {
	db *sql.DB
}

func NewPartnerCategoryRepository(db *sql.DB) *partnerCategoryRepository {
	return &partnerCategoryRepository{db: db}
}

func (r *partnerCategoryRepository) Create(category *e.PartnerCategory) error {
	stmt, err := r.db.Prepare("INSERT INTO categories (id, name) values(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		category.ID.String(),
		category.Name,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *partnerCategoryRepository) Update(category *e.PartnerCategory) error {
	stmt, err := r.db.Prepare("UPDATE categories SET name=? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		category.Name,
		category.ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *partnerCategoryRepository) FindAll() ([]e.PartnerCategory, error) {
	rows, err := r.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Partners []e.PartnerCategory
	for rows.Next() {
		inputPartner := e.PartnerCategory{}
		err := rows.Scan(
			&inputPartner.ID,
			&inputPartner.Name,
		)
		if err != nil {
			return nil, err
		}
		Partners = append(Partners, inputPartner)
	}
	return Partners, nil
}

func (r *partnerCategoryRepository) FindByID(id string) (*e.PartnerCategory, error) {
	stmt, err := r.db.Prepare("SELECT id, name FROM categories WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	inputPartner := e.PartnerCategory{}
	rows := stmt.QueryRow(id)
	rows.Scan(
		&inputPartner.ID,
		&inputPartner.Name,
	)
	if err != nil {
		return nil, err
	}

	return &inputPartner, nil
}

func (r *partnerCategoryRepository) Delete(id string) error {
	stmt, err := r.db.Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
