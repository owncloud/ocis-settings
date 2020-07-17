package svc

import settings "github.com/owncloud/ocis-settings/pkg/proto/v0"

const (
	BundleUuidRoleAdmin = "71881883-1768-46bd-a24d-a356a2afdf7f"
	BundleUuidRoleUser  = "d7beeea8-8ff4-406b-8fb6-ab2dd81e6b11"
	BundleUuidRoleGuest = "38071a68-456a-4553-846a-fa67bf5596cc"
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
		Id:          BundleUuidRoleAdmin,
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
		Id:          BundleUuidRoleUser,
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
		Id:          BundleUuidRoleGuest,
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
