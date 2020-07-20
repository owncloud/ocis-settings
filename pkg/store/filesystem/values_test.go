package store

import (
	"testing"

	olog "github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/stretchr/testify/assert"
)

var valueScenarios = []struct {
	name string
	in   struct {
		record   *proto.SettingsValue
		filePath string
	}
	out interface{}
}{
	{
		name: "generic-test-with-system-resource",
		in: struct {
			record   *proto.SettingsValue
			filePath string
		}{
			record: &proto.SettingsValue{
				Id:          value1,
				BundleId:    bundle1,
				SettingId:   setting1,
				AccountUuid: accountUUID1,
				Resource: &proto.Resource{
					Type: proto.Resource_TYPE_SYSTEM,
				},
				Value: &proto.SettingsValue_StringValue{
					StringValue: "lalala",
				},
			},
			filePath: "/var/tmp/herecomesthesun", // TODO replace this with a testing temp file.
		},
		out: nil,
	},
	{
		name: "generic-test-with-file-resource",
		in: struct {
			record   *proto.SettingsValue
			filePath string
		}{
			record: &proto.SettingsValue{
				Id:          value2,
				BundleId:    bundle1,
				SettingId:   setting2,
				AccountUuid: accountUUID1,
				Resource: &proto.Resource{
					Type: proto.Resource_TYPE_FILE,
					Id:   "adfba82d-919a-41c3-9cd1-5a3f83b2bf76",
				},
				Value: &proto.SettingsValue_StringValue{
					StringValue: "tralala",
				},
			},
			filePath: "/var/tmp/herecomesthesun", // TODO replace this with a testing temp file.
		},
		out: nil,
	},
}

func TestWriteSettingsValueToFile(t *testing.T) {
	s := Store{
		dataPath: "/var/tmp/herecomesthesun",
		Logger: olog.NewLogger(
			olog.Color(true),
			olog.Pretty(true),
			olog.Level("info"),
		),
	}
	for i := range valueScenarios {
		index := i
		t.Run(valueScenarios[index].name, func(t *testing.T) {
			t.Parallel()

			filePath := s.buildFilePathForValue(valueScenarios[index].in.record.Id, true)
			if err := s.writeRecordToFile(valueScenarios[index].in.record, filePath); err != nil {
				t.Error(err)
			}
			assert.FileExists(t, filePath)
		})
	}
}
