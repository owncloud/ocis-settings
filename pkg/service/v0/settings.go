package svc

import settings "github.com/owncloud/ocis-settings/pkg/proto/v0"

func generateSettingsBundlesDefaultRoles() []*settings.SettingsBundle {
	return []*settings.SettingsBundle{
		generateSettingsBundleAdminRole(),
		generateSettingsBundleUserRole(),
		generateSettingsBundleGuestRole(),
	}
}

func generateSettingsBundleAdminRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Identifier: &settings.Identifier{
			Extension: "ocis-settings",
			Bundle: "admin",
		},
		DisplayName: "Admin",
		Settings: []*settings.Setting{
		},
		Resource: &settings.Resource{
			Type: settings.Resource_SYSTEM,
		},
	}
}

func generateSettingsBundleUserRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Identifier: &settings.Identifier{
			Extension: "ocis-settings",
			Bundle: "user",
		},
		DisplayName: "User",
		Settings: []*settings.Setting{
		},
		Resource: &settings.Resource{
			Type: settings.Resource_SYSTEM,
		},
	}
}

func generateSettingsBundleGuestRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Identifier: &settings.Identifier{
			Extension: "ocis-settings",
			Bundle: "guest",
		},
		DisplayName: "Guest",
		Settings: []*settings.Setting{
		},
		Resource: &settings.Resource{
			Type: settings.Resource_SYSTEM,
		},
	}
}
