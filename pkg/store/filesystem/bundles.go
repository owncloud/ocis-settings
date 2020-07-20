// Package store implements the go-micro store interface
package store

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

var m = &sync.RWMutex{}

// ListBundles returns all bundles in the dataPath folder belonging to the given extension
func (s Store) ListBundles(bundleType proto.SettingsBundle_Type) ([]*proto.SettingsBundle, error) {
	// FIXME: list requests should be ran against a cache, not FS
	m.RLock()
	defer m.RUnlock()

	var records []*proto.SettingsBundle
	bundlesFolder := s.buildFolderPathForBundles(false)
	bundleFiles, err := ioutil.ReadDir(bundlesFolder)
	if err != nil {
		return records, nil
	}

	for _, bundleFile := range bundleFiles {
		record := proto.SettingsBundle{}
		err = s.parseRecordFromFile(&record, filepath.Join(bundlesFolder, bundleFile.Name()))
		if err != nil {
			s.Logger.Warn().Msgf("error reading %v", bundleFile)
			continue
		}
		if record.Type != bundleType {
			continue
		}
		records = append(records, &record)
	}

	return records, nil
}

// ReadBundle tries to find a bundle by the given identifier within the dataPath.
// Extension and BundleKey within the identifier are required.
func (s Store) ReadBundle(bundleId string) (*proto.SettingsBundle, error) {
	// FIXME: locking should happen on the file here, not globally.
	m.RLock()
	defer m.RUnlock()

	filePath := s.buildFilePathForBundle(bundleId, false)
	record := proto.SettingsBundle{}
	if err := s.parseRecordFromFile(&record, filePath); err != nil {
		return nil, err
	}

	s.Logger.Debug().Msgf("read contents from file: %v", filePath)
	return &record, nil
}

// WriteBundle writes the given record into a file within the dataPath
// Extension and BundleKey within the record identifier are required.
func (s Store) WriteBundle(record *proto.SettingsBundle) (*proto.SettingsBundle, error) {
	// FIXME: locking should happen on the file here, not globally.
	m.Lock()
	defer m.Unlock()
	if record.Id == "" {
		record.Id = uuid.Must(uuid.NewV4()).String()
	}
	filePath := s.buildFilePathForBundle(record.Id, true)
	if err := s.writeRecordToFile(record, filePath); err != nil {
		return nil, err
	}

	s.Logger.Debug().Msgf("request contents written to file: %v", filePath)
	return record, nil
}

// AddSettingToBundle adds the given setting to the settings bundle which was identified by the given identifier and resource
func (s Store) AddSettingToBundle(bundleId string, setting *proto.Setting) (*proto.Setting, error) {
	bundle, err := s.ReadBundle(bundleId)
	if err != nil {
		return nil, err
	}
	if setting.Id == "" {
		setting.Id = uuid.Must(uuid.NewV4()).String()
	}
	setSetting(bundle, setting)
	_, err = s.WriteBundle(bundle)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

// RemoveSettingFromBundle removes the setting that was identified by the given identifier and resource
func (s Store) RemoveSettingFromBundle(bundleId string, settingId string) error {
	bundle, err := s.ReadBundle(bundleId)
	if err != nil {
		return nil
	}
	removeSetting(bundle, settingId)
	_, err = s.WriteBundle(bundle)
	return err
}

// indexOfSetting finds the index of the given setting within the given bundle.
// returns -1 if the setting was not found.
func indexOfSetting(bundle *proto.SettingsBundle, settingId string) int {
	for index := range bundle.Settings {
		s := bundle.Settings[index]
		if s.Id == settingId {
			return index
		}
	}
	return -1
}

// setSetting will append or overwrite the given setting within the given bundle
func setSetting(bundle *proto.SettingsBundle, setting *proto.Setting) {
	m.Lock()
	defer m.Unlock()
	index := indexOfSetting(bundle, setting.Id)
	if index == -1 {
		bundle.Settings = append(bundle.Settings, setting)
	} else {
		bundle.Settings[index] = setting
	}
}

// removeSetting will remove the given setting from the given bundle
func removeSetting(bundle *proto.SettingsBundle, settingId string) bool {
	m.Lock()
	defer m.Unlock()
	index := indexOfSetting(bundle, settingId)
	if index == -1 {
		return false
	}
	bundle.Settings = append(bundle.Settings[:index], bundle.Settings[index+1:]...)
	return true
}
