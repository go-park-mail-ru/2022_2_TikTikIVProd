package postgres_test

import (
	"regexp"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/faker"
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// func TestRepositoryGetPostImagcces(t *testing.T) {
// 	db, mock, err := sqlmock.New()
//     if err != nil {
// 		t.Fatal(err)
// 	}
//     defer db.Close()

// 	dialector := postgres.New(postgres.Config{
//         DSN:                  "sqlmock_db_0",
//         DriverName:           "postgres",
//         Conn:                 db,
//         PreferSimpleProtocol: true,
//     })
//     gdb, err := gorm.Open(dialector, &gorm.Config{})
//     if err != nil {
//         t.Fatal(err)
//     }

// 	gdb.Logger.LogMode(logger.Info)

//     mock.ExpectBegin()

// 	var mockUser models.User
// 	err = faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

//     mockUser.Id = 1

//     var userId uint64 = 1

//     mock.ExpectQuery(regexp.QuoteMeta(
//     `INSERT INTO "users" ("first_name","last_name","nick_name","email",`+
//     `"password","created_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).WithArgs(mockUser.FirstName,
//         mockUser.LastName, mockUser.NickName, mockUser.Email, mockUser.Password, mockUser.CreatedAt, mockUser.Id).
//     WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userId))

//     mock.ExpectCommit()

// 	repository := userRep.New(gdb)

//     err = repository.CreateUser(&mockUser)
//     require.NoError(t, err)
//     assert.Equal(t, userId, mockUser.Id)

//     err = mock.ExpectationsWereMet()
//     assert.NoError(t, err)
// }

func TestRepositoryGetPostImages(t *testing.T) {
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

	mockImages := make([]*models.Image, 0, 10)
	err = faker.FakeData(&mockImages)
	assert.NoError(t, err)

    var postId uint64 = 1

    rows := sqlmock.NewRows([]string{"id", "img_link"})

    for _, mockImage := range mockImages {
        rows.AddRow(mockImage.ID, mockImage.ImgLink)
    }
    
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT "images"."id","images"."img_link" FROM "images"`+
	` JOIN user_posts_images upi ON upi.img_id = images.id AND upi.user_post_id = $1`)).
	WillReturnRows(rows)

	repository := imageRep.NewImageRepository(gdb)

    actualImgs, err := repository.GetPostImages(postId)
    require.NoError(t, err)
    assert.Equal(t, mockImages, actualImgs)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositoryGetImage(t *testing.T) {
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

	var mockImage models.Image
	err = faker.FakeData(&mockImage)
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "images" WHERE "images"."id" = $1 ORDER BY "images"."id" LIMIT 1`)).
		WithArgs(mockImage.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "img_link"}).
		AddRow(mockImage.ID, mockImage.ImgLink))
	

	repository := imageRep.NewImageRepository(gdb)

    img, err := repository.GetImage(mockImage.ID)
    require.NoError(t, err)
    assert.Equal(t, &mockImage, img)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositoryCreateImage(t *testing.T) {
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

	var mockImage models.Image
	err = faker.FakeData(&mockImage)
	assert.NoError(t, err)

	mock.ExpectBegin()

    mock.ExpectQuery(regexp.QuoteMeta(
    `INSERT INTO "images" ("img_link","id") VALUES ($1,$2) RETURNING "id"`)).WithArgs(mockImage.ImgLink, mockImage.ID).
    WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockImage.ID))
	
	mock.ExpectCommit()
	
	repository := imageRep.NewImageRepository(gdb)

    err = repository.CreateImage(&mockImage)
    require.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

