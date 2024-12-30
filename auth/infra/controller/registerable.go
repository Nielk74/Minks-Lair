package controller

import "cosmink/libs/route"

// registerable is an interface that defines a method to register routes
type Registerable interface {
	RegisterRoutes(server *route.Server)
}
