package utils

import (
	"os"
	"path"
)

// ******** Exported ******** //
type FileRepo struct {
	RootDir string
}

func GetFileRepo(rootDir string) *FileRepo {
	// check if rootDir exists
	// if not, return nil

	if !isDirExist(rootDir) {
		return nil
	}

	return &FileRepo{RootDir: rootDir}
}

func (fileRepo *FileRepo) Save(filepath string, data []byte) error {
	// split by '/' and get file name

	relDir := path.Dir(filepath)

	err := fileRepo.makeDir(relDir)
	if err != nil {
		return err
	}

	fullPath := path.Join(fileRepo.RootDir, filepath)

	err = os.WriteFile(fullPath, data, 0755)
	if err != nil {
		return err
	}

	return nil
}

// ******** Internal ******** //
func isDirExist(fullPath string) bool {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func (fileRepo *FileRepo) makeDir(path string) error {
	fullPath := Str_Concate(fileRepo.RootDir, "/", path)
	if isDirExist(fullPath) {
		return nil
	}

	err := os.MkdirAll(fullPath, 0755)

	if err != nil {
		return err
	}

	return nil
}
