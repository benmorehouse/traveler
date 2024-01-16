package config

// Config is the configuration for the api
type Configuration struct {
	Host string `json:"host"`
	Port string `json:"port"`

	DBHost string `json:"db_host"`
	DBPort string `json:"db_port"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`

	APIURL string `json:"api_url"`
}

// DefaultConfig returns the default config for the application
func DefaultConfig() Configuration {
	return Configuration{
		Host:   "localhost",
		Port:   "8080",
		DBHost: "127.0.0.1",
		DBPort: "3306",
		DBUser: "root",
		DBPass: "password",
		APIURL: "https://restcountries.com/v3.1/all",
	}
}
