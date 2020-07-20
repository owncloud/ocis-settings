package svc

import settings "github.com/owncloud/ocis-settings/pkg/proto/v0"

const (
	BundleUUIDRoleAdmin = "71881883-1768-46bd-a24d-a356a2afdf7f"
	BundleUUIDRoleUser  = "d7beeea8-8ff4-406b-8fb6-ab2dd81e6b11"
	BundleUUIDRoleGuest = "38071a68-456a-4553-846a-fa67bf5596cc"
)

// generateSettingsBundlesDefaultRoles bootstraps the default roles.
func generateSettingsBundlesDefaultRoles() []*settings.SettingsBundle {
	return []*settings.SettingsBundle{
		generateSettingsBundleAdminRole(),
		generateSettingsBundleUserRole(),
		generateSettingsBundleGuestRole(),
	}
}

func generateSettingsBundleAdminRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Id:          BundleUUIDRoleAdmin,
		Name:        "admin",
		Type:        settings.SettingsBundle_TYPE_ROLE,
		Extension:   "ocis-roles",
		DisplayName: "Admin role",
		Resource: &settings.Resource{
			Type: settings.Resource_TYPE_SYSTEM,
		},
		Settings: []*settings.Setting{},
	}
}

func generateSettingsBundleUserRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Id:          BundleUUIDRoleUser,
		Name:        "user",
		Type:        settings.SettingsBundle_TYPE_ROLE,
		Extension:   "ocis-roles",
		DisplayName: "User role",
		Resource: &settings.Resource{
			Type: settings.Resource_TYPE_SYSTEM,
		},
		Settings: []*settings.Setting{},
	}
}

func generateSettingsBundleGuestRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Id:          BundleUUIDRoleGuest,
		Name:        "guest",
		Type:        settings.SettingsBundle_TYPE_ROLE,
		Extension:   "ocis-roles",
		DisplayName: "Guest role",
		Resource: &settings.Resource{
			Type: settings.Resource_TYPE_SYSTEM,
		},
		Settings: []*settings.Setting{},
	}
}
