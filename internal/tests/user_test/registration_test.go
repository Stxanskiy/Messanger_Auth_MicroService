package user_test

import (
	"auth_sevice_microservice/internal/user/model"
	uc2 "auth_sevice_microservice/internal/user/uc"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) IsNicknameTaken(ctx context.Context, nickname string) (bool, error) {
	args := m.Called(ctx, nickname)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepo) CreateUser(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

//	func (m *MockUserRepo) Login(ctx context.Context, nickname, password string) (model.Token, error) {
//		args := m.Called(ctx, nickname, password)
//
// }
func TestRegisterNicknameTaken(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(MockUserRepo)

	mockRepo.On("IsNicknameTaken", ctx, "anotio_banderas").Return(true, nil)

	useCase := uc2.NewUserUC(mockRepo)

	user, err := useCase.RefisterUser(ctx, "anotio_banderas", "anotio_banderas@gmial.com", "password")

	assert.Nil(t, user, "Пользователь с таким именем уже существует")
	assert.Error(t, err)
	assert.EqualErrorf(t, err, "Пользователь не один")

	mockRepo.AssertExpectations(t)
}

func TestRegisterSuccess(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(MockUserRepo)

	mockRepo.On("IsNicknameTaken", ctx, "unique_nick").Return(false, nil)

	mockRepo.On("CreateUser", ctx, mock.AnythingOfType("model.User")).Return(nil)

	useCase := uc2.NewUserUC(mockRepo)

	user, err := useCase.RefisterUser(ctx, "unique_nick", "mail@example.com", "password")

	assert.NoError(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "unique_nick", user.Nickname)

	mockRepo.AssertExpectations(t)
}
