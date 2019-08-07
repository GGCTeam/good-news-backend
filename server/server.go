package server

import (
	"github.com/kanzitelli/good-news-backend/crawler"
	"github.com/kanzitelli/good-news-backend/db"
	"github.com/kanzitelli/good-news-backend/utils"
)

// Init <function>
// is used to initialize server and all the corresponding services such as DB, Utils, Workers
func Init() {
	// utils
	utils.InitEnvVars()

	// services
	db.InitService()
	db.GetClient().FillSeedsInformation() // a bit ugly but everything is going to work, only needed to fill seeds information

	// workers
	crawler.Start()

	r := NewRouter()
	r.Run(":6969")
}
