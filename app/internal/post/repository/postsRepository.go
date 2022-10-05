package postsRepository

import (
	"gorm.io/gorm"
	postsModel "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/model"
)

type RepositoryI interface {
	SelectPost(id int) (*postsModel.Post, error)
	CreatePost(u postsModel.Post) (*postsModel.Post, error)
	SelectAllPosts() (*[]postsModel.Post, error)
}

type DataBasePosts struct {
	db *gorm.DB
}

func NewDataBasePosts(db *gorm.DB) RepositoryI {
	return &DataBasePosts{
		db: db,
	}
}

func (dbPosts *DataBasePosts) SelectPost(id int) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}

func (dbPosts *DataBasePosts) CreatePost(u postsModel.Post) (*postsModel.Post, error) {
	return &postsModel.Post{}, nil
}

func (dbPosts *DataBasePosts) SelectAllPosts() (*[]postsModel.Post, error) {
	var posts []postsModel.Post
	dbPosts.db.Table("user_posts").Find(&posts) //TODO оттрекать ошибки
	return &posts, nil
}
