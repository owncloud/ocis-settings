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
		bundleId := role.Extension + "." + role.Id
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to register settings bundle %v", bundleId)
		}
		logger.Debug().Msgf("Successfully registered settings bundle %v", bundleId)
	}
	return service
}

// TODO: check permissions on every request

// SaveSettingsBundle implements the BundleServiceHandler interface
func (g Service) SaveSettingsBundle(c context.Context, req *proto.SaveSettingsBundleRequest, res *proto.SaveSettingsBundleResponse) error {
	cleanUpResource(c, req.SettingsBundle.Resource)
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
	if validationError := validateGetSettingsBundle(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ReadBundle(req.BundleId)
	if err != nil {
		return err
	}
	res.SettingsBundle = r
	return nil
}

// ListSettingsBundles implements the BundleServiceHandler interface
func (g Service) ListSettingsBundles(c context.Context, req *proto.ListSettingsBundlesRequest, res *proto.ListSettingsBundlesResponse) error {
	req.AccountUuid = getValidatedAccountUUID(c, req.AccountUuid)
	if validationError := validateListSettingsBundles(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ListBundles(req.AccountUuid, proto.SettingsBundle_DEFAULT)
	if err != nil {
		return err
	}
	res.SettingsBundles = r
	return nil
}

// AddSettingToSettingsBundle implements the BundleServiceHandler interface
func (g Service) AddSettingToSettingsBundle(c context.Context, req *proto.AddSettingToSettingsBundleRequest, res *proto.AddSettingToSettingsBundleResponse) error {
	cleanUpResource(c, req.Setting.Resource)
	if validationError := validateAddSettingToSettingsBundle(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.AddSettingToBundle(req.BundleId, req.Setting)
	if err != nil {
		return err
	}
	res.Setting = r
	return nil
}

// RemoveSettingFromSettingsBundle implements the BundleServiceHandler interface
func (g Service) RemoveSettingFromSettingsBundle(c context.Context, req *proto.RemoveSettingFromSettingsBundleRequest, _ *empty.Empty) error {
	if validationError := validateRemoveSettingFromSettingsBundle(req); validationError != nil {
		return validationError
	}
	return g.manager.RemoveSettingFromBundle(req.BundleId, req.SettingId)
}

// SaveSettingsValue implements the ValueServiceHandler interface
func (g Service) SaveSettingsValue(c context.Context, req *proto.SaveSettingsValueRequest, res *proto.SaveSettingsValueResponse) error {
	cleanUpResource(c, req.SettingsValue.Resource)
	// TODO: we need to check, if the authenticated user has permission to write the value for the specified resource (e.g. global, file with id xy, ...)
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
	if validationError := validateGetSettingsValue(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ReadValue(req.Id, )
	if err != nil {
		return err
	}
	res.SettingsValue = r
	return nil
}

// ListSettingsValues implements the ValueServiceHandler interface
func (g Service) ListSettingsValues(c context.Context, req *proto.ListSettingsValuesRequest, res *proto.ListSettingsValuesResponse) error {
	req.AccountUuid = getValidatedAccountUUID(c, req.AccountUuid)
	if validationError := validateListSettingsValues(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ListValues(req.BundleId, req.AccountUuid)
	if err != nil {
		return err
	}
	res.SettingsValues = r
	return nil
}

// ListRoles implements the RoleServiceHandler interface
func (g Service) ListRoles(c context.Context, req *proto.ListSettingsBundlesRequest, res *proto.ListSettingsBundlesResponse) error {
	req.AccountUuid = getValidatedAccountUUID(c, req.AccountUuid)
	if validationError := validateListRoles(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ListBundles(req.AccountUuid, proto.SettingsBundle_ROLE)
	if err != nil {
		return err
	}
	res.SettingsBundles = r
	return nil
}

// ListRoleAssignments implements the RoleServiceHandler interface
func (g Service) ListRoleAssignments(c context.Context, req *proto.ListRoleAssignmentsRequest, res *proto.ListRoleAssignmentsResponse) error {
	req.AccountUuid = getValidatedAccountUUID(c, req.AccountUuid)
	if validationError := validateListRoleAssignments(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.ListRoleAssignments(req.AccountUuid)
	if err != nil {
		return err
	}
	res.Assignments = r
	return nil
}

// AssignRoleToUser implements the RoleServiceHandler interface
func (g Service) AssignRoleToUser(c context.Context, req *proto.AssignRoleToUserRequest, res *proto.AssignRoleToUserResponse) error {
	req.Assignment.AccountUuid = getValidatedAccountUUID(c, req.Assignment.AccountUuid)
	if validationError := validateAssignRoleToUser(req); validationError != nil {
		return validationError
	}
	r, err := g.manager.WriteRoleAssignment(req.Assignment)
	if err != nil {
		return err
	}
	res.Assignment = r
	return nil
}

// RemoveRoleFromUser implements the RoleServiceHandler interface
func (g Service) RemoveRoleFromUser(c context.Context, req *proto.RemoveRoleFromUserRequest, _ *empty.Empty) error {
	if validationError := validateRemoveRoleFromUser(req); validationError != nil {
		return validationError
	}
	return g.manager.RemoveRoleAssignment(req.Id)
}

// cleanUpResource makes sure that the account uuid of the authenticated user is injected if needed.
func cleanUpResource(c context.Context, resource *proto.Resource) {
	if resource != nil && resource.Type == proto.Resource_USER {
		resource.Id = getValidatedAccountUUID(c, resource.Id)
	}
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
