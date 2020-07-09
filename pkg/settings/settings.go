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
	ReadBundle(identifier *proto.Identifier, resource *proto.Resource) (*proto.SettingsBundle, error)
	WriteBundle(bundle *proto.SettingsBundle) (*proto.SettingsBundle, error)
	ListBundles(identifier *proto.Identifier, resource *proto.Resource) ([]*proto.SettingsBundle, error)
}

// ValueManager is a value service interface for abstraction of storage implementations
type ValueManager interface {
	ReadValue(identifier *proto.Identifier, resource *proto.Resource) (*proto.SettingsValue, error)
	WriteValue(value *proto.SettingsValue) (*proto.SettingsValue, error)
	ListValues(identifier *proto.Identifier, resource *proto.Resource) ([]*proto.SettingsValue, error)
}

// RoleAssignmentManager is a role assignment service interface for abstraction of storage implementations
type RoleAssignmentManager interface {
	ListRoleAssignments(assignment *proto.RoleAssignmentIdentifier) (*proto.UserRoleAssignments, error)
	WriteRoleAssignment(assignment *proto.RoleAssignmentIdentifier) error
	DeleteRoleAssignment(assignment *proto.RoleAssignmentIdentifier) error
}
