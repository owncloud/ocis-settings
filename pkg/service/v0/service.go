package svc

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-pkg/v2/middleware"
	"github.com/owncloud/ocis-settings/pkg/config"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/owncloud/ocis-settings/pkg/settings"
	store "github.com/owncloud/ocis-settings/pkg/store/filesystem"
)

// Service represents a service.
type Service struct {
	config  *config.Config
	logger  log.Logger
	manager settings.Manager
}

// NewService returns a service implementation for Service.
func NewService(cfg *config.Config, logger log.Logger) Service {
	service := Service{
		config:  cfg,
		logger:  logger,
		manager: store.New(cfg),
	}
	for _, role := range GenerateSettingsBundlesDefaultRoles() {
		_, err := service.manager.WriteBundle(role)
		bundleID := role.Identifier.Extension + "." + role.Identifier.Bundle
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to register settings bundle %v", bundleID)
		}
		logger.Debug().Msgf("Successfully registered settings bundle %v", bundleID)
	}
	return service
}

// SaveSettingsBundle implements the BundleServiceHandler interface
func (g Service) SaveSettingsBundle(c context.Context, req *proto.SaveSettingsBundleRequest, res *proto.SaveSettingsBundleResponse) error {
	req.SettingsBundle.Identifier = getFailsafeIdentifier(c, req.SettingsBundle.Identifier)
	if validationError := validateSaveSettingsBundle(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.WriteBundle(req.SettingsBundle)
	if err != nil {
		return err
	}
	res.SettingsBundle = r
	return nil
}

// GetSettingsBundle implements the BundleServiceHandler interface
func (g Service) GetSettingsBundle(c context.Context, req *proto.GetSettingsBundleRequest, res *proto.GetSettingsBundleResponse) error {
	req.Identifier = getFailsafeIdentifier(c, req.Identifier)
	if validationError := validateGetSettingsBundle(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ReadBundle(req.Identifier)
	if err != nil {
		return err
	}
	res.SettingsBundle = r
	return nil
}

// ListSettingsBundles implements the BundleServiceHandler interface
func (g Service) ListSettingsBundles(c context.Context, req *proto.ListSettingsBundlesRequest, res *proto.ListSettingsBundlesResponse) error {
	req.Identifier = getFailsafeIdentifier(c, req.Identifier)
	if validationError := validateListSettingsBundles(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ListBundles(req.Identifier)
	if err != nil {
		return err
	}
	res.SettingsBundles = r
	return nil
}

// SaveSettingsValue implements the ValueServiceHandler interface
func (g Service) SaveSettingsValue(c context.Context, req *proto.SaveSettingsValueRequest, res *proto.SaveSettingsValueResponse) error {
	req.SettingsValue.Identifier = getFailsafeIdentifier(c, req.SettingsValue.Identifier)
	if validationError := validateSaveSettingsValue(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.WriteValue(req.SettingsValue)
	if err != nil {
		return err
	}
	res.SettingsValue = r
	return nil
}

// GetSettingsValue implements the ValueServiceHandler interface
func (g Service) GetSettingsValue(c context.Context, req *proto.GetSettingsValueRequest, res *proto.GetSettingsValueResponse) error {
	req.Identifier = getFailsafeIdentifier(c, req.Identifier)
	if validationError := validateGetSettingsValue(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ReadValue(req.Identifier)
	if err != nil {
		return err
	}
	res.SettingsValue = r
	return nil
}

// ListSettingsValues implements the ValueServiceHandler interface
func (g Service) ListSettingsValues(c context.Context, req *proto.ListSettingsValuesRequest, res *proto.ListSettingsValuesResponse) error {
	req.Identifier = getFailsafeIdentifier(c, req.Identifier)
	if validationError := validateListSettingsValues(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ListValues(req.Identifier)
	if err != nil {
		return err
	}
	res.SettingsValues = r
	return nil
}

// ListRoleAssignments implements the RoleServiceHandler interface
func (g Service) ListRoleAssignments(c context.Context, req *proto.ListRoleAssignmentsRequest, res *proto.UserRoleAssignments) error {
	req.Assignment = getFailsafeRoleAssignment(c, req.Assignment)
	// TODO: validation. At least the accountUuid is required, rest is optional (but must comply if present).
	r, err := g.manager.ListRoleAssignments(req.Assignment)
	if err != nil {
		return err
	}
	res.Assignments = r.Assignments
	return nil
}

// AssignRoleToUser implements the RoleServiceHandler interface
func (g Service) AssignRoleToUser(c context.Context, req *proto.AssignRoleToUserRequest, _ *empty.Empty) error {
	req.Assignment = getFailsafeRoleAssignment(c, req.Assignment)
	// TODO: validation. If a role is not associated with a resource, it has to be set to SYSTEM.
	err := g.manager.WriteRoleAssignment(req.Assignment)
	if err != nil {
		return err
	}
	return nil
}

// RemoveRoleFromUser implements the RoleServiceHandler interface
func (g Service) RemoveRoleFromUser(c context.Context, req *proto.RemoveRoleFromUserRequest, _ *empty.Empty) error {
	req.Assignment = getFailsafeRoleAssignment(c, req.Assignment)
	// TODO: validation. If a role is not associated with a resource, it has to be set to SYSTEM.
	err := g.manager.DeleteRoleAssignment(req.Assignment)
	if err != nil {
		return err
	}
	return nil
}

// getFailsafeIdentifier makes sure that there is an identifier, and that the account uuid is injected if needed.
func getFailsafeIdentifier(c context.Context, identifier *proto.Identifier) *proto.Identifier {
	if identifier == nil {
		identifier = &proto.Identifier{}
	}
	identifier.AccountUuid = getValidatedAccountUUID(c, identifier.AccountUuid)
	return identifier
}

// getFailsafeRoleAssignment makes sure that there is an assignment identifier, and that the account uuid is injected if needed.
func getFailsafeRoleAssignment(c context.Context, assignment *proto.RoleAssignmentIdentifier) *proto.RoleAssignmentIdentifier {
	if assignment == nil {
		assignment = &proto.RoleAssignmentIdentifier{}
	}
	assignment.AccountUuid = getValidatedAccountUUID(c, assignment.AccountUuid)
	return assignment
}

// getValidatedAccountUUID converts `me` into an actual account uuid from the context, if possible.
// the result of this function will always be a valid lower-case UUID or an empty string.
func getValidatedAccountUUID(c context.Context, accountUUID string) string {
	if accountUUID == "me" {
		ownAccountUUID := c.Value(middleware.UUIDKey)
		if ownAccountUUID != nil && len(ownAccountUUID.(string)) > 0 {
			return strings.ToLower(ownAccountUUID.(string))
		}
		// might be valid for the request not having an AccountUUID in the context.
		// but clear it, instead of passing on the provided `me`.
		return ""
	}
	return strings.ToLower(accountUUID)
}
