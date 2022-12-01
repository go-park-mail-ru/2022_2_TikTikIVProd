package usecase

import (
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository"
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type PostUseCaseI interface {
	GetPostById(id uint64, userId uint64) (*models.Post, error)
	GetUserPosts(userId uint64, curUserId uint64) ([]*models.Post, error)
	GetCommunityPosts(id uint64, userId uint64) ([]*models.Post, error)
	CreatePost(p *models.Post) error
	UpdatePost(p *models.Post) error
	GetAllPosts(userId uint64) ([]*models.Post, error)
	DeletePost(id uint64, userId uint64) error
	LikePost(id uint64, userId uint64) error
	UnLikePost(id uint64, userId uint64) error
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

func (p *postsUsecase) GetPostById(id uint64, userId uint64) (*models.Post, error) {
	resPost, err := p.postsRepo.GetPostById(id)

	if err != nil {
		return nil, errors.Wrap(err, "postsUsecase.GetPostById error while get post info")
	}

	err = addAdditionalFieldsToPost(resPost, p, userId)

	if err != nil {
		return nil, errors.Wrap(err, "postsUsecase.GetPostById error while get additional info")
	}

	return resPost, nil
}

func (p *postsUsecase) DeletePost(id uint64, userId uint64) error {
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
		return errors.Wrap(err, "post repository error in delete")
	}

	return nil
}

func (p *postsUsecase) LikePost(id uint64, userId uint64) error {
	existedPost, err := p.postsRepo.GetPostById(id)
	if err != nil {
		return err
	}

	if existedPost == nil {
		return errors.New("Post not found")
	}

	err = p.postsRepo.LikePost(id, userId)

	if err != nil {
		return errors.Wrap(err, "post repository error in like")
	}

	return nil
}

func (p *postsUsecase) UnLikePost(id uint64, userId uint64) error {
	existedPost, err := p.postsRepo.GetPostById(id)
	if err != nil {
		return err
	}

	if existedPost == nil {
		return errors.New("Post not found")
	}

	err = p.postsRepo.UnLikePost(id, userId)

	if err != nil {
		return errors.Wrap(err, "post repository error in like")
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

func addCountLikesForPost(post *models.Post, postsRepo repository.RepositoryI) error {
	count, err := postsRepo.GetCountLikesPost(post.ID)

	if err != nil {
		return errors.Wrap(err, "Post repository error in func addPostImagesAuthors")
	}

	post.CountLikes = count

	return nil
}

func addIsLikedForPost(post *models.Post, postRepo repository.RepositoryI, userId uint64) error {
	isLiked, err := postRepo.CheckLikePost(post.ID, userId)
	if err != nil {
		return errors.Wrap(err, "postsUsecase.GetPostById error while check like")
	}

	post.IsLiked = isLiked
	return nil
}

func addAdditionalFieldsToPost(post *models.Post, postsUsecase *postsUsecase, userId uint64) error {
	err := addImagesForPost(post, postsUsecase.imageRepo)

	if err != nil {
		return errors.Wrap(err, "error while get images")
	}

	err = addAuthorForPost(post, postsUsecase.userRepo)

	if err != nil {
		return errors.Wrap(err, "error while get users")
	}

	err = addCountLikesForPost(post, postsUsecase.postsRepo)

	if err != nil {
		return errors.Wrap(err, "error while get likes")
	}

	err = addIsLikedForPost(post, postsUsecase.postsRepo, userId)

	if err != nil {
		return errors.Wrap(err, "error while get isLiked")
	}

	return nil
}

func (p *postsUsecase) GetAllPosts(userId uint64) ([]*models.Post, error) {
	posts, err := p.postsRepo.GetAllPosts()

	if err != nil {
		return nil, errors.Wrap(err, "Error in func postsUsecase.GetAllPosts")
	}

	for idx := range posts {
		err = addAdditionalFieldsToPost(posts[idx], p, userId)

		if err != nil {
			return nil, errors.Wrap(err, "postsUsecase.GetAllPosts error while add additional fields")
		}
	}

	return posts, nil
}

func (p *postsUsecase) GetUserPosts(userId uint64, curUserId uint64) ([]*models.Post, error) {
	posts, err := p.postsRepo.GetUserPosts(userId)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func postsUsecase.GetUserPosts")
	}

	for idx := range posts {
		err = addAdditionalFieldsToPost(posts[idx], p, curUserId)

		if err != nil {
			return nil, errors.Wrap(err, "postsUsecase.GetUserPosts error while add additional fields")
		}
	}

	return posts, nil
}

func (p *postsUsecase) GetCommunityPosts(communityID uint64, userId uint64) ([]*models.Post, error) {
	posts, err := p.postsRepo.GetCommunityPosts(communityID)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func postsUsecase.GetUserPosts")
	}

	for idx := range posts {
		err = addAdditionalFieldsToPost(posts[idx], p, userId)

		if err != nil {
			return nil, errors.Wrap(err, "postsUsecase.GetUserPosts error while add additional fields")
		}
	}

	return posts, nil
}
