package config

import (
	"credens/apps/cli/command"
	"github.com/defval/inject"

	accountAppCreate "credens/libs/accounts/application/create"
	accountAppRead "credens/libs/accounts/application/read"
	accountDomain "credens/libs/accounts/domain"
	accountInfraPersistence "credens/libs/accounts/infrastructure/persistence"

	sharedDomainBus "credens/libs/shared/domain/bus"
	sharedInfraBus "credens/libs/shared/infrastructure/bus"
	shareInfraLogging "credens/libs/shared/infrastructure/logging"
	sharedInfraLoggingFmt "credens/libs/shared/infrastructure/logging/fmt"
)

func BuildContainer(env Environment) (*inject.Container, error) {
	return inject.New(
		inject.Provide(
			sharedInfraLoggingFmt.NewFmtLogger,
			inject.As(new(shareInfraLogging.Logger)),
		),

		inject.Provide(
			accountInfraPersistence.NewInMemoryAccountRepository([]*accountDomain.Account{}),
			inject.As(new(accountDomain.AccountRepository)),
		),

		inject.Provide(
			sharedInfraBus.NewInMemoryEventPublisher,
			inject.As(new(sharedDomainBus.EventPublisher)),
		),

		inject.Provide(
			accountAppCreate.NewCreateAccountCommandHandler,
			inject.As(new(sharedDomainBus.CommandHandler)),
		),

		inject.Provide(
			sharedInfraBus.NewSyncCommandBus,
			inject.As(new(sharedDomainBus.CommandBus)),
		),

		inject.Provide(
			accountAppRead.NewReadAccountQueryHandler,
			inject.As(new(sharedDomainBus.QueryHandler)),
		),

		inject.Provide(
			sharedInfraBus.NewSyncQueryBus,
			inject.As(new(sharedDomainBus.QueryBus)),
		),

		inject.Provide(
			command.NewCreateAccountCommand,
			inject.As(new(command.Command)),
			inject.WithName("CreateAccountCommand"),
		),
		inject.Provide(
			command.NewReadAccountCommand,
			inject.As(new(command.Command)),
			inject.WithName("ReadAccountCommand"),
		),
	)
}
