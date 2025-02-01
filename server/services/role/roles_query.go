package role

import (
	"server/database/repository"
	"server/models"
	"slices"
)

func (rs *RoleService) GetRoleByUserAndBoard(userID, boardID uint) (models.BoardRole, error) {
	userBoardRole, err := rs.db.UserBoardRoleRepository.GetFirst(repository.WithWhere("user_id = ? AND board_id = ?", userID, boardID))
	if err != nil {
		return models.BoardRole{}, err
	}

	boardRole, err := rs.db.BoardRoleRepository.GetFirst(repository.WithWhere("id = ?", userBoardRole.BoardRoleID))
	if err != nil {
		return models.BoardRole{}, err
	}

	return boardRole, nil
}

type UserWithAppRole struct {
	models.User
	AppRole string `json:"app_role"`
}

func (rs *RoleService) GetUsersWithAccessToBoard(boardID uint) ([]UserWithAppRole, error) {
	perms, err := rs.db.UserBoardRoleRepository.GetAll(
		repository.WithWhere("board_id = ?", boardID),
		repository.WithPreload("User"),
		repository.WithPreload("BoardRole"),
	)
	if err != nil {
		return nil, err
	}

	users := make([]UserWithAppRole, 0)
	for _, perm := range perms {
		users = append(users, UserWithAppRole{
			User:    perm.User,
			AppRole: roleMap[perm.BoardRole.Name].Name,
		})
	}

	globalAdmins, err := rs.db.UserRepository.GetAll(repository.WithWhere("role = ?", models.RoleAdmin))
	if err != nil {
		return nil, err
	}

	globalAdminsWithRoles := make([]UserWithAppRole, len(globalAdmins))
	for i, globalAdmin := range globalAdmins {
		globalAdminsWithRoles[i] = UserWithAppRole{
			User:    globalAdmin,
			AppRole: AdminRole.Name,
		}
	}

	for _, globalAdmin := range globalAdminsWithRoles {
		if !slices.Contains(users, globalAdmin) {
			users = append(users, globalAdmin)
		}
	}

	return users, nil
}
