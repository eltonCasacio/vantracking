package driver

type User struct {
	ID   string
	Name string
}

type OutputDTO struct {
	AccessToken string `json:"token"`
	User        `json:"user"`
}
