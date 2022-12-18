package postgres

import (
	fileRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/file/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type File struct {
	ID       uint64
	FileLink string
}

func (File) TableName() string {
	return "files"
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) fileRep.RepositoryI {
	return &fileRepository{
		db: db,
	}
}

func toPostgresFile(im *models.File) *File {
	return &File{
		ID:       im.ID,
		FileLink: im.FileLink,
	}
}

func toModelFile(f *File) *models.File {
	return &models.File{
		ID:       f.ID,
		FileLink: f.FileLink,
	}
}

func toModelFiles(files []*File) []*models.File {
	out := make([]*models.File, len(files))

	for i, b := range files {
		out[i] = toModelFile(b)
	}

	return out
}

func (dbFile *fileRepository) GetPostFiles(postID uint64) ([]*models.File, error) {
	var files []*File
	tx := dbFile.db.Model(File{}).Joins("JOIN user_posts_files upf ON upf.file_id = files.id AND upf.user_post_id = ?", postID).Scan(&files)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toModelFiles(files), nil
}

func (dbFile *fileRepository) GetFileById(fileID uint64) (*models.File, error) {
	var file File
	tx := dbFile.db.Table("files").First(&file, fileID)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "Get file repository error")
	}

	return toModelFile(&file), nil
}

func (dbFile *fileRepository) CreateFile(file *models.File) error {
	pfile := toPostgresFile(file)

	tx := dbFile.db.Create(pfile)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "fileRepository.CreateFile error while insert file")
	}

	file.ID = pfile.ID

	return nil
}
