package postsRep

import (
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type PostUseCaseI interface {
	GetPostById(id int) (*models.Post, error)
	CreatePost(u *models.Post) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
}

type postsUsecase struct {
	postsRepo repository.RepositoryI
	imageRepo imageRep.RepositoryI
	userRepo  userRep.RepositoryI
}

func NewPostUsecase(ps repository.RepositoryI, ir imageRep.RepositoryI, ur userRep.RepositoryI) PostUseCaseI {
	return &postsUsecase{
		postsRepo: ps,
		imageRepo: ir,
		userRepo:  ur,
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

func addPostImagesAuthors(posts []*models.Post, repImg imageRep.RepositoryI, repUsers userRep.RepositoryI) error {
	for idx := range posts {
		postImages, err := repImg.GetPostImages(posts[idx].ID)

		if err != nil {
			return err
		}

		for _, img := range postImages {
			posts[idx].Images = append(posts[idx].Images, *img)
		}

		author, err := repUsers.SelectUserById(posts[idx].ID)
		posts[idx].UserFirstName = author.FirstName
		posts[idx].UserLastName = author.LastName
	}

	return nil
}

func (p *postsUsecase) GetAllPosts() ([]*models.Post, error) {
	posts, err := p.postsRepo.GetAllPosts() //TODO ошибки

	if err != nil {
		return nil, err
	}

	err = addPostImagesAuthors(posts, p.imageRepo, p.userRepo)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
