package svc

import "github.com/owncloud/ocis-settings/pkg/proto/v0"

func (g Service) hasPermission(assignments []*proto.UserRoleAssignment, resource *proto.Resource, operation proto.PermissionSetting_Operation) bool {
	return true
}
