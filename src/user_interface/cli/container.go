package main

import (
	"credens/src/application/create"
	"credens/src/domain/account"
	infraDomainAccount "credens/src/infrastructure/domain/account"
	"credens/src/infrastructure/logging"
	"credens/src/infrastructure/logging/fmt"
	"credens/src/shared/domain/bus"
	sharedBus "credens/src/shared/infrastructure/bus"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"credens/src/user_interface/cli/command"
	"github.com/spf13/cobra"
)

const (
	LoggerKey                      string = "credens/src/infrastructure/logging/Logger"
	AccountRepositoryKey                  = "credens/src/domain/account/account_repository/AccountRepository"
	AccountBuilderKey                     = "credens/src/domain/account/account_builder/AccountBuilder"
	EventPublisherKey                     = "credens/src/shared/domain/bus/EventPublisherKey"
	CommandHandlerDictKey                 = "credens/src/shared/domain/bus/CommandHandler[]"
	CommandBusKey                         = "credens/src/shared/infrastructure/bus/CommandBus"
	CreateAccountCommandHandlerKey        = "credens/src/application/create/create_account_command_handler/CreateAccountCommandHandler"
	EventSubscriberDictKey                = "credens/src/shared/domain/bus/EventSubscriber[]"
	EventBusKey                           = "credens/src/shared/domain/bus/EventBus"
	QueryBusKey                           = "credens/src/shared/domain/bus/QueryBus"
	QueryHandlerSliceKey                  = "credens/src/shared/domain/bus/QueryHandler[]"
	RootCmdKey                            = "github.com/spf13/cobra/Command"
)

func NewContainer(env config.Env, debug config.Debug) *user_interface.Container {
	ctx := user_interface.NewContainer()

	ctx.Set(LoggerKey, func(_ *user_interface.Container) interface{} {
		return fmt.NewLogger()
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

	ctx.SetEmptySlice(EventSubscriberDictKey)

	ctx.Set(
		EventBusKey,
		func(container *user_interface.Container) interface{} {
			return sharedBus.NewSyncEventBus(container.GetSlice(EventSubscriberDictKey))
		},
	)

	ctx.SetEmptySlice(QueryHandlerSliceKey)

	ctx.Set(
		QueryBusKey,
		func(container *user_interface.Container) interface{} {
			return sharedBus.NewSyncQueryBus(container.GetSlice(QueryHandlerSliceKey))
		},
	)

	ctx.Set(RootCmdKey, func(container *user_interface.Container) interface{} {
		rootCmd := cobra.Command{}

		logger := ctx.Get(LoggerKey).(logging.Logger)

		rootCmd.AddCommand(
			command.NewHelloCommand(logger),
			command.NewCreateAccountCommand(
				container.GetInDict(CommandHandlerDictKey, CreateAccountCommandHandlerKey).(create.CreateAccountCommandHandler),
			),
		)

		return rootCmd
	})

	return ctx
}
