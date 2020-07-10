package svc

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

var (
	regexForKeys = regexp.MustCompile(`^[A-Za-z0-9\-_]*$`)
	keyRule      = []validation.Rule{
		validation.Required,
		validation.Match(regexForKeys),
	}
	settingKeyRule = []validation.Rule{
		validation.Required,
		validation.Match(regexForKeys),
	}
	accountUUIDRule = []validation.Rule{
		validation.Required,
		is.UUID,
	}
)

func validateSaveSettingsBundle(req *proto.SaveSettingsBundleRequest) error {
	if err := validateBundleIdentifier(req.SettingsBundle.Identifier); err != nil {
		return err
	}
	return validation.ValidateStruct(
		req.SettingsBundle,
		validation.Field(&req.SettingsBundle.DisplayName, validation.Required),
		validation.Field(&req.SettingsBundle.Settings, validation.Required),
	)
}

func validateGetSettingsBundle(req *proto.GetSettingsBundleRequest) error {
	return validateBundleIdentifier(req.Identifier)
}

func validateListSettingsBundles(req *proto.ListSettingsBundlesRequest) error {
	return validation.ValidateStruct(
		req.Identifier,
		validation.Field(&req.Identifier.Extension, validation.Match(regexForKeys)),
	)
}

func validateSaveSettingsValue(req *proto.SaveSettingsValueRequest) error {
	return validateValueIdentifier(req.SettingsValue.Identifier)
}

func validateGetSettingsValue(req *proto.GetSettingsValueRequest) error {
	return validateValueIdentifier(req.Identifier)
}

func validateListSettingsValues(req *proto.ListSettingsValuesRequest) error {
	return validation.ValidateStruct(
		req.Identifier,
		validation.Field(&req.Identifier.AccountUuid, is.UUID),
		validation.Field(&req.Identifier.Extension, validation.Match(regexForKeys)),
		validation.Field(&req.Identifier.Extension, validation.When(req.Identifier.Bundle != "", validation.Required)),
		validation.Field(&req.Identifier.Bundle, validation.Match(regexForKeys)),
		validation.Field(&req.Identifier.Bundle, validation.When(req.Identifier.Setting != "", validation.Required)),
		validation.Field(&req.Identifier.Setting, validation.Match(regexForKeys)),
	)
}

func validateListRoleAssignments(req *proto.ListRoleAssignmentsRequest) error {
	return validation.ValidateStruct(
		req.Assignment,
		validation.Field(&req.Assignment.AccountUuid, is.UUID),
		validation.Field(&req.Assignment.Role, validation.When(req.Assignment.Role != "", validation.Match(regexForKeys))),
		// TODO: if resource is present (it's not required), how do we validate the resource types and ids?
	)
}

func validateAssignRoleToUser(req *proto.AssignRoleToUserRequest) error {
	return validateRoleAssignmentIdentifier(req.Assignment)
}

func validateRemoveRoleFromUser(req *proto.RemoveRoleFromUserRequest) error {
	return validateRoleAssignmentIdentifier(req.Assignment)
}

// validateBundleIdentifier is an internal helper for request validation.
func validateBundleIdentifier(identifier *proto.Identifier) error {
	return validation.ValidateStruct(
		identifier,
		validation.Field(&identifier.Extension, keyRule...),
		validation.Field(&identifier.Bundle, keyRule...),
	)
}

// validateValueIdentifier is an internal helper for request validation.
func validateValueIdentifier(identifier *proto.Identifier) error {
	return validation.ValidateStruct(
		identifier,
		validation.Field(&identifier.Extension, keyRule...),
		validation.Field(&identifier.Bundle, keyRule...),
		validation.Field(&identifier.Setting, settingKeyRule...),
		validation.Field(&identifier.AccountUuid, accountUUIDRule...),
	)
}

// validateRoleAssignmentIdentifier is an internal helper for request validation.
func validateRoleAssignmentIdentifier(assignment *proto.RoleAssignmentIdentifier) error {
	return validation.ValidateStruct(
		assignment,
		validation.Field(&assignment.AccountUuid, is.UUID),
		validation.Field(&assignment.Role, validation.Match(regexForKeys)),
		// TODO: if resource is present (it's not required), how do we validate the resource types and ids?
	)
}
