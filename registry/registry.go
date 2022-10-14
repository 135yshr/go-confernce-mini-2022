package registry

import (
	"github.com/135yshr/go-confernce-mini-2022/domain/repository"
	"github.com/135yshr/go-confernce-mini-2022/infrastructure/persistence"
	handler "github.com/135yshr/go-confernce-mini-2022/interfaces/handlers"
	"github.com/135yshr/go-confernce-mini-2022/usecase"
)

type Registry struct {
	UserRepository repository.UserRepository
	UserUseCase    usecase.UserUsecase
	UserHandler    handler.UserHandler
}

func NewRegistry() *Registry {
	r := &Registry{}

	r.UserRepository = persistence.NewUserRepository(persistence.Instance())
	r.UserUseCase = usecase.NewUserUsecase(r.UserRepository)
	r.UserHandler = handler.NewUserHandler(r.UserUseCase)

	return r
}
