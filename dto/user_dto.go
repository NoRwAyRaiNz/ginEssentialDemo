package dto

import "ginEssential/Model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user Model.User) UserDto { //仅返回部分信息
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
