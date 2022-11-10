package postgres

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/pkg/errors"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         int
	UserID     int
	Message    string
	CreateDate time.Time
}

func (Post) TableName() string {
	return "user_posts"
}

type PostImagesRelation struct {
	PostID  int `gorm:"column:user_post_id"`
	ImageID int `gorm:"column:img_id"`
}

func (PostImagesRelation) TableName() string {
	return "user_posts_images"
}

func toPostgresPost(p *models.Post) *Post {
	return &Post{
		ID:         p.ID,
		UserID:     p.UserID,
		Message:    p.Message,
		CreateDate: p.CreateDate,
	}
}

func toModelPost(p *Post) *models.Post {
	return &models.Post{
		ID:         p.ID,
		UserID:     p.UserID,
		Message:    p.Message,
		CreateDate: p.CreateDate,
	}
}

func toModelPosts(posts []*Post) []*models.Post {
	out := make([]*models.Post, len(posts))

	for i, b := range posts {
		out[i] = toModelPost(b)
	}

	return out
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.RepositoryI {
	return &postRepository{
		db: db,
	}
}

func (dbPost *postRepository) UpdatePost(p *models.Post) error {
	post := toPostgresPost(p)

	tx := dbPost.db.Omit("id").Updates(post)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.UpdatePost error while insert post")
	}

	tx = dbPost.db.Where(&PostImagesRelation{PostID: p.ID}).Delete(&PostImagesRelation{})
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.UpdatePost error while delete relation")
	}

	postImages := make([]PostImagesRelation, 0, 10)
	for _, elem := range p.Images {
		postImages = append(postImages, PostImagesRelation{PostID: p.ID, ImageID: elem.ID})
	}

	if len(postImages) > 0 {
		tx = dbPost.db.Create(&postImages)
		if tx.Error != nil {
			return errors.Wrap(tx.Error, "postRepository.CreatePost error while insert relation")
		}
	}

	return nil
}

func (dbPost *postRepository) CreatePost(p *models.Post) error {
	post := toPostgresPost(p)

	tx := dbPost.db.Create(post)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.CreatePost error while insert post")
	}

	p.ID = post.ID
	p.CreateDate = time.Now()

	postImages := make([]PostImagesRelation, 0, 10)
	for _, elem := range p.Images {
		postImages = append(postImages, PostImagesRelation{PostID: p.ID, ImageID: elem.ID})
	}

	if len(postImages) > 0 {
		tx = dbPost.db.Create(&postImages)
		fmt.Println(postImages)
		if tx.Error != nil {
			return errors.Wrap(tx.Error, "postRepository.CreatePost error while insert relation")
		}
	}

	return nil
}

func (dbPost *postRepository) GetPostById(id int) (*models.Post, error) {
	var post Post
	tx := dbPost.db.First(&post, id)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetPostById error")
	}

	return toModelPost(&post), nil
}

func (dbPost *postRepository) DeletePostById(id int) error {
	tx := dbPost.db.Delete(&Post{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.DeletePostById error")
	}

	return nil
}

func (dbPost *postRepository) GetAllPosts() ([]*models.Post, error) {
	posts := make([]*Post, 0, 10)
	tx := dbPost.db.Table("user_posts").Find(&posts)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetAllPosts error")
	}

	return toModelPosts(posts), nil
}

func (dbPost *postRepository) GetUserPosts(userId int) ([]*models.Post, error) {
	posts := make([]*Post, 0, 10)
	tx := dbPost.db.Where(&Post{UserID: userId}).Find(&posts)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetAllPosts error")
	}

	return toModelPosts(posts), nil
}
