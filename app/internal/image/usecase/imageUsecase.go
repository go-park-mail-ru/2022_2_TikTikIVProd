package imageUsecase

type Image struct {
	ID      int `json:"id"`
	ImgLink int `json:"img_link"`
}

type ImageReposiroty interface {
	SelectImagesInPost(postID int) (*[]Image, error)
}
