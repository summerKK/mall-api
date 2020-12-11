package admin

import "github.com/summerKK/mall-api/internal/model"

type UserLoginRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Icon     string `form:"icon" json:"icon"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	NickName string `form:"nickName" json:"nickName" binding:"required"`
	Note     string `form:"note" json:"note"`
}

func (u *UserRegisterRequest) Convert2Model() *model.UmsAdmin {
	return &model.UmsAdmin{
		Username: u.UserName,
		Password: u.Password,
		Icon:     u.Icon,
		Email:    u.Email,
		NickName: u.NickName,
		Note:     u.Note,
	}
}
