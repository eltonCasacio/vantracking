package repository

import (
	"database/sql"

	entity "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
)

type driverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *driverRepository {
	return &driverRepository{db: db}
}

func (d *driverRepository) Create(driver *entity.Driver) error {
	stmt, err := d.db.Prepare("insert into drivers(id, nome, senha, ativo) values(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		driver.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *driverRepository) Update(driver *entity.Driver) error {
	stmt, err := d.db.Prepare("update drivers set id = ?, nome = ?, senha = ?, ativo = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(driver.GetID())
	if err != nil {
		return err
	}
	return nil
}

func (d *driverRepository) Delete(id string) error {
	stmt, err := d.db.Prepare("delete from drivers where id = ?")
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

func (d *driverRepository) FindAll() ([]entity.Driver, error) {
	rows, err := d.db.Query("select * from drivers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// var drivers []entity.Driver

	// for rows.Next() {
	// 	var usuario entity.Driver
	// 	err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha, &usuario.Ativo)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	var u entity.Usuario
	// 	u.ChangeID(usuario.Id)
	// 	u.ChangeNome(usuario.Nome)
	// 	u.ChangeSenha(usuario.Senha)
	// 	u.Desativar()
	// 	if usuario.Ativo {
	// 		u.Ativar()
	// 	}

	// 	usuarios = append(usuarios, u)
	// }
	return nil, nil

}

func (d *driverRepository) FindByID(id string) (*entity.Driver, error) {
	stmt, err := d.db.Prepare("select * from drivers where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// var usuarioModel usuario.UsuarioModel
	// err = stmt.QueryRow(id).Scan(&usuarioModel.Id, &usuarioModel.Nome, &usuarioModel.Senha, &usuarioModel.Ativo)
	// if err != nil {
	// 	return nil, err
	// }

	// var usuario entity.Usuario
	// usuario.ChangeID(usuarioModel.Id)
	// usuario.ChangeNome(usuarioModel.Nome)
	// usuario.ChangeSenha(usuarioModel.Senha)
	// usuario.Desativar()
	// if usuarioModel.Ativo {
	// 	usuario.Ativar()
	// }
	return nil, nil
}
