package repository

import (
	"github.com/google/uuid"
	"myapi/internal/domain"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	AddUser(user domain.User) (domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userRepo struct {
	users []domain.User
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (ur *userRepo) AddUser(user domain.User) (domain.User, error) {
	user.ID = uuid.New().String()
	ur.users = append(ur.users, user)
	return user, nil
}

func (ur *userRepo) GetUsers() ([]domain.User, error) {
	return ur.users, nil
}

// -----------------------------

// MySQLRepository は MySQL データベースへのアクセスを提供します
type MySQLRepository struct {
	DB *gorm.DB
}

// NewMySQLRepository は MySQLRepository の新しいインスタンスを作成します
func NewMySQLRepository(db *gorm.DB) UserRepository {
	return &MySQLRepository{
		DB: db,
	}
}

func (mr *MySQLRepository) AddUser(user domain.User) (domain.User, error) {
    if err := mr.DB.Create(&user).Error; err != nil {
        return domain.User{}, err
    }
    return user, nil
}

func (mr *MySQLRepository) GetUsers() ([]domain.User, error) {
    var users []domain.User
    if err := mr.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}



// AddUser は新しいユーザーを MySQL データベースに追加します
// func (m *MySQLRepository) AddUser(user domain.User) (domain.User, error) {
// 	// MySQLのINSERT文を作成し、ユーザー情報をデータベースに挿入します
// 	query := "INSERT INTO users (email, created_at, updated_at) VALUES (?,?,?)"
// 	result, err := m.DB.Exec(query, user.Email, time.Now(), time.Now() )
// 	if err != nil {
// 		return domain.User{}, err
// 	}

// 	// ユーザー情報を返す（例えば、IDが自動生成される場合など）
// 	tmpID, _ := result.LastInsertId()
// 	user.ID = strconv.FormatInt(tmpID, 10)
// 	return user, nil
// }

// GetUsers は全てのユーザーを取得します
// func (m *MySQLRepository) GetUsers() ([]domain.User, error) {
//     query := "SELECT id, email, created_at FROM users"
//     rows, err := m.DB.Query(query)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var users []domain.User
//     for rows.Next() {
//         var user domain.User
//         var createdAtString string // created_at カラムを文字列として格納する変数

//         if err := rows.Scan(&user.ID, &user.Email, &createdAtString); err != nil {
//             return nil, err
//         }

//         createdAtTime, err := stringToTime(createdAtString)
//         if err != nil {
//             return nil, err
//         }
//         user.CreatedAt = createdAtTime

//         users = append(users, user)
//     }

//     if err := rows.Err(); err != nil {
//         return nil, err
//     }

//     return users, nil
// }

func stringToTime(createdAtString string) (time.Time, error) {
    // 実際のフォーマットに合わせてパースの方法を指定
    createdAtTime, err := time.Parse("2006-01-02T15:04:05.999999-07:00", createdAtString)
    if err != nil {
        return time.Time{}, err
    }
    return createdAtTime, nil
}