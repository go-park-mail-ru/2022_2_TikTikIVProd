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

type PostsUsecase struct {
	postsRep postsRepository.RepositoryI
	imageRep imgUsecase.ImageReposiroty
}

func NewPostsUsecase(ps postsRepository.RepositoryI, ir imgUsecase.ImageReposiroty) UseCaseI {
	return &PostsUsecase{
		postsRep: ps,
		imageRep: ir,
	}
}

func (pr *PostsUsecase) SelectPost(id int) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}

func (pr *PostsUsecase) CreatePost(u *postsModel.Post) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}
func (pr *PostsUsecase) SelectAllPosts() (*[]postsModel.Post, error) {
	res, err := pr.postsRep.SelectAllPosts() //TODO ошибки
	return res, err
}
