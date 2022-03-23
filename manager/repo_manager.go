package manager

import (
	"enigmacamp.com/warung_makan/repository"
)

type RepoManager interface {
	FoodRepo() repository.FoodRepo
	TableRepo() repository.TableRepo
	CustomersRepo() repository.CustomersRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) FoodRepo() repository.FoodRepo {
	return repository.NewFoodRepo(r.infra.SqlDb())
}

func (r *repoManager) TableRepo() repository.TableRepo {
	return repository.NewTableRepo(r.infra.SqlDb())
}

func (r *repoManager) CustomersRepo() repository.CustomersRepo {
	return repository.NewCustomersRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
