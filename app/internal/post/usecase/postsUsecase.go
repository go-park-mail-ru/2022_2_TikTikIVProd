package postsRep

import (
	imgUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	postsRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	postsModel "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/model"
)

type UseCaseI interface {
	SelectPost(id int) (*postsModel.Post, error)
	CreatePost(u *postsModel.Post) (*postsModel.Post, error)
	SelectAllPosts() (*[]postsModel.Post, error)
}

type postsUsecase struct {
	postsRep postsRepository.RepositoryI
	imageRep imgUsecase.ImageReposiroty
}

func NewPostsUsecase(ps postsRepository.RepositoryI, ir imgUsecase.ImageReposiroty) UseCaseI {
	return &postsUsecase{
		postsRep: ps,
		imageRep: ir,
	}
}

func (pr *postsUsecase) SelectPost(id int) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}

func (pr *postsUsecase) CreatePost(u *postsModel.Post) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}
func (pr *postsUsecase) SelectAllPosts() (*[]postsModel.Post, error) {
	res, err := pr.postsRep.SelectAllPosts() //TODO ошибки
	return res, err
}
