package config

import (
	"credens/libs/shared/infrastructure/di"

	accountAppCreate "credens/libs/accounts/application/create"
	accountAppRead "credens/libs/accounts/application/read"
	accountDomain "credens/libs/accounts/domain"
	accountInfraPersistence "credens/libs/accounts/infrastructure/persistence"

	sharedDomainBus "credens/libs/shared/domain/bus"
	sharedInfraBus "credens/libs/shared/infrastructure/bus"
	sharedInfraLogging "credens/libs/shared/infrastructure/logging/logrus"
)

const (
	AccountRepositoryKey = "AccountRepository"

	CommandBusKey         = "CommandBus"
	CommandHandlerDictKey = "CommandHandler[]"

	QueryBusKey          = "QueryBus"
	QueryHandlerSliceKey = "QueryHandler[]"

	EventPublisherKey = "EventPublisher"

	LoggerKey = "Logger"
)

func BuildContainer(env Environment) *di.Container {
	ctx := di.NewContainer()

	setCommonDependencies(ctx, env)
	setCommandDependencies(ctx, env)
	setQueryDependencies(ctx, env)

	return ctx
}

func setCommonDependencies(ctx *di.Container, env Environment) {
	ctx.Set(LoggerKey, func(_ *di.Container) interface{} {
		return sharedInfraLogging.NewLogger()
	})

	ctx.Set(AccountRepositoryKey, func(_ *di.Container) interface{} {
		return *accountInfraPersistence.NewInMemoryAccountRepository([]*accountDomain.Account{})
	})

	ctx.Set(EventPublisherKey, func(_ *di.Container) interface{} {
		return *sharedInfraBus.NewInMemoryEventPublisher()
	})
}

func setCommandDependencies(ctx *di.Container, env Environment) {
	ctx.SetEmptyDict(CommandHandlerDictKey)

	ctx.SetInDict(
		CommandHandlerDictKey,
		"CreateAccountCommandHandler",
		func(container *di.Container) interface{} {
			return *accountAppCreate.NewCreateAccountCommandHandler(
				*accountAppCreate.NewCreateAccountService(
					container.Get(AccountRepositoryKey).(accountDomain.AccountRepository),
					container.Get(EventPublisherKey).(sharedDomainBus.EventPublisher),
				),
			)
		},
	)

	ctx.Set(
		CommandBusKey,
		func(container *di.Container) interface{} {
			return sharedInfraBus.NewSyncCommandBus(container.GetDictAsSlice(CommandHandlerDictKey))
		},
	)
}

func setQueryDependencies(ctx *di.Container, env Environment) {
	ctx.SetEmptySlice(QueryHandlerSliceKey)

	ctx.SetInSlice(
		QueryHandlerSliceKey,
		func(container *di.Container) interface{} {
			return accountAppRead.NewReadAccountQueryHandler(
				*accountAppRead.NewReadAccountService(
					container.Get(AccountRepositoryKey).(accountDomain.AccountRepository),
				),
			)
		},
	)

	ctx.Set(
		QueryBusKey,
		func(container *di.Container) interface{} {
			return sharedInfraBus.NewSyncQueryBus(container.GetSlice(QueryHandlerSliceKey))
		},
	)
}
