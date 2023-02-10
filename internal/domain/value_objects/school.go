package value_objects

type School struct {
	code    string
	name    string
	address Address
}

func NewSchool(name string, addr *Address) (*School, error) {
	s := &School{
		code:    "123",
		name:    name,
		address: *addr,
	}
	err := s.IsValid()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *School) IsValid() error {
	return nil
}

func (s *School) GetCode() string {
	return s.code
}
func (s *School) GetName() string {
	return s.name
}
func (s *School) GetAddress() Address {
	return s.address
}
