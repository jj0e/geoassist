package files

type Manager struct {
	Directory      string
	ConfigJSONPath string
}

type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
