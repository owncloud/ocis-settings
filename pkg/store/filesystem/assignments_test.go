package store

import (
	"path/filepath"
	"testing"

	olog "github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/stretchr/testify/assert"
)

var (
	einstein = "a4d07560-a670-4be9-8d60-9b547751a208"
	marie    = "3c054db3-eec1-4ca4-b985-bc56dcf560cb"

	logger = olog.NewLogger(
		olog.Color(true),
		olog.Pretty(true),
		olog.Level("info"),
	)

	assignmentScenarios = []struct {
		Bundle      *proto.SettingsBundle
		AccountUUID string
	}{
		{
			AccountUUID: einstein,
			Bundle: &proto.SettingsBundle{
				Id:          "f36db5e6-a03c-40df-8413-711c67e40b47",
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
		},
	}
	// TODO: add red tests based on specs. For instance, should be having 2 role assignments to the same accountID
	// pointing to the same role be allowed?
)

func TestRoleAssignments(t *testing.T) {
	s := Store{
		dataPath: dataRoot,
		Logger:   logger,
	}

	// write role assignments
	for i := range assignmentScenarios {
		res, err := s.WriteBundle(assignmentScenarios[i].Bundle)
		assert.NoError(t, err)

		roleAssignment, err := s.WriteRoleAssignment(assignmentScenarios[i].AccountUUID, res.Id)

		assert.NoError(t, err)
		assert.FileExists(t, filepath.Join(dataRoot, "assignments", roleAssignment.Id+".json"))
	}

	// list roles
	for i := range assignmentScenarios {
		// list role assignment for the current account
		roleAssignments, err := s.ListRoleAssignments(assignmentScenarios[i].AccountUUID)
		if err != nil {
			t.Error(err)
		}

		// a role is stored as a SettingBundle of type "Role", so once we get the assignment we need
		// to fetch the SettingBundle by its ID.
		for j := range roleAssignments {
			role, err := s.ReadBundle(roleAssignments[j].RoleId)

			assert.NoError(t, err)
			assert.Equal(t, role.GetName(), assignmentScenarios[i].Bundle.GetName())
		}
	}

	// Remove Assignment
	for i := range assignmentScenarios {
		roleAssignments, err := s.ListRoleAssignments(assignmentScenarios[i].AccountUUID)
		if err != nil {
			t.Error(err)
		}

		for j := range roleAssignments {
			err = s.RemoveRoleAssignment(roleAssignments[j].Id)
			if err != nil {
				t.Error(err)
			}
			assert.NoFileExists(t, filepath.Join(dataRoot, "assignments", roleAssignments[i].Id+".json"))
			assert.FileExists(t, filepath.Join(dataRoot, "bundles", assignmentScenarios[i].Bundle.Id+".json"))
		}
	}
	burnRoot()
}
