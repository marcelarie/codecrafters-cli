package custom_storage

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/filesystem/dotgit"
)

type ObjectStorage = filesystem.ObjectStorage
type IndexStorage = CustomIndexStorage
type ShallowStorage = CustomShallowStorage
type ConfigStorage = CustomConfigStorage
type ModuleStorage = CustomModuleStorage
type ReferenceStorage = CustomReferenceStorage

type CustomStorage struct {
	fs  billy.Filesystem
	dir *dotgit.DotGit

	ObjectStorage
	ReferenceStorage
	IndexStorage
	ShallowStorage
	ConfigStorage
	ModuleStorage
}

// NewCustomStorage returns a new CustomStorage that uses an in-memory Index storage
// backed by a given `fs.Filesystem` and cache.
func NewCustomStorage(fs billy.Filesystem, cache cache.Object) *CustomStorage {
	dir := dotgit.NewWithOptions(fs, dotgit.Options{})

	return &CustomStorage{
		fs:  fs,
		dir: dir,

		ObjectStorage:    *filesystem.NewObjectStorageWithOptions(dir, cache, filesystem.Options{}),
		ReferenceStorage: CustomReferenceStorage{dir: dir},
		IndexStorage:     CustomIndexStorage{dir: dir},
		ShallowStorage:   CustomShallowStorage{dir: dir},
		ConfigStorage:    CustomConfigStorage{dir: dir},
		ModuleStorage:    CustomModuleStorage{dir: dir},
	}
}

// Filesystem returns the underlying filesystem
func (s *CustomStorage) Filesystem() billy.Filesystem {
	return s.fs
}