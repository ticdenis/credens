package handler_test

import (
	"credens/apps/http/handler"
	"credens/libs/accounts/application/create"
	"credens/libs/shared/domain/bus"
	"credens/tests/apps/http/mocks"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type CreateAccountPostHandlerSuite struct {
	suite.Suite

	accountRepository *mocks.AccountRepositoryMock
	eventPublisher    *mocks.EventPublisherMock
	commandBus        *mocks.TestCommandBus

	sut *handler.CreateAccountPostHandler
}

func (suite *CreateAccountPostHandlerSuite) SetupTest() {
	suite.accountRepository = new(mocks.AccountRepositoryMock)
	suite.eventPublisher = new(mocks.EventPublisherMock)
	suite.commandBus = mocks.NewTestCommandBus([]bus.CommandHandler{
		create.NewCreateAccountCommandHandler(suite.accountRepository, suite.eventPublisher),
	})

	suite.sut = handler.NewCreateAccountPostHandler(*suite.commandBus)

	gin.SetMode(gin.TestMode)
}

func (suite *CreateAccountPostHandlerSuite) TearDownTest() {
	suite.accountRepository.AssertExpectations(suite.T())
	suite.eventPublisher.AssertExpectations(suite.T())
}

func (suite *CreateAccountPostHandlerSuite) TestShouldCreateAccount() {
	suite.accountRepository.On("Add", mock.AnythingOfType("*domain.Account")).Return(nil).Once()
	suite.eventPublisher.On("Publish", mock.AnythingOfType("[]bus.Event")).Once()

	actualRes, actualErr := suite.sut.Handle(mocks.MakeJSONRequestAndGetGinContext(
		http.MethodPost,
		"https://github.com/ticdenis/credens",
		fmt.Sprintf(
			`{"name":"%s","username":"%s","password":"%s"}`,
			"go", "gopher", "secret",
		),
	))

	assert.Nil(suite.T(), actualErr)
	assert.Equal(suite.T(), 201, actualRes.Status)
	assert.Contains(suite.T(), actualRes.Content, "data")
	assert.Contains(suite.T(), actualRes.Content["data"], "type")
	assert.Contains(suite.T(), actualRes.Content["data"].(map[string]interface{})["type"], "accounts")
	assert.Contains(suite.T(), actualRes.Content["data"], "id")
	assert.NotContains(suite.T(), actualRes.Content["data"], "attributes")
}

func TestCreateAccountPostHandlerSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountPostHandlerSuite))
}
