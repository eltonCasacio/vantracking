package repository

import (
	"database/sql"

	entity "github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"
)

type MonitorRepository struct {
	db *sql.DB
}

func NewMonitorRepository(db *sql.DB) *MonitorRepository {
	return &MonitorRepository{db: db}
}

func (m *MonitorRepository) Create(monitor *entity.Monitor) error {
	stmt, err := m.db.Prepare("insert into monitors(id, nome, senha, ativo) values(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		monitor.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) Update(monitor *entity.Monitor) error {
	stmt, err := m.db.Prepare("update monitors set id = ?, nome = ?, senha = ?, ativo = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(monitor.GetID())
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) Delete(id string) error {
	stmt, err := m.db.Prepare("delete from monitors where id = ?")
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

func (m *MonitorRepository) FindAll() ([]entity.Monitor, error) {
	rows, err := m.db.Query("select * from monitors")
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

func (m *MonitorRepository) FindByID(id string) (*entity.Monitor, error) {
	stmt, err := m.db.Prepare("select * from monitors where id = ?")
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
