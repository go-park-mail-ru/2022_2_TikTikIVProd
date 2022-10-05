package postsRepository

import (
	"database/sql"

	postsModel "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/model"
	"gorm.io/gorm"
)

type RepositoryI interface {
	SelectPost(id int) (*postsModel.Post, error)
	CreatePost(u postsModel.Post) (*postsModel.Post, error)
	SelectAllPosts() (*[]postsModel.Post, error)
}

type dataBasePosts struct {
	db *gorm.DB
}

func NewDataBasePosts(db *gorm.DB) RepositoryI {
	return &dataBasePosts{
		db: db,
	}
}

func (dbPosts *dataBasePosts) SelectPost(id int) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}

func (dbPosts *dataBasePosts) CreatePost(u postsModel.Post) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
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

func (dbPosts *dataBasePosts) SelectAllPosts() (*[]postsModel.Post, error) {
	var posts []postsModel.Post
	tx := dbPosts.db.Table("user_posts").Select("user_posts.id, user_posts.message", "user_posts.create_date", "users.first_name", "users.last_name").Joins("JOIN users ON users.id = user_posts.user_id ").Scan(&posts) //TODO оттрекать ошибки

	if tx.Error != nil {
		return nil, tx.Error
	}

	for i, _ := range posts {
		linkRows, err := dbPosts.db.Table("images").Select("img_link").Joins("JOIN user_posts_images upi ON upi.img_id = images.id AND upi.user_post_id = ?", posts[i].ID).Rows()

		if err != nil {
			return nil, err
		}
		links, err := getLinksFromRows(linkRows)

		if err != nil {
			return nil, err
		}

		posts[i].ImageLinks = links
	}

	return &posts, nil
}
