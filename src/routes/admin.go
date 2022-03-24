package routes

import (
	"github.com/Neeraj1997-dev/go-fiber-mongodb-docker-boillerplate/handlers/admin"
)

func admin_routes() {
	API.Post("/admin", admin.AdminRegistration)

}
