package configs

//go mod tidy pra baixar o viper
import "github.com/spf13/viper"

var cfg *config // ponteiro de config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct { // estrutura do API
	Port string
}

type DBConfig struct { // estrutura do DB
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

//chamada sempre no start das aplicaçoes com valores default
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error { // carrega os valores das informações inseridas
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err

		}

	}

	cfg = new(config) // cria ponteiro de config

	cfg.API = APIConfig{

		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil

}

//retorna as informaçoes de BD 14:50
func GetDB() DBConfig {
	return cfg.DB
}

// retorna a informação da porta
func GetServerPort() string {
	return cfg.API.Port
}
