package service

import (
	"errors"
	"time"

	"github.com/papu/gAgent/internal/models"
)

// UserService 处理用户相关的业务逻辑
type UserService struct{}

// NewUserService 创建新的用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(username, email, password string) (*models.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("用户名、邮箱和密码不能为空")
	}

	user := &models.User{
		Username:  username,
		Email:     email,
		Password:  password, // 注意：实际应用中应该对密码进行加密
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}
