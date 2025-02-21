package servicecontainer

import (
	"database/sql"
	"goservertemplate/persistence/repositories"
	"goservertemplate/types"
)

type RepositoryContainer struct {
	UserRepository *repositories.UserRepository
}

type ServiceContainer struct {
}

type Container struct {
	Repositories *RepositoryContainer
	Services     *ServiceContainer
	Config       *types.Configuration
	DB           *sql.DB
}

func NewServiceContainer(config *types.Configuration, db *sql.DB) *Container {
	container := &Container{}

	container.Config = config
	container.DB = db
	container.Repositories = getRepositories(db)
	container.Services = getServiceContainer(container)

	return container
}

func getServiceContainer(container *Container) *ServiceContainer {
	s := &ServiceContainer{}

	return s
}

func getRepositories(db *sql.DB) *RepositoryContainer {
	r := &RepositoryContainer{}

	r.UserRepository = repositories.NewUserRepository(db)

	return r
}
