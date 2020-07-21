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
	ListBundles(bundleType proto.SettingsBundle_Type) ([]*proto.SettingsBundle, error)
	ReadBundle(bundleID string) (*proto.SettingsBundle, error)
	WriteBundle(bundle *proto.SettingsBundle) (*proto.SettingsBundle, error)
	ReadSetting(settingID string) (*proto.Setting, error)
	AddSettingToBundle(bundleID string, setting *proto.Setting) (*proto.Setting, error)
	RemoveSettingFromBundle(bundleID, settingID string) error
}

// ValueManager is a value service interface for abstraction of storage implementations
type ValueManager interface {
	ListValues(bundleID, accountUUID string) ([]*proto.SettingsValue, error)
	ReadValue(valueID string) (*proto.SettingsValue, error)
	WriteValue(value *proto.SettingsValue) (*proto.SettingsValue, error)
}

// RoleAssignmentManager is a role assignment service interface for abstraction of storage implementations
type RoleAssignmentManager interface {
	ListRoleAssignments(accountUUID string) ([]*proto.UserRoleAssignment, error)
	WriteRoleAssignment(accountUUID, roleID string) (*proto.UserRoleAssignment, error)
	RemoveRoleAssignment(assignmentID string) error
}
