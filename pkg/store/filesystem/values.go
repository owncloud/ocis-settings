// Package store implements the go-micro store interface
package store

import (
	"os"
	"path/filepath"

	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

// ReadValue tries to find a value by the given identifier attributes within the dataPath
// All identifier fields are required.
func (s Store) ReadValue(identifier *proto.Identifier, resource *proto.Resource) (*proto.SettingsValue, error) {
	filePath := s.buildFilePathFromValueArgs(identifier.AccountUuid, identifier.Extension, identifier.Bundle, false)
	values, err := s.readValuesMapFromFile(filePath)
	if err != nil {
		return nil, err
	}
	if value := values.Values[identifier.Setting]; value != nil {
		return value, nil
	}
	// TODO: we want to return sensible defaults here, when the value was not found
	return nil, gstatus.Error(codes.NotFound, "SettingsValue not set")
}

// WriteValue writes the given SettingsValue into a file within the dataPath
// All identifier fields within the value are required.
func (s Store) WriteValue(value *proto.SettingsValue) (*proto.SettingsValue, error) {
	s.Logger.Debug().Str("value", value.String()).Msg("writing value")
	filePath := s.buildFilePathFromValue(value, true)
	values, err := s.readValuesMapFromFile(filePath)
	if err != nil {
		return nil, err
	}
	values.Values[value.Identifier.Setting] = value
	if err := s.writeRecordToFile(values, filePath); err != nil {
		return nil, err
	}
	return value, nil
}

// ListValues reads all values within the scope of the given identifier
// AccountUuid is required.
func (s Store) ListValues(identifier *proto.Identifier, resource *proto.Resource) ([]*proto.SettingsValue, error) {
	accountFolderPath := filepath.Join(s.dataPath, folderNameValues, identifier.AccountUuid)
	var values []*proto.SettingsValue
	if _, err := os.Stat(accountFolderPath); err != nil {
		return values, nil
	}

	// depending on the set values in the identifier arg, collect all SettingValues files for the account
	var valueFilePaths []string
	if len(identifier.Extension) < 1 {
		if err := filepath.Walk(accountFolderPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			valueFilePaths = append(valueFilePaths, path)
			return nil
		}); err != nil {
			s.Logger.Err(err).Msgf("error reading %v", accountFolderPath)
			return values, nil
		}
	} else if len(identifier.Bundle) < 1 {
		extensionPath := filepath.Join(accountFolderPath, identifier.Extension)
		if err := filepath.Walk(extensionPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			valueFilePaths = append(valueFilePaths, path)
			return nil
		}); err != nil {
			s.Logger.Err(err).Msgf("error reading %v", extensionPath)
			return values, nil
		}
	} else {
		bundlePath := filepath.Join(accountFolderPath, identifier.Extension, identifier.Bundle+".json")
		valueFilePaths = append(valueFilePaths, bundlePath)
	}

	// parse the SettingValues from the collected files
	for _, filePath := range valueFilePaths {
		bundleValues, err := s.readValuesMapFromFile(filePath)
		if err != nil {
			s.Logger.Err(err).Msgf("error reading %v", filePath)
		} else {
			for _, value := range bundleValues.Values {
				values = append(values, value)
			}
		}
	}
	return values, nil
}

// readValuesMapFromFile reads SettingsValues as map from the given file or returns an
// empty map if the file doesn't exist.
func (s Store) readValuesMapFromFile(filePath string) (*proto.SettingsValues, error) {
	values := &proto.SettingsValues{}
	err := s.parseRecordFromFile(values, filePath)
	if err != nil {
		if os.IsNotExist(err) {
			values.Values = map[string]*proto.SettingsValue{}
		} else {
			return nil, err
		}
	}
	return values, nil
}
