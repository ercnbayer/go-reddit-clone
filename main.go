package main

import (
	// _ "emreddit/config"
	// _ "emreddit/logger"
	// _ "emreddit/migration-utils"
	// migrationutils "emreddit/migration-utils"

	"emreddit/app/api"
)

func main() {
	api.ListenPort()
}
