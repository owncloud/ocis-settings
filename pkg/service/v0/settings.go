package svc

import settings "github.com/owncloud/ocis-settings/pkg/proto/v0"

func generateSettingsBundleAdminRole() *settings.SettingsBundle {
	return &settings.SettingsBundle{
		Identifier: &settings.Identifier{
			Extension: "ocis-settings",
			Bundle: "admin",
		},
		DisplayName: "Admin",
		Settings: []*settings.Setting{
			
		},
	}
}
