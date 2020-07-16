package store

import (
	"testing"

	olog "github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/stretchr/testify/assert"
)

var bundleScenarios = []struct {
	name string
	in   struct {
		record   *proto.SettingsBundle
		filePath string
	}
	out interface{}
}{
	{
		name: "generic-test-file-resource",
		in: struct {
			record   *proto.SettingsBundle
			filePath string
		}{
			record: &proto.SettingsBundle{
				Id:          bundle1,
				Extension:   extension1,
				DisplayName: "test1",
				Resource: &proto.Resource{
					Type: proto.Resource_FILE,
					Id:   "beep",
				},
				Settings: []*proto.Setting{
					{
						Id:          setting1,
						Description: "test-desc-1",
						DisplayName: "test-displayname-1",
						Resource: &proto.Resource{
							Type: proto.Resource_FILE,
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
			filePath: "/var/tmp/herecomesthesun", // TODO replace this with a testing temp file.
		},
		out: nil,
	},
	{
		name: "generic-test-system-resource",
		in: struct {
			record   *proto.SettingsBundle
			filePath string
		}{
			record: &proto.SettingsBundle{
				Id:          bundle2,
				Extension:   extension2,
				DisplayName: "test1",
				Resource: &proto.Resource{
					Type: proto.Resource_SYSTEM,
				},
				Settings: []*proto.Setting{
					{
						Id:          setting2,
						Description: "test-desc-2",
						DisplayName: "test-displayname-2",
						Resource: &proto.Resource{
							Type: proto.Resource_SYSTEM,
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
}

func TestWriteSettingsBundleToFile(t *testing.T) {
	s := Store{
		dataPath: "/var/tmp/herecomesthesun",
		Logger: olog.NewLogger(
			olog.Color(true),
			olog.Pretty(true),
			olog.Level("info"),
		),
	}
	for i := range bundleScenarios {
		index := i
		t.Run(bundleScenarios[index].name, func(t *testing.T) {
			t.Parallel()
			filePath := s.buildFilePathForBundle(bundleScenarios[index].in.record.Id, true)
			if err := s.writeRecordToFile(bundleScenarios[index].in.record, filePath); err != nil {
				t.Error(err)
			}
			assert.FileExists(t, filePath)
		})
	}
}
