package postgres_test

import (
	"regexp"
	"testing"
	"time"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	postRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestRepositoryCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockPost := models.Post{
		ID:          1,
		UserID:      1,
		Message:     "message",
		CommunityID: 1,
	}
	mockPost.Attachments = make([]models.Attachment, 2)

	mockPost.Attachments[0] = models.Attachment{
		ID:      1,
		ImgLink: "link1",
	}

	mockPost.Attachments[1] = models.Attachment{
		ID:      2,
		ImgLink: "link2",
	}

	mockPost.CreateDate = time.Now()

	var postId uint64 = 1

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "user_posts" ("user_id","description","community_id","created_at","id")`+
			` VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).WithArgs(
		mockPost.UserID, mockPost.Message, mockPost.CommunityID, mockPost.CreateDate, mockPost.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(postId))

	mock.ExpectCommit()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "user_posts_attachments" ("user_post_id","att_id") VALUES ($1,$2),($3,$4)`)).
		WithArgs(mockPost.ID, mockPost.Attachments[0].ID, mockPost.ID, mockPost.Attachments[1].ID).
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	mock.ExpectCommit()

	repository := postRep.NewPostRepository(gdb)

	err = repository.CreatePost(&mockPost)
	require.NoError(t, err)
	assert.Equal(t, postId, mockPost.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositoryUpdatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockPost := models.Post{
		ID:          1,
		UserID:      1,
		Message:     "message",
		CommunityID: 1,
	}
	mockPost.Attachments = make([]models.Attachment, 2)

	mockPost.Attachments[0] = models.Attachment{
		ID:      1,
		ImgLink: "link1",
	}

	mockPost.Attachments[1] = models.Attachment{
		ID:      2,
		ImgLink: "link2",
	}

	mockPost.CreateDate = time.Now()

	var postId uint64 = 1

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "user_posts" SET "user_id"=$1,"description"=$2,"community_id"=$3,"created_at"=$4`+
			` WHERE "id" = $5`)).WithArgs(mockPost.UserID, mockPost.Message, mockPost.CommunityID,
		mockPost.CreateDate, mockPost.ID).WillReturnResult(sqlmock.NewResult(int64(mockPost.ID), 1))

	mock.ExpectCommit()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "user_posts_attachments" WHERE "user_posts_attachments"."user_post_id" = $1`)).
		WithArgs(mockPost.ID).
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	mock.ExpectCommit()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "user_posts_attachments" ("user_post_id","att_id") VALUES ($1,$2),($3,$4)`)).
		WithArgs(mockPost.ID, mockPost.Attachments[0].ID, mockPost.ID, mockPost.Attachments[1].ID).
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	mock.ExpectCommit()

	repository := postRep.NewPostRepository(gdb)

	err = repository.UpdatePost(&mockPost)
	require.NoError(t, err)
	assert.Equal(t, postId, mockPost.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositoryGetPostById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	gdb.Logger.LogMode(logger.Info)

	mockPost := models.Post{
		ID:          1,
		UserID:      1,
		Message:     "message",
		CommunityID: 1,
		CreateDate:  time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "user_posts" WHERE id = $1 LIMIT 1`)).WithArgs(mockPost.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "community_id",
			"created_at"}).AddRow(mockPost.ID, mockPost.UserID, mockPost.Message, mockPost.CommunityID,
			mockPost.CreateDate))

	repository := postRep.NewPostRepository(gdb)

	actualPost, err := repository.GetPostById(mockPost.ID)
	require.NoError(t, err)
	assert.Equal(t, mockPost, *actualPost)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
