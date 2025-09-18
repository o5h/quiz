package db

import "embed"

//go:embed sql/schema/*.sql
var schemaSQLFiles embed.FS
