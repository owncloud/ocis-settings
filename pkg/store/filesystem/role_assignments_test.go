package store

import (
	"os"
	"path/filepath"
	"testing"

	olog "github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/stretchr/testify/assert"
)

var (
	dataRoot = "/var/tmp/herecomesthesun"
	logger   = olog.NewLogger(
		olog.Color(true),
		olog.Pretty(true),
		olog.Level("info"),
	)

	bundles = []*proto.SettingsBundle{
		{
			Type:        proto.SettingsBundle_TYPE_ROLE,
			DisplayName: "test role - reads | update",
			Name:        "TEST_ROLE",
			Extension:   "ocis-settings",
			Resource: &proto.Resource{
				Type: proto.Resource_TYPE_BUNDLE,
			},
			Settings: []*proto.Setting{
				{
					Name: "update",
					Value: &proto.Setting_PermissionValue{
						PermissionValue: &proto.PermissionSetting{
							Operation: proto.PermissionSetting_OPERATION_UPDATE,
						},
					},
				},
				{
					Name: "read",
					Value: &proto.Setting_PermissionValue{
						PermissionValue: &proto.PermissionSetting{
							Operation: proto.PermissionSetting_OPERATION_READ,
						},
					},
				},
			},
		},
	}
)

func TestRoleAssignments(t *testing.T) {
	// start from a zero state on the store datapath
	os.RemoveAll(filepath.Join(dataRoot, "role-assignments"))
	os.RemoveAll(filepath.Join(dataRoot, "bundles"))
	s := Store{
		dataPath: dataRoot,
		Logger:   logger,
	}

	res, err := s.WriteBundle(bundles[0])
	if err != nil {
		t.Error(err)
	}

	roleAssignment, err := s.WriteRoleAssignment(&proto.UserRoleAssignment{
		AccountUuid: "a4d07560-a670-4be9-8d60-9b547751a208",
		RoleId:      res.Id,
	})
	if err != nil {
		t.Error(err)
	}

	assert.FileExists(t, filepath.Join(dataRoot, "role-assignments", roleAssignment.Id+".json"))
}
