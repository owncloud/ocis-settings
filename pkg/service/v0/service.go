package svc

import (
	"context"
	"strings"

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
	if _, err := service.manager.WriteBundle(generateSettingsBundleAdminRole()); err != nil {
		logger.Error().Err(err).Msg("Failed to register roles")
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

// getFailsafeIdentifier makes sure that there is an identifier, and that the account uuid is injected if needed.
func getFailsafeIdentifier(c context.Context, identifier *proto.Identifier) *proto.Identifier {
	if identifier == nil {
		identifier = &proto.Identifier{}
	}
	if identifier.AccountUuid == "me" {
		ownAccountUUID := c.Value(middleware.UUIDKey).(string)
		if len(ownAccountUUID) > 0 {
			identifier.AccountUuid = ownAccountUUID
		} else {
			// might be valid for the request not having an AccountUuid in the identifier.
			// but clear it, instead of passing on `me`.
			identifier.AccountUuid = ""
		}
	}
	identifier.AccountUuid = strings.ToLower(identifier.AccountUuid)
	return identifier
}
