package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"myapi/internal/domain"
)

// MockUserRepositoryはUserRepositoryのモックです
type MockUserRepository struct {
	users []domain.User
	err   error
}

// GetUsersはユーザーリストを取得するモック関数です
func (m *MockUserRepository) GetUsers() ([]domain.User, error) {
	return m.users, m.err
}

// AddUserは新しいユーザーを追加するモック関数です
func (m *MockUserRepository) AddUser(user domain.User) (domain.User, error) {
	if m.err != nil {
		return domain.User{}, m.err
	}
	m.users = append(m.users, user)
	return user, nil
}

func TestGetUserList(t *testing.T) {
	// テスト用のユーザーリスト
	mockUsers := []domain.User{
		{ID: "1", Username: "user1", Email: "user1@example.com"},
		{ID: "2", Username: "user2", Email: "user2@example.com"},
	}

	// モックリポジトリを作成
	mockRepo := &MockUserRepository{
		users: mockUsers,
		err:   nil,
	}

	// ユースケースを作成
	userUsecase := NewUserUsecase(mockRepo)

	// ユーザーリストを取得するテスト
	users, err := userUsecase.GetUserList()

	// テスト結果を検証
	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)
}

func TestAddUser(t *testing.T) {
	// 追加する新しいユーザー
	newUser := domain.User{ID: "3", Username: "user3", Email: "user3@example.com"}

	// モックリポジトリを作成
	mockRepo := &MockUserRepository{
		users: nil,
		err:   nil,
	}

	// ユースケースを作成
	userUsecase := NewUserUsecase(mockRepo)

	// 新しいユーザーを追加するテスト
	addedUser, err := userUsecase.AddUser(newUser)

	// テスト結果を検証
	assert.NoError(t, err)
	assert.Equal(t, newUser, addedUser)
}