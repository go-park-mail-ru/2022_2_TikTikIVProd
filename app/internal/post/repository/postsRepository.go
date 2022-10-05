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

func (dbPosts *dataBasePosts) SelectAllPosts() (*[]postsModel.Post, error) {
	var posts []postsModel.Post
	dbPosts.db.Table("user_posts").Find(&posts) //TODO оттрекать ошибки
	return &posts, nil
}
