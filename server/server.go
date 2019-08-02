package server

import (
	"github.com/kanzitelli/good-news-backend/crawler"
)

// Init <function>
// is used to initialize server and all the corresponding services such as DB, Utils, Workers
func Init() {
	// utils
	// utils.InitEnvVars()

	// services
	// db.InitService()

	// workers
	crawler.Start()
	// workers.InitRecieptsDeletion()

	r := NewRouter()
	r.Run(":6969")
}
