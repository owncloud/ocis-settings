// Package store implements the go-micro store interface
package store

import (
	"os"

	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

// ListRoleAssignments loads and returns all role assignments matching the given assignment identifier.
func (s Store) ListRoleAssignments(assignment *proto.RoleAssignmentIdentifier) (*proto.UserRoleAssignments, error) {
	filePath := s.buildFilePathFromRoleAssignmentArgs(assignment.AccountUuid, false)
	assignments := &proto.UserRoleAssignments{
		Assignments: []*proto.RoleAssignmentIdentifier{},
	}
	err := s.parseRecordFromFile(assignments, filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	// TODO: filter assignments by optional fields in the assignment identifier.
	// E.g. only return the assignments that match a certain role name or a certain resource.
	return assignments, nil
}

// WriteRoleAssignment appends the given role assignment to the existing assignments of the respective account.
func (s Store) WriteRoleAssignment(assignment *proto.RoleAssignmentIdentifier) error {
	assignments, err := s.ListRoleAssignments(assignment)
	if err != nil {
		return err
	}

	// check if role assignment exists
	for _, a := range assignments.Assignments {
		if equalAssignments(assignment, a) {
			return nil
		}
	}

	// role assignment doesn't exist, yet
	assignments.Assignments = append(assignments.Assignments, assignment)
	filePath := s.buildFilePathFromRoleAssignmentArgs(assignment.AccountUuid, true)
	err = s.writeRecordToFile(assignments, filePath)
	return err
}

// DeleteRoleAssignment deletes the given role assignment from the existing assignments of the respective account.
func (s Store) DeleteRoleAssignment(assignment *proto.RoleAssignmentIdentifier) error {
	assignments, err := s.ListRoleAssignments(assignment)
	if err != nil {
		return err
	}

	// delete role assignment if it exists
	for index, a := range assignments.Assignments {
		if equalAssignments(assignment, a) {
			assignments.Assignments = append(assignments.Assignments[:index], assignments.Assignments[index+1:]...)
			filePath := s.buildFilePathFromRoleAssignmentArgs(assignment.AccountUuid, true)
			err = s.writeRecordToFile(assignments, filePath)
			return err
		}
	}
	return nil
}

// equalAssignments checks if the two given role assignments have the same properties.
func equalAssignments(a1, a2 *proto.RoleAssignmentIdentifier) bool {
	return a1.AccountUuid == a2.AccountUuid &&
		a1.Role == a2.Role &&
		a1.ResourceType == a2.ResourceType &&
		a1.ResourceValue == a2.ResourceValue
}
