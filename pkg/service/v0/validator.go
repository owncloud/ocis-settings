package svc

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

var (
	regexForKeys        = regexp.MustCompile(`^[A-Za-z0-9\-_]*$`)
	requireAlphanumeric = []validation.Rule{
		validation.Required,
		validation.Match(regexForKeys),
	}
)

func validateSaveSettingsBundle(req *proto.SaveSettingsBundleRequest) error {
	if err := validation.ValidateStruct(
		req.SettingsBundle,
		validation.Field(&req.SettingsBundle.Id, validation.When(req.SettingsBundle.Id != "", is.UUID)),
		validation.Field(&req.SettingsBundle.Type, validation.NotIn(proto.SettingsBundle_TYPE_UNKNOWN)),
		validation.Field(&req.SettingsBundle.Extension, requireAlphanumeric...),
		validation.Field(&req.SettingsBundle.DisplayName, validation.Required),
		validation.Field(&req.SettingsBundle.Settings, validation.Required),
	); err != nil {
		return err
	}
	if err := validateResource(req.SettingsBundle.Resource); err != nil {
		return err
	}
	for i := range req.SettingsBundle.Settings {
		if err := validateSetting(req.SettingsBundle.Settings[i]); err != nil {
			return err
		}
	}
	return nil
}

func validateGetSettingsBundle(req *proto.GetSettingsBundleRequest) error {
	return validation.Validate(&req.BundleId, is.UUID)
}

func validateListSettingsBundles(req *proto.ListSettingsBundlesRequest) error {
	return validation.Validate(&req.AccountUuid, is.UUID)
}

func validateAddSettingToSettingsBundle(req *proto.AddSettingToSettingsBundleRequest) error {
	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.BundleId, is.UUID),
	); err != nil {
		return err
	}
	return validateSetting(req.Setting)
}

func validateRemoveSettingFromSettingsBundle(req *proto.RemoveSettingFromSettingsBundleRequest) error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.BundleId, is.UUID),
		validation.Field(&req.SettingId, is.UUID),
	)
}

func validateSaveSettingsValue(req *proto.SaveSettingsValueRequest) error {
	if err := validation.ValidateStruct(
		req.SettingsValue,
		validation.Field(&req.SettingsValue.Id, validation.When(req.SettingsValue.Id != "", is.UUID)),
		validation.Field(&req.SettingsValue.SettingId, validation.When(req.SettingsValue.SettingId != "", is.UUID)),
		validation.Field(&req.SettingsValue.AccountUuid, is.UUID),
	); err != nil {
		return err
	}

	// TODO: validate values against the respective setting. need to check if constraints of the setting are fulfilled.
	return nil
}

func validateGetSettingsValue(req *proto.GetSettingsValueRequest) error {
	return validation.Validate(req.Id, is.UUID)
}

func validateListSettingsValues(req *proto.ListSettingsValuesRequest) error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.BundleId, validation.When(req.BundleId != "", is.UUID)),
		validation.Field(&req.AccountUuid, validation.When(req.AccountUuid != "", is.UUID)),
	)
}

func validateListRoles(req *proto.ListSettingsBundlesRequest) error {
	return validation.Validate(&req.AccountUuid, is.UUID)
}

func validateListRoleAssignments(req *proto.ListRoleAssignmentsRequest) error {
	return validation.Validate(req.AccountUuid, is.UUID)
}

func validateAssignRoleToUser(req *proto.AssignRoleToUserRequest) error {
	return validation.ValidateStruct(
		req.Assignment,
		validation.Field(&req.Assignment.Id, validation.Empty), // Assignment has an Id field but for assigning it must be empty.
		validation.Field(&req.Assignment.AccountUuid, is.UUID),
		validation.Field(&req.Assignment.RoleId, is.UUID),
	)
}

func validateRemoveRoleFromUser(req *proto.RemoveRoleFromUserRequest) error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.Id, is.UUID),
	)
}

// validateResource is an internal helper for validating the content of a resource.
func validateResource(resource *proto.Resource) error {
	if err := validation.Validate(&resource, validation.Required); err != nil {
		return err
	}
	return validation.Validate(&resource, validation.NotIn(proto.Resource_TYPE_UNKNOWN))
}

// validateSetting is an internal helper for validating the content of a setting.
func validateSetting(setting *proto.Setting) error {
	// TODO: make sanity checks, like for int settings, min <= default <= max.
	if err := validation.ValidateStruct(
		setting,
		validation.Field(&setting.Id, validation.When(setting.Id != "", is.UUID)),
	); err != nil {
		return err
	}
	return validateResource(setting.Resource)
}
