package repository

import (
	"database/sql"

	e "github.com/eltoncasacio/vantracking/internal/domain/partner/entity"
)

type partnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) *partnerRepository {
	return &partnerRepository{db: db}
}

func (r *partnerRepository) Create(partner *e.Partner) error {
	stmt, err := r.db.Prepare("INSERT INTO partners (id, name, description, price, phone_number, uf, city, street, number, cep, complement, category_id) values(?,?,?,?,?,?,?,?,?,?,?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		partner.ID.String(),
		partner.Name,
		partner.Description,
		partner.Price,
		partner.PhoneNumber,
		partner.Address.UF,
		partner.Address.City,
		partner.Address.Street,
		partner.Address.Number,
		partner.Address.CEP,
		partner.Address.Complement,
		partner.CategoryID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *partnerRepository) Update(partner *e.Partner) error {
	stmt, err := r.db.Prepare("UPDATE partners SET name=?, description=?, price=?, phone_number=?,  uf=?, city=?, street=?, number=?, cep=?, complement=?, category_id=? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		partner.Name,
		partner.Description,
		partner.Price,
		partner.PhoneNumber,
		partner.Address.UF,
		partner.Address.City,
		partner.Address.Street,
		partner.Address.Number,
		partner.Address.CEP,
		partner.Address.Complement,
		partner.CategoryID,
		partner.ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *partnerRepository) FindAll() ([]e.Partner, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, phone_number,  uf, city, street, number, cep, complement, category_id FROM partners")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Partners []e.Partner
	for rows.Next() {
		inputPartner := e.Partner{}
		err := rows.Scan(
			&inputPartner.ID,
			&inputPartner.Name,
			&inputPartner.Description,
			&inputPartner.Price,
			&inputPartner.PhoneNumber,
			&inputPartner.Address.UF,
			&inputPartner.Address.City,
			&inputPartner.Address.Street,
			&inputPartner.Address.Number,
			&inputPartner.Address.CEP,
			&inputPartner.Address.Complement,
			&inputPartner.CategoryID,
		)
		if err != nil {
			return nil, err
		}
		Partners = append(Partners, inputPartner)
	}
	return Partners, nil
}

func (r *partnerRepository) FindByID(id string) (*e.Partner, error) {
	stmt, err := r.db.Prepare("SELECT id, name, description, price, phone_number,  uf, city, street, number, cep, complement, category_id FROM partners WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	inputPartner := e.Partner{}
	rows := stmt.QueryRow(id)
	rows.Scan(
		&inputPartner.ID,
		&inputPartner.Name,
		&inputPartner.Description,
		&inputPartner.Price,
		&inputPartner.PhoneNumber,
		&inputPartner.Address.UF,
		&inputPartner.Address.City,
		&inputPartner.Address.Street,
		&inputPartner.Address.Number,
		&inputPartner.Address.CEP,
		&inputPartner.Address.Complement,
		&inputPartner.CategoryID,
	)
	if err != nil {
		return nil, err
	}

	return &inputPartner, nil
}

func (r *partnerRepository) Delete(id string) error {
	stmt, err := r.db.Prepare("DELETE FROM partners WHERE id = ?")
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

func (r *partnerRepository) ListByCity(city string) ([]e.Partner, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, phone_number,  uf, city, street, number, cep, complement, category_id FROM partners WHERE city = ?", city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Partners []e.Partner
	for rows.Next() {
		inputPartner := e.Partner{}
		err := rows.Scan(
			&inputPartner.ID,
			&inputPartner.Name,
			&inputPartner.Description,
			&inputPartner.Price,
			&inputPartner.PhoneNumber,
			&inputPartner.Address.UF,
			&inputPartner.Address.City,
			&inputPartner.Address.Street,
			&inputPartner.Address.Number,
			&inputPartner.Address.CEP,
			&inputPartner.Address.Complement,
			&inputPartner.CategoryID,
		)
		if err != nil {
			return nil, err
		}
		Partners = append(Partners, inputPartner)
	}
	return Partners, nil
}
