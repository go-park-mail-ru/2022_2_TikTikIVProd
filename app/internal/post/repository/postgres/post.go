package postgres

import (
	"database/sql"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         int
	UserID     int
	Message    string
	CreateDate time.Time
	FirstName  string
	LastName   string
}

func toPostgresPost(p *models.Post) *Post {
	return &Post{
		ID:         p.ID,
		UserID:     p.UserID,
		Message:    p.Message,
		CreateDate: p.CreateDate,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
	}
}

func toModelPost(p *Post) *models.Post {
	return &models.Post{
		ID:         p.ID,
		UserID:     p.UserID,
		Message:    p.Message,
		CreateDate: p.CreateDate,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
	}
}

func toModelPosts(posts []*Post) []*models.Post {
	out := make([]*models.Post, len(posts))

	for i, b := range posts {
		out[i] = toModelPost(b)
	}

	return out
}

type RepositoryI interface {
	GetPostById(id int) (*models.Post, error)
	CreatePost(u models.Post) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) RepositoryI {
	return &postRepository{
		db: db,
	}
}

func (dbPost *postRepository) CreatePost(u models.Post) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (dbPost *postRepository) GetPostById(id int) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func getLinksFromRows(rows *sql.Rows) ([]string, error) {
	var links []string
	defer rows.Close()
	for rows.Next() {
		var link string
		err := rows.Scan(&link)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func (dbPost *postRepository) GetAllPosts() ([]*models.Post, error) {
	var posts []*Post
	tx := dbPost.db.Table("user_posts").Select("user_posts.id, user_posts.message", "user_posts.create_date", "users.first_name", "users.last_name").Joins("JOIN users ON users.id = user_posts.user_id ").Scan(&posts) //TODO оттрекать ошибки

	if tx.Error != nil {
		log.Println()
		return nil, tx.Error
	}

	log.Println("Fetch all posts from postgres: ", posts)

	return toModelPosts(posts), nil
}

//func (dbPosts *postRepository) SelectAllPosts() (*[]Post, error) {
//	var posts []Post
//	tx := dbPosts.db.Table("user_posts").Select("user_posts.id, user_posts.message", "user_posts.create_date", "users.first_name", "users.last_name").Joins("JOIN users ON users.id = user_posts.user_id ").Scan(&posts) //TODO оттрекать ошибки
//
//	if tx.Error != nil {
//		return nil, tx.Error
//	}
//
//	for i := range posts {
//		linkRows, err := dbPosts.db.Table("images").Select("img_link").Joins("JOIN user_posts_images upi ON upi.img_id = images.id AND upi.user_post_id = ?", posts[i].ID).Rows()
//
//		if err != nil {
//			return nil, err
//		}
//		links, err := getLinksFromRows(linkRows)
//
//		if err != nil {
//			return nil, err
//		}
//
//		posts[i].ImageLinks = links
//	}
//
//	return &posts, nil
//}
