package main

import (
	"credens/libs/accounts/application/create"
	"credens/libs/accounts/application/read"
	"credens/libs/accounts/domain"
	"credens/libs/accounts/infrastructure/persistence"
	"credens/libs/shared/application/serializer/json_iterator"
	"credens/libs/shared/domain/bus"
	sharedBus "credens/libs/shared/infrastructure/bus"
	"credens/libs/shared/infrastructure/di"
	"credens/libs/shared/infrastructure/logging/logrus"
	"github.com/gin-gonic/gin"
)

const (
	domainPath       = "credens/libs/persistence"
	appPath          = "credens/libs/application"
	infraPath        = "credens/libs/infrastructure"
	sharedDomainPath = "credens/libs/shared/persistence"
	sharedAppPath    = "credens/libs/shared/application"
	sharedInfraPath  = "credens/libs/shared/infrastructure"
)

const (
	LoggerKey                      = infraPath + "/logging/Logger"
	AccountRepositoryKey           = domainPath + "/accounts/account_repository/AccountRepository"
	EventPublisherKey              = sharedDomainPath + "/bus/EventPublisherKey"
	CommandHandlerDictKey          = sharedDomainPath + "/bus/CommandHandler[]"
	CommandBusKey                  = sharedInfraPath + "/bus/CommandBus"
	CreateAccountCommandHandlerKey = appPath + "/create/create_account_command_handler/CreateAccountCommandHandler"
	QueryBusKey                    = sharedDomainPath + "/bus/QueryBus"
	QueryHandlerSliceKey           = sharedDomainPath + "/bus/QueryHandler[]"
	JSONSerializerKey              = sharedAppPath + "/serializer/json_serializer/JSONSerializer"
	HttpServerKey                  = "github.com/gin-gonic/gin"
)

func BuildContainer(env Environment) *di.Container {
	ctx := di.NewContainer()

	ctx.Set(LoggerKey, func(_ *di.Container) interface{} {
		return logrus.NewLogger()
	})

	ctx.Set(AccountRepositoryKey, func(_ *di.Container) interface{} {
		return *persistence.NewInMemoryAccountRepository([]*domain.Account{})
	})

	ctx.Set(EventPublisherKey, func(_ *di.Container) interface{} {
		return *sharedBus.NewInMemoryEventPublisher()
	})

	ctx.SetEmptyDict(CommandHandlerDictKey)

	ctx.SetInDict(
		CommandHandlerDictKey,
		CreateAccountCommandHandlerKey,
		func(container *di.Container) interface{} {
			return *create.NewCreateAccountCommandHandler(
				*create.NewCreateAccountService(
					container.Get(AccountRepositoryKey).(domain.AccountRepository),
					container.Get(EventPublisherKey).(bus.EventPublisher),
				),
			)
		},
	)

	ctx.Set(
		CommandBusKey,
		func(container *di.Container) interface{} {
			return sharedBus.NewSyncCommandBus(container.GetDictAsSlice(CommandHandlerDictKey))
		},
	)

	ctx.SetEmptySlice(QueryHandlerSliceKey)

	ctx.SetInSlice(
		QueryHandlerSliceKey,
		func(container *di.Container) interface{} {
			return read.NewReadAccountQueryHandler(
				*read.NewReadAccountService(
					container.Get(AccountRepositoryKey).(domain.AccountRepository),
				),
			)
		},
	)

	ctx.Set(
		QueryBusKey,
		func(container *di.Container) interface{} {
			return sharedBus.NewSyncQueryBus(container.GetSlice(QueryHandlerSliceKey))
		},
	)

	ctx.Set(
		JSONSerializerKey,
		func(container *di.Container) interface{} {
			return *json_iterator.NewJSONIteratorJSONSerializer()
		},
	)

	ctx.Set(
		HttpServerKey,
		func(container *di.Container) interface{} {
			var mode = gin.ReleaseMode
			if env.Debug {
				switch env.Env {
				case "test", "testing":
					mode = gin.TestMode
				default:
					mode = gin.DebugMode
				}
			}
			gin.SetMode(mode)

			server := gin.Default()

			AddRoutes(server, container)

			return *server
		},
	)

	return ctx
}
