package auth_service

import "github.com/luanluanxu/go-gin-app/internal/my_app/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
