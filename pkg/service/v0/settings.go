package svc

import settings "github.com/owncloud/ocis-settings/pkg/proto/v0"

// GenerateSettingsBundlesDefaultRoles bootstraps the default roles.
func GenerateSettingsBundlesDefaultRoles() []*settings.SettingsBundle {
	return []*settings.SettingsBundle{
		generateSettingsBundleAdminRole(),
		generateSettingsBundleUserRole(),
		generateSettingsBundleGuestRole(),
	}
}

func generateSettingsBundleAdminRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Id:          "71881883-1768-46bd-a24d-a356a2afdf7f",
		Type:        settings.SettingsBundle_ROLE,
		Extension:   "ocis-settings",
		DisplayName: "Admin role",
		Resource: &settings.Resource{
			Type: settings.Resource_SYSTEM,
		},
		Settings: []*settings.Setting{},
	}
}

func generateSettingsBundleUserRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Id:          "d7beeea8-8ff4-406b-8fb6-ab2dd81e6b11",
		Type:        settings.SettingsBundle_ROLE,
		Extension:   "ocis-settings",
		DisplayName: "User role",
		Resource: &settings.Resource{
			Type: settings.Resource_SYSTEM,
		},
		Settings: []*settings.Setting{},
	}
}

func generateSettingsBundleGuestRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Id:          "38071a68-456a-4553-846a-fa67bf5596cc",
		Type:        settings.SettingsBundle_ROLE,
		Extension:   "ocis-settings",
		DisplayName: "Guest role",
		Resource: &settings.Resource{
			Type: settings.Resource_SYSTEM,
		},
		Settings: []*settings.Setting{},
	}
}
