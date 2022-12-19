package postgres

import (
	"time"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint64
	UserID      uint64
	Description string
	CommunityID uint64
	CreatedAt   time.Time
}

func (Post) TableName() string {
	return "user_posts"
}

type PostAttachmentsRelation struct {
	PostID       uint64 `gorm:"column:user_post_id"`
	AttachmentID uint64 `gorm:"column:att_id"`
}

func (PostAttachmentsRelation) TableName() string {
	return "user_posts_attachments"
}

func toPostgresPost(p *models.Post) *Post {
	return &Post{
		ID:          p.ID,
		UserID:      p.UserID,
		CommunityID: p.CommunityID,
		Description: p.Message,
		CreatedAt:   p.CreateDate,
	}
}

func toModelPost(p *Post) *models.Post {
	return &models.Post{
		ID:          p.ID,
		UserID:      p.UserID,
		CommunityID: p.CommunityID,
		Message:     p.Description,
		CreateDate:  p.CreatedAt,
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

	tx = dbPost.db.Where(&PostAttachmentsRelation{PostID: p.ID}).Delete(&PostAttachmentsRelation{})
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.UpdatePost error while delete relation")
	}

	postAttachments := make([]PostAttachmentsRelation, 0, 10)
	for _, elem := range p.Attachments {
		postAttachments = append(postAttachments, PostAttachmentsRelation{PostID: p.ID, AttachmentID: elem.ID})
	}

	if len(postAttachments) > 0 {
		tx = dbPost.db.Create(&postAttachments)
		if tx.Error != nil {
			return errors.Wrap(tx.Error, "postRepository.CreatePost error while insert relation")
		}
	}

	return nil
}

func (dbPost *postRepository) CreatePost(p *models.Post) error {
	post := toPostgresPost(p)

	var tx *gorm.DB = nil

	p.CreateDate = time.Now()

	if p.CommunityID == 0 {
		tx = dbPost.db.Omit("community_id").Create(post)
	} else {
		tx = dbPost.db.Create(post)
	}

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.CreatePost error while insert post")
	}

	p.ID = post.ID

	postAttachments := make([]PostAttachmentsRelation, 0, 10)
	for _, elem := range p.Attachments {
		postAttachments = append(postAttachments, PostAttachmentsRelation{PostID: p.ID, AttachmentID: elem.ID})
	}

	if len(postAttachments) > 0 {
		tx = dbPost.db.Create(&postAttachments)
		if tx.Error != nil {
			return errors.Wrap(tx.Error, "postRepository.CreatePost error while insert relation")
		}
	}

	return nil
}

func (dbPost *postRepository) GetPostById(id uint64) (*models.Post, error) {
	var post Post
	tx := dbPost.db.Where("id = ?", id).Take(&post)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetPostById error")
	}

	return toModelPost(&post), nil
}

func (dbPost *postRepository) DeletePostById(id uint64) error {
	tx := dbPost.db.Delete(&Post{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.DeletePostById error")
	}

	return nil
}

func (dbPost *postRepository) LikePost(id uint64, userId uint64) error {
	tx := dbPost.db.Create(&models.LikePost{UserID: userId, PostID: id})

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table like_post) on create")
	}

	return nil
}

func (dbPost *postRepository) UnLikePost(id uint64, userId uint64) error {
	tx := dbPost.db.Where(&models.LikePost{UserID: userId, PostID: id}).Delete(&models.LikePost{})

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "\"database error (table like_post) on delete")
	}

	return nil
}

func (dbPost *postRepository) GetCountLikesPost(id uint64) (uint64, error) {
	var count int64
	tx := dbPost.db.Model(&models.LikePost{}).Where("user_post_id = ?", id).Count(&count)

	if tx.Error != nil {
		return 0, errors.Wrap(tx.Error, "database error (table like_post) on count")
	}

	return uint64(count), nil
}

func (dbPost *postRepository) CheckLikePost(id uint64, userID uint64) (bool, error) {
	var count int64
	tx := dbPost.db.Model(&models.LikePost{}).Where(&models.LikePost{UserID: userID, PostID: id}).Count(&count)

	if tx.Error != nil {
		return false, errors.Wrap(tx.Error, "database error (table like_post) on check")
	}

	return count > 0, nil
}
func (dbPost *postRepository) GetAllPosts() ([]*models.Post, error) {
	posts := make([]*Post, 0, 10)
	tx := dbPost.db.Order("created_at desc").Find(&posts)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetAllPosts error")
	}

	return toModelPosts(posts), nil
}

func (dbPost *postRepository) GetUserPosts(userId uint64) ([]*models.Post, error) {
	posts := make([]*Post, 0, 10)
	tx := dbPost.db.Where("community_id is NULL").Where(&Post{UserID: userId}).Find(&posts)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetAllPosts error") // TODO
	}

	return toModelPosts(posts), nil
}

func (dbPost *postRepository) GetCommunityPosts(communityID uint64) ([]*models.Post, error) {
	posts := make([]*Post, 0, 10)
	tx := dbPost.db.Where(&Post{CommunityID: communityID}).Find(&posts)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetAllPosts error")
	}

	return toModelPosts(posts), nil
}

func (dbPost *postRepository) GetComments(postId uint64) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0, 10)
	tx := dbPost.db.Table("comments").Where(&models.Comment{PostID: postId}).Find(&comments)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetComments error")
	}

	return comments, nil
}

func (dbPost *postRepository) GetCommentById(id uint64) (*models.Comment, error) {
	var comment models.Comment
	tx := dbPost.db.Table("comments").Where("id = ?", id).Take(&comment)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "postRepository.GetPostGetCommentByIdById error")
	}

	return &comment, nil
}

func (dbPost *postRepository) AddComment(comment *models.Comment) error {
	tx := dbPost.db.Table("comments").Create(comment)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.AddComment error while insert comment")
	}

	return nil
}

func (dbPost *postRepository) UpdateComment(comment *models.Comment) error {
	tx := dbPost.db.Table("comments").Omit("id").Updates(comment)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.UpdateComment error while UPDATE comment")
	}

	return nil
}

func (dbPost *postRepository) DeleteComment(id uint64) error {
	tx := dbPost.db.Table("comments").Delete(&models.Comment{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "postRepository.DeleteComment error")
	}

	return nil
}
