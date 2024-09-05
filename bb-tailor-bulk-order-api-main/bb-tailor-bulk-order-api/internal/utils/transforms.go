package utils

import "git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"

func TransformRolesToStringArray(roles []*models.Role) []string {
	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}
	return roleNames
}
