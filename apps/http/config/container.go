package config

import (
	"database/sql"
	"github.com/defval/inject"

	accountAppCreate "credens/libs/accounts/application/create"
	accountAppRead "credens/libs/accounts/application/read"
	accountDomain "credens/libs/accounts/domain"
	accountInfraPersistence "credens/libs/accounts/infrastructure/persistence/mysql"

	sharedDomainBus "credens/libs/shared/domain/bus"
	sharedInfraBus "credens/libs/shared/infrastructure/bus"
	shareInfraLogging "credens/libs/shared/infrastructure/logging"
	sharedInfraLoggingLogrus "credens/libs/shared/infrastructure/logging/logrus"
)

func BuildContainer(env Environment, db *sql.DB) (*inject.Container, error) {
	return inject.New(
		inject.Provide(
			sharedInfraLoggingLogrus.NewLogrusLogger,
			inject.As(new(shareInfraLogging.Logger)),
		),

		inject.Provide(
			accountInfraPersistence.NewMysqlAccountRepository(db),
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
	)
}
