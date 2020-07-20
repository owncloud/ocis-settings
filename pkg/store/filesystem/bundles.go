// Package store implements the go-micro store interface
package store

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

var m *sync.RWMutex = &sync.RWMutex{}

// ListBundles returns all bundles in the dataPath folder belonging to the given extension
func (s Store) ListBundles(identifier *proto.Identifier, resource *proto.Resource) ([]*proto.SettingsBundle, error) {
	m.RLock()
	var records []*proto.SettingsBundle
	bundlesFolder := s.buildFolderPathForBundles(resource, false)
	extensionFolders, err := ioutil.ReadDir(bundlesFolder)
	if err != nil {
		return records, nil
	}

	if len(identifier.Extension) < 1 {
		s.Logger.Info().Msg("listing all bundles")
	} else {
		s.Logger.Info().Msgf("listing bundles by extension %v", identifier.Extension)
	}
	for _, extensionFolder := range extensionFolders {
		extensionPath := filepath.Join(bundlesFolder, extensionFolder.Name())
		bundleFiles, err := ioutil.ReadDir(extensionPath)
		if err == nil {
			for _, bundleFile := range bundleFiles {
				record := proto.SettingsBundle{}
				bundlePath := filepath.Join(extensionPath, bundleFile.Name())
				err = s.parseRecordFromFile(&record, bundlePath)
				if err != nil {
					s.Logger.Warn().Msgf("error reading %v", bundlePath)
					continue
				}
				if len(identifier.Extension) == 0 || identifier.Extension == record.Identifier.Extension {
					records = append(records, &record)
				}
			}
		} else {
			s.Logger.Err(err).Msgf("error reading %v", extensionPath)
		}
	}

	m.RUnlock()

	return records, nil
}

// ReadBundle tries to find a bundle by the given identifier within the dataPath.
// Extension and BundleKey within the identifier are required.
func (s Store) ReadBundle(identifier *proto.Identifier, resource *proto.Resource) (*proto.SettingsBundle, error) {
	m.RLock()
	filePath := s.buildFilePathForBundle(identifier, resource, false)
	record := proto.SettingsBundle{}
	if err := s.parseRecordFromFile(&record, filePath); err != nil {
		return nil, err
	}
	m.RUnlock()

	s.Logger.Debug().Msgf("read contents from file: %v", filePath)
	return &record, nil
}

// WriteBundle writes the given record into a file within the dataPath
// Extension and BundleKey within the record identifier are required.
func (s Store) WriteBundle(record *proto.SettingsBundle) (*proto.SettingsBundle, error) {
	m.Lock()
	filePath := s.buildFilePathForBundle(record.Identifier, record.Resource, true)
	if err := s.writeRecordToFile(record, filePath); err != nil {
		return nil, err
	}

	m.Unlock()

	s.Logger.Debug().Msgf("request contents written to file: %v", filePath)
	return record, nil
}
