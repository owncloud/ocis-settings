// Package store implements the go-micro store interface
package store

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gofrs/uuid"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

// ListValues reads all values that match the given bundleId and accountUUID
func (s Store) ListValues(bundleID, accountUUID string) ([]*proto.SettingsValue, error) {
	var records []*proto.SettingsValue
	valuesFolder := s.buildFolderPathForValues(false)
	valueFiles, err := ioutil.ReadDir(valuesFolder)
	if err != nil {
		return records, nil
	}

	for _, valueFile := range valueFiles {
		record := proto.SettingsValue{}
		err := s.parseRecordFromFile(&record, filepath.Join(valuesFolder, valueFile.Name()))
		if err != nil {
			s.Logger.Warn().Msgf("error reading %v", valueFile)
			continue
		}
		if record.BundleId != bundleID {
			continue
		}
		if record.AccountUuid != "" && record.AccountUuid != accountUUID {
			continue
		}
		records = append(records, &record)
	}

	return records, nil
}

// ReadValue tries to find a value by the given valueId within the dataPath
func (s Store) ReadValue(valueID string) (*proto.SettingsValue, error) {
	filePath := s.buildFilePathForValue(valueID, false)
	record := proto.SettingsValue{}
	if err := s.parseRecordFromFile(&record, filePath); err != nil {
		return nil, err
	}

	s.Logger.Debug().Msgf("read contents from file: %v", filePath)
	return &record, nil
}

// WriteValue writes the given SettingsValue into a file within the dataPath
func (s Store) WriteValue(value *proto.SettingsValue) (*proto.SettingsValue, error) {
	s.Logger.Debug().Str("value", value.String()).Msg("writing value")
	if value.Id == "" {
		value.Id = uuid.Must(uuid.NewV4()).String()
	}
	filePath := s.buildFilePathForValue(value.Id, true)
	if err := s.writeRecordToFile(value, filePath); err != nil {
		return nil, err
	}
	return value, nil
}
