package api

import (
	"github.com/kwisnia/inzynierka-backend/internal/api/router"
	"github.com/kwisnia/inzynierka-backend/internal/pkg/config"
	"github.com/kwisnia/inzynierka-backend/internal/pkg/config/database"
)

func Run()  {
	config.LoadConfig()
	r := router.Setup()
	database.SetupDB()
	r.Run()
}