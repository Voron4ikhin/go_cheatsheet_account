package mapper

import (
	"github.com/voron4ikhin/go_cheatsheet_account/internal/model"
	repomodel "github.com/voron4ikhin/go_cheatsheet_account/internal/repository/model"
)

func UserToRepoUser(user model.User) repomodel.User {
	return repomodel.User{
		ID:         user.ID,
		Login:      user.Login,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func RepoUserToUser(user repomodel.User) model.User {
	res := model.User{
		ID:         user.ID,
		Login:      user.Login,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
	return res
}

func RepoUsersToUsers(users []repomodel.User) []model.User {
	res := make([]model.User, len(users))
	for i, user := range users {
		res[i] = RepoUserToUser(user)
	}
	return res
}

func UpdateUserToRepoUser(user model.UpdateUser) repomodel.User {
	return repomodel.User{
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
	}
}
