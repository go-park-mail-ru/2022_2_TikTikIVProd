package fileUsecase

import (
	fileRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/file/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type FileUseCaseI interface {
	GetPostFiles(postID uint64) ([]*models.File, error)
	GetFileById(fileID uint64) (*models.File, error)
	CreateFile(img *models.File) error
}

type fileUsecase struct {
	fileRep fileRepository.RepositoryI
}

func NewFileUsecase(ir fileRepository.RepositoryI) FileUseCaseI {
	return &fileUsecase{
		fileRep: ir,
	}
}

func (i *fileUsecase) GetPostFiles(postID uint64) ([]*models.File, error) {
	files, err := i.fileRep.GetPostFiles(postID)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (i *fileUsecase) GetFileById(fileID uint64) (*models.File, error) {
	file, err := i.fileRep.GetFileById(fileID)

	if err != nil {
		return nil, errors.Wrap(err, "GetFile usecase error")
	}

	return file, nil
}

func (i *fileUsecase) CreateFile(img *models.File) error {
	err := i.fileRep.CreateFile(img)

	if err != nil {
		return errors.Wrap(err, "fileUsecase.CreateFile error")
	}

	return nil
}
