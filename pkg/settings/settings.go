package settings

import (
	"github.com/owncloud/ocis-settings/pkg/config"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

var (
	// Registry uses the strategy pattern as a registry
	Registry = map[string]RegisterFunc{}
)

// RegisterFunc stores store constructors
type RegisterFunc func(*config.Config) Manager

// Manager combines service interfaces for abstraction of storage implementations
type Manager interface {
	BundleManager
	ValueManager
	RoleAssignmentManager
}

// BundleManager is a bundle service interface for abstraction of storage implementations
type BundleManager interface {
	ListBundles(accountUUID string, bundleType proto.SettingsBundle_Type) ([]*proto.SettingsBundle, error)
	ReadBundle(bundleId string) (*proto.SettingsBundle, error)
	WriteBundle(bundle *proto.SettingsBundle) (*proto.SettingsBundle, error)
	AddSettingToBundle(bundleId string, setting *proto.Setting) (*proto.Setting, error)
	RemoveSettingFromBundle(bundleId, settingId string) error
}

// ValueManager is a value service interface for abstraction of storage implementations
type ValueManager interface {
	ListValues(bundleId, accountUUID string) ([]*proto.SettingsValue, error)
	ReadValue(valueId string) (*proto.SettingsValue, error)
	WriteValue(value *proto.SettingsValue) (*proto.SettingsValue, error)
}

// RoleAssignmentManager is a role assignment service interface for abstraction of storage implementations
type RoleAssignmentManager interface {
	ListRoleAssignments(accountUUID string) ([]*proto.UserRoleAssignment, error)
	WriteRoleAssignment(assignment *proto.UserRoleAssignment) (*proto.UserRoleAssignment, error)
	RemoveRoleAssignment(assignmentId string) error
}
