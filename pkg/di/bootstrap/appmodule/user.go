package appmodule

import (
	"fmt"
	"github.com/alihdrii/go-hexagonal/internal/adapter/http/handler"
	"github.com/alihdrii/go-hexagonal/internal/core/usecase/user"
	bootstrap "github.com/alihdrii/go-hexagonal/pkg/di"
	"github.com/alihdrii/go-hexagonal/pkg/di/container"

	repoCmd "github.com/alihdrii/go-hexagonal/internal/adapter/repository/command"
	repoQry "github.com/alihdrii/go-hexagonal/internal/adapter/repository/query"
)

type UserModule struct{}

func (m *UserModule) Register(mc container.MainContainerI) {
	if ac, ok := mc.(*container.AppContainer); !ok {
		fmt.Println("Failed to type assertion Module")
	} else {
		userCmdRepo := repoCmd.NewUserCommandRepository(ac.DB.Conn)
		userQryRepo := repoQry.NewUserQueryRepository(ac.DB.Conn)
		userUseCase := user.NewUserUseCase(userCmdRepo, userQryRepo)
		userHandler := handler.NewUserHandler(userUseCase)

		ac.Handlers.User = userHandler
	}

}

func init() {
	bootstrap.RegisterModule(&UserModule{})
}
