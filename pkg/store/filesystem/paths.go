package store

import (
	"os"
	"path"
	"strings"

	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

const virtualSystemFolder = "system"
const folderNameBundles = "bundles"
const folderNameValues = "values"
const folderNameRoleAssignments = "role-assignments"

// Builds the folder path for storing settings bundles. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFolderPathBundles(resource *proto.Resource, mkdir bool) string {
	folderPath := path.Join(s.dataPath, folderNameBundles)
	if resource != nil {
		folderPath = path.Join(folderPath, strings.ToLower(proto.ResourceType_name[int32(resource.Type)]) + "-" + resource.Id)
	} else {
		folderPath = path.Join(folderPath, virtualSystemFolder)
	}
	if mkdir {
		s.ensureFolderExists(folderPath)
	}
	return folderPath
}

// Builds a unique file name from the given params. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathForBundle(identifier *proto.Identifier, resource *proto.Resource, mkdir bool) string {
	extensionFolder := path.Join(s.buildFolderPathBundles(resource, mkdir), identifier.Extension)
	if mkdir {
		s.ensureFolderExists(extensionFolder)
	}
	return path.Join(extensionFolder, identifier.Bundle+".json")
}

// Builds a unique file name from the given settings value. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathFromValue(value *proto.SettingsValue, mkdir bool) string {
	return s.buildFilePathFromValueArgs(value.Identifier.AccountUuid, value.Identifier.Extension, value.Identifier.Bundle, mkdir)
}

// Builds a unique file name from the given params. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathFromValueArgs(accountUUID string, extension string, bundleKey string, mkdir bool) string {
	extensionFolder := path.Join(s.dataPath, folderNameValues, accountUUID, extension)
	if mkdir {
		s.ensureFolderExists(extensionFolder)
	}
	return path.Join(extensionFolder, bundleKey+".json")
}

// Builds a unique file name from the given params. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathFromRoleAssignmentArgs(accountUUID string, mkdir bool) string {
	roleAssignmentsFolder := path.Join(s.dataPath, folderNameRoleAssignments)
	if mkdir {
		s.ensureFolderExists(roleAssignmentsFolder)
	}
	return path.Join(roleAssignmentsFolder, accountUUID+".json")
}

// Checks if the given path is an existing folder and creates one if not existing
func (s Store) ensureFolderExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0700)
		if err != nil {
			s.Logger.Err(err).Msgf("Error creating folder %v", path)
		}
	}
}
