package postsRepository

import (
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"gorm.io/gorm"
)

type DataBasePosts struct {
	db *gorm.DB
}

func NewDataBasePosts(db *gorm.DB) *DataBasePosts {
	return &DataBasePosts{
		db: db,
	}
}

func (dbPosts *DataBasePosts) SelectPost(id int) (*postsUsecase.Post, error) {
	return &postsUsecase.Post{}, nil
}

func (dbPosts *DataBasePosts) CreatePost(u postsUsecase.Post) (*postsUsecase.Post, error) {
	return &postsUsecase.Post{}, nil
}

func (dbPosts *DataBasePosts) SelectAllPosts() (*[]postsUsecase.Post, error) {
	var posts []postsUsecase.Post
	dbPosts.db.Take(&posts)
	return nil, nil
}
