package store

import (
	"testing"

	olog "github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

var scenarios = []struct {
	name string
	in   struct {
		record   *proto.SettingsBundle
		filePath string
	}
	out interface{}
}{
	{
		name: "generic-test",
		in: struct {
			record   *proto.SettingsBundle
			filePath string
		}{
			record: &proto.SettingsBundle{
				DisplayName: "test1",
				Identifier: &proto.Identifier{
					AccountUuid: "c4572da7-6142-4383-8fc6-efde3d463036",
					Bundle:      "test-bundle-1",
					Extension:   "test-extension-1",
					Setting:     "test-settings",
				},
				Resource: &proto.Resource{
					Type: proto.ResourceType_FILE,
					Id: "beep",
				},
				Settings: []*proto.Setting{
					{
						Description: "test-desc-1",
						DisplayName: "test-displayname-1",
						Name:        "test-name-1",
						Resource: &proto.Resource{
							Type: proto.ResourceType_FILE,
							Id: "bleep",
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
			filePath: "/var/tmp/herecomesthesun", // TODO replace this with a testing temp file.
		},
		out: nil,
	},
	{
		name: "generic-test-2",
		in: struct {
			record   *proto.SettingsBundle
			filePath string
		}{
			record: &proto.SettingsBundle{
				DisplayName: "test1",
				Identifier: &proto.Identifier{
					AccountUuid: "c4572da7-6142-4383-8fc6-efde3d463034",
					Bundle:      "test-bundle-2",
					Extension:   "test-extension-2",
					Setting:     "test-settings",
				},
				Resource: &proto.Resource{
					Type: proto.ResourceType_FILE,
					Id: "boop",
				},
				Settings: []*proto.Setting{
					{
						Description: "test-desc-2",
						DisplayName: "test-displayname-2",
						Name:        "test-name-2",
						Resource: &proto.Resource{
							Type: proto.ResourceType_FILE,
							Id: "blorg",
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
			filePath: "/var/tmp/herecomesthesun", // TODO replace this with a testing temp file.
		},
		out: nil,
	},
	{
		name: "settings-bundle-without-resource",
		in: struct {
			record   *proto.SettingsBundle
			filePath string
		}{
			record: &proto.SettingsBundle{
				DisplayName: "test1",
				Identifier: &proto.Identifier{
					AccountUuid: "c4572da7-6142-4383-8fc6-efde3d463038",
					Bundle:      "test-global-bundle",
					Extension:   "test-extension-1",
					Setting:     "test-settings",
				},
				Settings: []*proto.Setting{
					{
						Description: "test-desc-2",
						DisplayName: "test-displayname-2",
						Name:        "test-name-2",
						Value: &proto.Setting_IntValue{
							IntValue: &proto.IntSetting{
								Min: 0,
								Max: 42,
							},
						},
					},
				},
			},
			filePath: "/var/tmp/herecomesthesun", // TODO replace this with a testing temp file.
		},
		out: nil,
	},
}

func TestWriteRecordToFile(t *testing.T) {
	s := Store{
		dataPath: "/var/tmp/herecomesthesun",
		Logger: olog.NewLogger(
			olog.Color(true),
			olog.Pretty(true),
			olog.Level("info"),
		),
	}
	for i := range scenarios {
		index := i
		t.Run(scenarios[index].name, func(t *testing.T) {
			t.Parallel()
			filePath := s.buildFilePathForBundle(scenarios[index].in.record.Identifier, scenarios[index].in.record.Resource, true)
			if err := s.writeRecordToFile(scenarios[index].in.record, filePath); err != nil {
				t.Error(err)
			}
		})
	}
}