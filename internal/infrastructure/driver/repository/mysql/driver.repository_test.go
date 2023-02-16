package repository

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/suite"
)

type DriverRepositoryTestSuite struct {
	suite.Suite
	Db         *sql.DB
	Repository *DriverRepository
	Driver     *entity.Driver
}

func (suite *DriverRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec(`CREATE TABLE drivers (
		id varchar(255) NOT NULL,
		cpf varchar(255) NOT NULL,
		name varchar(255) NOT NULL,
		nickname varchar(255),
		phone varchar(255),
		uf varchar(5),
		city varchar(80),
		street varchar(100),
		number varchar(50),
		cep varchar(10),
		active BOOLEAN DEFAULT true,
		PRIMARY KEY (id))`,
	)
	suite.Db = db

	addr, err := vo.NewAddress("uf", "city", "street", "number", "123")
	driver, err := entity.NewDriver("", "any_cpf", "any_name", "any_nickname", "any_phone", *addr)

	repo := NewDriverRepository(suite.Db)

	suite.Repository = repo
	suite.Driver = driver
}

func (suite *DriverRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DriverRepositoryTestSuite))
}

func (s *DriverRepositoryTestSuite) CreateDriver() {
	query := `INSERT INTO 
	drivers (id, cpf, name, nickname, phone, uf, city, street, number, cep, active)
	values(?,?,?,?,?,?,?,?,?,?,?)`
	stmt, err := s.Db.Prepare(query)
	s.Nil(err)

	addr := s.Driver.GetAddress()

	_, err = stmt.Exec(
		s.Driver.GetID().String(),
		s.Driver.GetCPF(),
		s.Driver.GetName(),
		s.Driver.GetNickName(),
		s.Driver.GetPhone(),
		addr.GetUF(),
		addr.GetCity(),
		addr.GetStreet(),
		addr.GetNumber(),
		addr.GetCEP(),
		true,
	)
	s.Nil(err)
}

func (s *DriverRepositoryTestSuite) TestCreate() {
	err := s.Repository.Create(s.Driver)
	s.Nil(err)

	stmt, err := s.Db.Prepare("SELECT * FROM drivers WHERE id = ? and active = true")
	if err != nil {
		panic(err)
	}
	var model DriverModel
	row := stmt.QueryRow(s.Driver.GetID().String())
	err = row.Scan(
		&model.id,
		&model.cpf,
		&model.name,
		&model.nickname,
		&model.phone,
		&model.uf,
		&model.city,
		&model.street,
		&model.number,
		&model.cep,
		&model.active,
	)
	s.Nil(err)
	s.Equal(s.Driver.GetID().String(), model.id)
}

func (s *DriverRepositoryTestSuite) TestCreate_InvalidEntity() {
	err := s.Repository.Create(&entity.Driver{})
	s.NotNil(err)

	stmt, err := s.Db.Prepare("SELECT * FROM drivers WHERE id = ?")
	if err != nil {
		panic(err)
	}
	var model DriverModel
	err = stmt.QueryRow(
		s.Driver.GetID().String()).Scan(
		&model.id,
		&model.cpf,
		&model.name,
		&model.nickname,
		&model.phone,
		&model.uf,
		&model.city,
		&model.street,
		&model.number,
		&model.cep,
	)
	s.NotNil(err)
}

func (s *DriverRepositoryTestSuite) TestFindByID() {
	s.CreateDriver()
	d, err := s.Repository.FindByID(s.Driver.GetID().String())
	s.Nil(err)
	s.NotNil(d)
	s.Equal(d.GetID(), s.Driver.GetID())
	s.Equal(d.GetCPF(), s.Driver.GetCPF())
}

func (s *DriverRepositoryTestSuite) TestFindByID_InvalidID() {
	d, err := s.Repository.FindByID("")
	s.Nil(d)
	s.NotNil(err)
	s.EqualError(err, "id is required")
}

// Para passar esse teste a query de repository deve ser alterada para dar erro
func (s *DriverRepositoryTestSuite) TestFindByID_NoSuchTable() {
	s.CreateDriver()
	d, err := s.Repository.FindByID(s.Driver.GetID().String())
	s.Nil(d)
	s.NotNil(err)
	s.EqualError(err, "no such table: driver")
}

func (s *DriverRepositoryTestSuite) TestUpdate() {
	s.CreateDriver()
	id := s.Driver.GetID()

	s.Equal(s.Driver.GetCPF(), "any_cpf")
	s.Equal(s.Driver.GetName(), "any_name")
	s.Equal(s.Driver.GetNickName(), "any_nickname")
	s.Equal(s.Driver.GetPhone(), "any_phone")

	address := s.Driver.GetAddress()

	s.Equal(address.GetCEP(), "123")
	s.Equal(address.GetCity(), "city")
	s.Equal(address.GetNumber(), "number")
	s.Equal(address.GetStreet(), "street")
	s.Equal(address.GetUF(), "uf")

	s.Driver.ChangeCPF("other_cpf")
	s.Driver.ChangeName("other_name")
	s.Driver.ChangeNickname("other_nickname")
	s.Driver.ChangePhone("other_phone")

	addr, _ := vo.NewAddress("other_uf", "other_city", "other_street", "other_number", "other_cep")
	s.Driver.ChangeAddress(*addr)

	err := s.Repository.Update(s.Driver)
	s.Nil(err)
	s.Equal(s.Driver.GetCPF(), "other_cpf")
	s.Equal(s.Driver.GetID(), id)
}

func (s *DriverRepositoryTestSuite) TestUpdate_InvalidValue() {
	err := s.Repository.Update(&entity.Driver{})
	s.NotNil(err)
	s.EqualError(err, "invalid driver")
}

func (s *DriverRepositoryTestSuite) TestDelete() {
	s.CreateDriver()
	s.Repository.Delete(s.Driver.GetID().String())
	_, err := s.Repository.FindByID(s.Driver.GetID().String())
	s.NotNil(err)
}

func (s *DriverRepositoryTestSuite) TestDelete_InvalidID() {
	s.CreateDriver()
	s.Repository.Delete(identity.NewID().String())
	_, err := s.Repository.FindByID(s.Driver.GetID().String())
	s.Nil(err)
}

func (s *DriverRepositoryTestSuite) TestFindAll() {
	addr, _ := vo.NewAddress("uf", "city", "street", "number", "123")
	driver1, _ := entity.NewDriver("", "any_cpf", "any_name", "any_nickname", "any_phone", *addr)
	driver2, _ := entity.NewDriver("", "any_cpf", "any_name", "any_nickname", "any_phone", *addr)
	driver3, _ := entity.NewDriver("", "any_cpf", "any_name", "any_nickname", "any_phone", *addr)

	repo := NewDriverRepository(s.Db)

	repo.Create(driver1)
	repo.Create(driver2)
	repo.Create(driver3)

	drivers, err := s.Repository.FindAll()
	s.Nil(err)
	s.Equal(len(drivers), 3)
}
