package user

import "github.com/alihdrii/go-hexagonal/internal/core/port/output/db"

type Usecase struct {
	userCmdRepo db.UserCommandRepository
	userQryRepo db.UserQueryRepository
}

func NewUserUseCase(
	ucr db.UserCommandRepository,
	uqr db.UserQueryRepository,
) *Usecase {
	return &Usecase{
		userCmdRepo: ucr,
		userQryRepo: uqr,
	}
}
