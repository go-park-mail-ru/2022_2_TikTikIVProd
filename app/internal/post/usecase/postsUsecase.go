package usecase

import (
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/pkg/errors"
)

type PostUseCaseI interface {
	GetPostById(id int) (*models.Post, error)
	GetUserPosts(userId int) ([]*models.Post, error)
	CreatePost(p *models.Post) error
	UpdatePost(p *models.Post) error
	GetAllPosts() ([]*models.Post, error)
	DeletePost(id int, userId int) error
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
	resPost, err := p.postsRepo.GetPostById(id)

	if err != nil {
		return nil, errors.Wrap(err, "postsUsecase.GetPostById error while get post info")
	}

	err = addAdditionalFieldsToPost(resPost, p.userRepo, p.imageRepo)

	if err != nil {
		return nil, errors.Wrap(err, "postsUsecase.GetPostById error while get additional info")
	}

	return resPost, nil
}

func (p *postsUsecase) DeletePost(id int, userId int) error {
	existedPost, err := p.postsRepo.GetPostById(id)
	if err != nil {
		return err
	}

	if existedPost == nil {
		return errors.New("Post not found")
	}

	if existedPost.UserID != userId {
		return errors.New("Permission denied")
	}

	err = p.postsRepo.DeletePostById(id)

	if err != nil {
		return errors.Wrap(err, "postsUsecase.DeletePost error")
	}

	return nil
}

func (p *postsUsecase) CreatePost(post *models.Post) error {
	err := p.postsRepo.CreatePost(post)

	if err != nil {
		return errors.Wrap(err, "Error in func postsUsecase.CreatePost")
	}

	user, err := p.userRepo.SelectUserById(post.UserID)

	if err != nil {
		return errors.Wrap(err, "Error in func postsUsecase.CreatePost")
	}

	post.UserFirstName = user.FirstName
	post.UserLastName = user.LastName
	post.AvatarID = user.Avatar

	return nil
}

func (p *postsUsecase) UpdatePost(post *models.Post) error {
	existedPost, err := p.postsRepo.GetPostById(post.ID)
	if err != nil {
		return err
	}

	if existedPost == nil {
		return errors.New("Post not found")
	}

	if existedPost.UserID != post.UserID {
		return errors.New("Permission denied")
	}

	err = p.postsRepo.UpdatePost(post)

	if err != nil {
		return errors.Wrap(err, "Error in func postsUsecase.UpdatePost")
	}

	err = addAuthorForPost(post, p.userRepo)

	if err != nil {
		return errors.Wrap(err, "Error in func postsUsecase.UpdatePost")
	}

	return nil
}

func addAuthorForPost(post *models.Post, repUsers userRep.RepositoryI) error {
	author, err := repUsers.SelectUserById(post.UserID)

	if err != nil {
		return errors.Wrap(err, "Error in func addAuthorForPost")
	}

	post.UserLastName = author.LastName
	post.UserFirstName = author.FirstName
	post.AvatarID = author.Avatar

	return nil
}

func addImagesForPost(post *models.Post, repImg imageRep.RepositoryI) error {
	images, err := repImg.GetPostImages(post.ID)

	if err != nil {
		return errors.Wrap(err, "Error in func addPostImagesAuthors")
	}

	post.Images = make([]models.Image, 0, 10)

	for _, image := range images {
		post.Images = append(post.Images, *image)
	}

	return nil
}

func addAdditionalFieldsToPost(post *models.Post, repUsers userRep.RepositoryI, repImg imageRep.RepositoryI) error {
	err := addImagesForPost(post, repImg)

	if err != nil {
		return errors.Wrap(err, "error while get images")
	}

	err = addAuthorForPost(post, repUsers)

	if err != nil {
		return errors.Wrap(err, "error while get users")
	}

	return nil
}

func (p *postsUsecase) GetAllPosts() ([]*models.Post, error) {
	posts, err := p.postsRepo.GetAllPosts()

	if err != nil {
		return nil, errors.Wrap(err, "Error in func postsUsecase.GetAllPosts")
	}

	for idx := range posts {
		err = addAdditionalFieldsToPost(posts[idx], p.userRepo, p.imageRepo)

		if err != nil {
			return nil, errors.Wrap(err, "postsUsecase.GetAllPosts error while add additional fields")
		}
	}

	return posts, nil
}

func (p *postsUsecase) GetUserPosts(userId int) ([]*models.Post, error) {
	posts, err := p.postsRepo.GetUserPosts(userId)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func postsUsecase.GetUserPosts")
	}

	for idx := range posts {
		err = addAdditionalFieldsToPost(posts[idx], p.userRepo, p.imageRepo)

		if err != nil {
			return nil, errors.Wrap(err, "postsUsecase.GetUserPosts error while add additional fields")
		}
	}

	return posts, nil
}
