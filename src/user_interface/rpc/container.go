package main

import (
	"credens/src/application/create"
	"credens/src/application/read"
	"credens/src/domain/account"
	infraDomainAccount "credens/src/infrastructure/domain/account"
	"credens/src/infrastructure/logging/logrus"
	"credens/src/shared/domain/bus"
	sharedBus "credens/src/shared/infrastructure/bus"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
)

const (
	domainPath       = "credens/src/domain"
	appPath          = "credens/src/application"
	infraPath        = "credens/src/infrastructure"
	sharedDomainPath = "credens/src/shared/domain"
	sharedInfraPath  = "credens/src/shared/infrastructure"
)

const (
	LoggerKey                      = infraPath + "/logging/Logger"
	AccountRepositoryKey           = domainPath + "/account/account_repository/AccountRepository"
	AccountBuilderKey              = domainPath + "/account/account_builder/AccountBuilder"
	EventPublisherKey              = sharedDomainPath + "/bus/EventPublisherKey"
	CommandHandlerDictKey          = sharedDomainPath + "/bus/CommandHandler[]"
	CommandBusKey                  = sharedInfraPath + "/bus/CommandBus"
	CreateAccountCommandHandlerKey = appPath + "/create/create_account_command_handler/CreateAccountCommandHandler"
	QueryBusKey                    = sharedDomainPath + "/bus/QueryBus"
	QueryHandlerSliceKey           = sharedDomainPath + "/bus/QueryHandler[]"
)

func NewContainer(env config.Env, debug config.Debug) *user_interface.Container {
	ctx := user_interface.NewContainer()

	ctx.Set(LoggerKey, func(_ *user_interface.Container) interface{} {
		return logrus.NewLogger()
	})

	ctx.Set(AccountRepositoryKey, func(_ *user_interface.Container) interface{} {
		return *infraDomainAccount.NewInMemoryAccountRepository()
	})

	ctx.Set(AccountBuilderKey, func(_ *user_interface.Container) interface{} {
		return *account.NewAccountBuilder()
	})

	ctx.Set(EventPublisherKey, func(_ *user_interface.Container) interface{} {
		return *sharedBus.NewInMemoryEventPublisher()
	})

	ctx.SetEmptyDict(CommandHandlerDictKey)

	ctx.SetInDict(
		CommandHandlerDictKey,
		CreateAccountCommandHandlerKey,
		func(container *user_interface.Container) interface{} {
			return *create.NewCreateAccountCommandHandler(
				container.Get(AccountRepositoryKey).(account.AccountRepository),
				container.Get(AccountBuilderKey).(account.AccountBuilder),
				container.Get(EventPublisherKey).(bus.EventPublisher),
			)
		},
	)

	ctx.Set(
		CommandBusKey,
		func(container *user_interface.Container) interface{} {
			return sharedBus.NewSyncCommandBus(container.GetDictAsSlice(CommandHandlerDictKey))
		},
	)

	ctx.SetEmptySlice(QueryHandlerSliceKey)

	ctx.SetInSlice(
		QueryHandlerSliceKey,
		func(container *user_interface.Container) interface{} {
			return read.NewReadAccountQueryHandler(
				container.Get(AccountRepositoryKey).(account.AccountRepository),
			)
		},
	)

	ctx.Set(
		QueryBusKey,
		func(container *user_interface.Container) interface{} {
			return sharedBus.NewSyncQueryBus(container.GetSlice(QueryHandlerSliceKey))
		},
	)

	return ctx
}