package store

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

const virtualSystemFolder = "system"
const folderNameBundles = "bundles"
const folderNameValues = "values"
const folderNameRoleAssignments = "role-assignments"

// buildFolderPathForBundles builds the folder path for storing settings bundles. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFolderPathForBundles(resource *proto.Resource, mkdir bool) string {
	folderPath := filepath.Join(s.dataPath, folderNameBundles, getResourceFolderName(resource))
	if mkdir {
		s.ensureFolderExists(folderPath)
	}
	return folderPath
}

// buildFilePathForBundle builds a unique file name from the given params. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathForBundle(identifier *proto.Identifier, resource *proto.Resource, mkdir bool) string {
	extensionFolder := filepath.Join(s.buildFolderPathForBundles(resource, mkdir), identifier.Extension)
	if mkdir {
		s.ensureFolderExists(extensionFolder)
	}
	return filepath.Join(extensionFolder, identifier.Bundle+".json")
}

// buildFolderPathForValues builds a unique folder path for storing settings values. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFolderPathForValues(identifier *proto.Identifier, resource *proto.Resource, mkdir bool) string {
	folderPath := filepath.Join(s.dataPath, folderNameValues, identifier.AccountUuid, getResourceFolderName(resource))
	if mkdir {
		s.ensureFolderExists(folderPath)
	}
	return folderPath
}

// buildFilePathForValue builds a unique file name from the given params. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathForValue(identifier *proto.Identifier, resource *proto.Resource, mkdir bool) string {
	extensionFolder := filepath.Join(s.buildFolderPathForValues(identifier, resource, mkdir), identifier.Extension)
	if mkdir {
		s.ensureFolderExists(extensionFolder)
	}
	return filepath.Join(extensionFolder, identifier.Bundle+".json")
}

// buildFilePathFromRoleAssignmentArgs builds a unique file name from the given params. If mkdir is true, folders in the path will be created if necessary.
func (s Store) buildFilePathFromRoleAssignmentArgs(accountUUID string, mkdir bool) string {
	roleAssignmentsFolder := filepath.Join(s.dataPath, folderNameRoleAssignments)
	if mkdir {
		s.ensureFolderExists(roleAssignmentsFolder)
	}
	return filepath.Join(roleAssignmentsFolder, accountUUID+".json")
}

// ensureFolderExists checks if the given path is an existing folder and creates one if not existing
func (s Store) ensureFolderExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0700)
		if err != nil {
			s.Logger.Err(err).Msgf("Error creating folder %v", path)
		}
	}
}

// getResourceFolderName returns a proper resource folder name for the provided resource
func getResourceFolderName(resource *proto.Resource) string {
	if resource != nil {
		return filepath.Join(strings.ToLower(proto.ResourceType_name[int32(resource.Type)]) + "-" + resource.Id)
	}
	return virtualSystemFolder
}
