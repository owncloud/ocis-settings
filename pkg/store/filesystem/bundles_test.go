package store

import (
	"testing"

	olog "github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/stretchr/testify/assert"
)

var bundleScenarios = []struct {
	name   string
	bundle *proto.SettingsBundle
}{
	{
		name: "generic-test-file-resource",
		bundle: &proto.SettingsBundle{
			Id:          bundle1,
			Type:        proto.SettingsBundle_TYPE_DEFAULT,
			Extension:   extension1,
			DisplayName: "test1",
			Resource: &proto.Resource{
				Type: proto.Resource_TYPE_FILE,
				Id:   "beep",
			},
			Settings: []*proto.Setting{
				{
					Id:          setting1,
					Description: "test-desc-1",
					DisplayName: "test-displayname-1",
					Resource: &proto.Resource{
						Type: proto.Resource_TYPE_FILE,
						Id:   "bleep",
					},
					Value: &proto.Setting_IntValue{
						IntValue: &proto.IntSetting{
							Min: 0,
							Max: 42,
						},
					},
				},
			},
		},
	},
	{
		name: "generic-test-system-resource",
		bundle: &proto.SettingsBundle{
			Id:          bundle2,
			Type:        proto.SettingsBundle_TYPE_DEFAULT,
			Extension:   extension2,
			DisplayName: "test1",
			Resource: &proto.Resource{
				Type: proto.Resource_TYPE_SYSTEM,
			},
			Settings: []*proto.Setting{
				{
					Id:          setting2,
					Description: "test-desc-2",
					DisplayName: "test-displayname-2",
					Resource: &proto.Resource{
						Type: proto.Resource_TYPE_SYSTEM,
					},
					Value: &proto.Setting_IntValue{
						IntValue: &proto.IntSetting{
							Min: 0,
							Max: 42,
						},
					},
				},
			},
		},
	},
	{
		name: "generic-test-role-bundle",
		bundle: &proto.SettingsBundle{
			Id:          bundle3,
			Type:        proto.SettingsBundle_TYPE_ROLE,
			Extension:   extension1,
			DisplayName: "Role1",
			Resource: &proto.Resource{
				Type: proto.Resource_TYPE_SYSTEM,
			},
			Settings: []*proto.Setting{
				{
					Id:          setting3,
					Description: "test-desc-3",
					DisplayName: "test-displayname-3",
					Resource: &proto.Resource{
						Type: proto.Resource_TYPE_SETTING,
						Id:   setting1,
					},
					Value: &proto.Setting_PermissionValue{
						PermissionValue: &proto.PermissionSetting{
							Operation:  proto.PermissionSetting_OPERATION_READ,
							Constraint: proto.PermissionSetting_CONSTRAINT_OWN,
						},
					},
				},
			},
		},
	},
}

func TestSettingsBundles(t *testing.T) {
	s := Store{
		dataPath: dataRoot,
		Logger: olog.NewLogger(
			olog.Color(true),
			olog.Pretty(true),
			olog.Level("info"),
		),
	}

	// write bundles
	for i := range bundleScenarios {
		index := i
		t.Run(bundleScenarios[index].name, func(t *testing.T) {
			filePath := s.buildFilePathForBundle(bundleScenarios[index].bundle.Id, true)
			if err := s.writeRecordToFile(bundleScenarios[index].bundle, filePath); err != nil {
				t.Error(err)
			}
			assert.FileExists(t, filePath)
		})
	}

	// check that ListBundles only returns bundles with type DEFAULT
	bundles, err := s.ListBundles(proto.SettingsBundle_TYPE_DEFAULT)
	if err != nil {
		t.Error(err)
	}
	for i := range bundles {
		assert.Equal(t, proto.SettingsBundle_TYPE_DEFAULT, bundles[i].Type)
	}

	// check that ListRoles only returns bundles with type ROLE
	roles, err := s.ListBundles(proto.SettingsBundle_TYPE_ROLE)
	if err != nil {
		t.Error(err)
	}
	for i := range roles {
		assert.Equal(t, proto.SettingsBundle_TYPE_ROLE, roles[i].Type)
	}

	burnRoot()
}
