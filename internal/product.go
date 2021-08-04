package Product

import (
	"fmt"
	"goods-square/internal/dependency"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"

	"goods-square/pkg/database"
	libgorm "goods-square/pkg/database/gorm"
)

type Server struct {
	*http.Server
}

type Env struct {
	DBUser     string `required:"true"`
	Port       string `required:"true"`
	DBHost     string `required:"true"`
	DBNet      string `required:"true"`
	DBPort     int    `required:"true"`
	DBPassword string `required:"true"`
	DBName     string `required:"true"`
}

func NewServer() (*Server, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}
	// ctx := context.Background()

	mux := chi.NewMux()

	// indexRoute := &routes.

	db := libgorm.NewRDBClient(&database.Config{
		DBHost:     env.DBHost,
		DBPort:     env.DBPort,
		DBNet:      env.DBNet,
		DBUser:     env.DBUser,
		DBPassword: env.DBPassword,
		DBName:     env.DBName,
	})
	dependency.NewProduct(mux, db)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", env.Port),
		Handler: mux,
	}

	return &Server{&server}, nil

}
