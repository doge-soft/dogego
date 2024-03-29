package serializer

import "{{cookiecutter.app_name}}/models"

// User 用户序列化器
type User struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	NickName    string `json:"nick_name"`
	Bio         string `json:"bio"`
	Role        string `json:"role"`
	Avatar      string `json:"avatar"`
	CreatedAt   int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		NickName:    user.NickName,
		Bio:         user.Bio,
		Role:        user.Role,
		Avatar:      user.Avatar,
		CreatedAt:   user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user models.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
