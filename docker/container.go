package docker

import "gimme/database"

type Container struct {
	Name     string
	Database database.Database
}
