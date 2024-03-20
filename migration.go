package main

import "emreddit/migration_manager"

func main() {
	migration_manager.RunUp()
	//migration_manager.RunDown()
}
