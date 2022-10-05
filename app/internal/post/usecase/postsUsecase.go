package postsRep

import (
	imgUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	"time"
)

type PostsRepository interface {
	SelectPost(id int) (*Post, error)
	SelectAllPosts() (*[]Post, error)
	CreatePost(u Post) (*Post, error)
}

type PostsUsecase struct {
	postsRep PostsRepository
	imageRep imgUsecase.ImageReposiroty
}

type Post struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Message    string    `json:"message"`
	CreateDate time.Time `json:"create_date"`
	ImageLinks []string  `json:"image_links" gorm:"-"`
}

func NewPostsUsecase(ps PostsRepository, ir imgUsecase.ImageReposiroty) *PostsUsecase {
	return &PostsUsecase{
		postsRep: ps,
		imageRep: ir,
	}
}

func (pr *PostsUsecase) SelectPost(id int) (*Post, error) {
	return &Post{}, nil
}

func (pr *PostsUsecase) CreatePost(u *Post) (*Post, error) {
	return &Post{}, nil
}
func (pr *PostsUsecase) SelectAllPosts() (*[]Post, error) {
	res, err := pr.postsRep.SelectAllPosts() //TODO ошибки
	return res, err
}
