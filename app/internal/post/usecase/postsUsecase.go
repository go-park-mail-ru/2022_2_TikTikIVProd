package postsRep

import (
	imageRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository/postgres"
	postsRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type PostUseCaseI interface {
	GetPostById(id int) (*models.Post, error)
	CreatePost(u *models.Post) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
}

type postsUsecase struct {
	postsRep postsRepository.RepositoryI
	imageRep imageRepository.RepositoryI
}

func NewPostUsecase(ps postsRepository.RepositoryI, ir imageRepository.RepositoryI) PostUseCaseI {
	return &postsUsecase{
		postsRep: ps,
		imageRep: ir,
	}
}

func (p *postsUsecase) GetPostById(id int) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postsUsecase) CreatePost(u *models.Post) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func addPostImages(posts []*models.Post, repImg imageRepository.RepositoryI) error {
	for idx := range posts {
		postImages, err := repImg.GetPostImages(posts[idx].ID)

		if err != nil {
			return err
		}

		for _, img := range postImages {
			posts[idx].Images = append(posts[idx].Images, *img)
		}
	}

	return nil
}

func (p *postsUsecase) GetAllPosts() ([]*models.Post, error) {
	posts, err := p.postsRep.GetAllPosts() //TODO ошибки

	if err != nil {
		return nil, err
	}

	err = addPostImages(posts, p.imageRep)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
