package infrastructure

// Configの構造体
type Config struct {
	DataBase
	Route
}

// DBの構造体
type DataBase struct {
	Develop
}

// 本番環境DB接続用の構造体
type Product struct {
	Host     string
	Username string
	Password string
	DBName   string
	DBPort   string
}

// 開発環境DB接続用の構造体
type Develop struct {
	Host     string
	Username string
	Password string
	DBName   string
	DBPort   string
}

// Routingの構造体
type Route struct {
	Port string
}

// Configの構造体
func NewConfig() *Config {
	c := &Config{
		DataBase: DataBase {
			Develop: Develop {
				Host    : "go_db",
				Username: "go_user",
				Password: "go",
				DBName  : "go_db",
				DBPort: "3306",
			},
		},
		Route: Route {
			Port : "8080",
		},
	}

	return c
}