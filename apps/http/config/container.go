package config

import (
	"credens/apps/http/migration/sql_migration"
	"github.com/defval/inject"

	accountAppCreate "credens/libs/accounts/application/create"
	accountAppRead "credens/libs/accounts/application/read"
	accountDomain "credens/libs/accounts/domain"
	accountInfraPersistence "credens/libs/accounts/infrastructure/persistence/mysql"

	sharedDomainBus "credens/libs/shared/domain/bus"
	sharedInfraBus "credens/libs/shared/infrastructure/bus"
	shareInfraLogging "credens/libs/shared/infrastructure/logging"
	sharedInfraLoggingLogrus "credens/libs/shared/infrastructure/logging/logrus"
	sharedInfraPersistence "credens/libs/shared/infrastructure/persistence"
)

func BuildContainer(env Environment) (*inject.Container, error) {
	return inject.New(
		inject.Provide(
			NewMySQLDBWrapper(env),
			inject.As(new(sharedInfraPersistence.SQLDb)),
		),

		inject.Provide(
			sql_migration.NewSQLMigratorWrapper,
			inject.As(new(sql_migration.SQLMigrator)),
		),

		inject.Provide(
			sharedInfraLoggingLogrus.NewLogrusLogger,
			inject.As(new(shareInfraLogging.Logger)),
		),

		inject.Provide(
			accountInfraPersistence.NewMysqlAccountRepository,
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
