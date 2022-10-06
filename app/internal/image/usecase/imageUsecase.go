package imageUsecase

import imagesModel "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/model"

type ImageReposiroty interface {
	SelectImagesInPost(postID int) (*[]imagesModel.Image, error)
}
