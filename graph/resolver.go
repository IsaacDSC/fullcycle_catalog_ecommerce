package graph

import "github.com/IsaacDSC/fullcycle_catalog_ecommerce/shared"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//
//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Repositories *shared.ContainerRepository
}
