package usecase

import (
	"myapi/internal/domain"
	"myapi/internal/repository"
)

// UserUsecase はユーザー関連のユースケースを定義します
type UserUsecase interface {
	GetUserList() ([]domain.User, error)
	AddUser(user domain.User) (domain.User, error)
}

// userUsecase は UserUsecase インターフェースの実装です
type userUsecase struct {
	userRepository repository.UserRepository // ユーザーのリポジトリを利用する例
}

// NewUserUsecase は新しい UserUsecase インスタンスを作成します
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

// GetUserList はユーザーリストを取得するメソッドです
func (uc *userUsecase) GetUserList() ([]domain.User, error) {
	// UserRepository を使ってユーザーリストを取得する例
	return uc.userRepository.GetUsers()
}

// AddUser は新しいユーザーを追加するメソッドです
func (uc *userUsecase) AddUser(user domain.User) (domain.User, error) {
	// UserRepository を使って新しいユーザーを追加する例
	return uc.userRepository.AddUser(user)
}
