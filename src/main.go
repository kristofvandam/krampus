package main

import (
	"go.hagfi.sh/krampus/database"
	"go.hagfi.sh/krampus/draw"
	"go.hagfi.sh/krampus/logger"
	"go.hagfi.sh/krampus/member"
)

var (
	log = logger.New()
	db  = database.New(log)

	drawCtrl   = draw.Controller{DB: db, Log: log}
	memberCtrl = member.Controller{DB: db, Log: log}
)

// @basePath /api/v1

// @title Krampus API
// @version 1.0
// @description A Privacy Friendly Secret Santa Service
func main() {
	draw.Migrate(db)
	member.Migrate(db)

	loadRouter()
}
