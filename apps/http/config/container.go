package config

import (
	"github.com/defval/inject"

	_ "github.com/go-sql-driver/mysql" // It loads MySQL driver

	appSQLMigration "credens/apps/http/migration/sql_migration"

	accountAppCreate "credens/libs/accounts/application/create"
	accountAppRead "credens/libs/accounts/application/read"
	accountDomain "credens/libs/accounts/domain"
	accountInfraPersistence "credens/libs/accounts/infrastructure/persistence/mysql"
	accountsInfraQueue "credens/libs/accounts/infrastructure/queue"

	sharedDomainBus "credens/libs/shared/domain/bus"
	sharedInfraBus "credens/libs/shared/infrastructure/bus"
	shareInfraLogging "credens/libs/shared/infrastructure/logging"
	sharedInfraLoggingLogrus "credens/libs/shared/infrastructure/logging/logrus"
	sharedInfraPersistence "credens/libs/shared/infrastructure/persistence"
	sharedInfraQueue "credens/libs/shared/infrastructure/queue"
	sharedInfraQueueRabbitMQ "credens/libs/shared/infrastructure/queue/rabbitmq"
)

func BuildContainer(env Environment) (*inject.Container, error) {
	return inject.New(
		inject.Provide(
			sharedInfraPersistence.NewSQLWrapper(env.Sql.Driver, env.Sql.Url),
			inject.As(new(sharedInfraPersistence.SQLDb)),
		),

		inject.Provide(
			appSQLMigration.NewSQLMigratorWrapper,
			inject.As(new(appSQLMigration.SQLMigrator)),
		),

		inject.Provide(
			sharedInfraQueueRabbitMQ.NewRabbitMQPublisher(
				*sharedInfraQueueRabbitMQ.NewRabbitMQConfig(env.Amqp.Url, "default"),
			),
			inject.As(new(sharedInfraQueue.Publisher)),
		),

		inject.Provide(
			sharedInfraLoggingLogrus.NewLogrusLogger,
			inject.As(new(shareInfraLogging.Logger)),
		),

		inject.Provide(
			accountDomain.NewAccountBuilder,
			inject.As(new(accountDomain.AccountBuilder)),
		),

		inject.Provide(
			accountInfraPersistence.NewMysqlAccountRepository,
			inject.As(new(accountDomain.AccountRepository)),
		),

		inject.Provide(
			accountsInfraQueue.NewRabbitMQEventPublisher,
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
