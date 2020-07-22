// Package store implements the go-micro store interface
package store

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gofrs/uuid"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

// ListValues reads all values that match the given bundleId and accountUUID.
// If the bundleId is empty, it's ignored for filtering.
// If the accountUUID is empty, only values with empty accountUUID are returned.
// If the accountUUID is not empty, values with an empty or with a matching accountUUID are returned.
func (s Store) ListValues(bundleID, accountUUID string) ([]*proto.Value, error) {
	var records []*proto.Value
	valuesFolder := s.buildFolderPathForValues(false)
	valueFiles, err := ioutil.ReadDir(valuesFolder)
	if err != nil {
		return records, nil
	}

	for _, valueFile := range valueFiles {
		record := proto.Value{}
		err := s.parseRecordFromFile(&record, filepath.Join(valuesFolder, valueFile.Name()))
		if err != nil {
			s.Logger.Warn().Msgf("error reading %v", valueFile)
			continue
		}
		if bundleID != "" && record.BundleId != bundleID {
			continue
		}
		// if requested accountUUID empty -> fetch all system level values
		if accountUUID == "" && record.AccountUuid != "" {
			continue
		}
		// if requested accountUUID empty -> fetch all individual + all system level values
		if accountUUID != "" && record.AccountUuid != "" && record.AccountUuid != accountUUID {
			continue
		}
		records = append(records, &record)
	}

	return records, nil
}

// ReadValue tries to find a value by the given valueId within the dataPath
func (s Store) ReadValue(valueID string) (*proto.Value, error) {
	filePath := s.buildFilePathForValue(valueID, false)
	record := proto.Value{}
	if err := s.parseRecordFromFile(&record, filePath); err != nil {
		return nil, err
	}

	s.Logger.Debug().Msgf("read contents from file: %v", filePath)
	return &record, nil
}

// WriteValue writes the given value into a file within the dataPath
func (s Store) WriteValue(value *proto.Value) (*proto.Value, error) {
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
