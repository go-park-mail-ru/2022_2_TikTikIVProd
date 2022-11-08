package postgres_test

import (
	"regexp"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/faker"
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestRepositoryCreateUser(t *testing.T) {
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

    mock.ExpectBegin()

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

    userId := 1

    mock.ExpectQuery(regexp.QuoteMeta(
    `INSERT INTO "users" ("first_name","last_name","nick_name","email",`+
    `"password","id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).WithArgs(mockUser.FirstName,
        mockUser.LastName, mockUser.NickName, mockUser.Email, mockUser.Password,mockUser.Id).
    WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userId))

    mock.ExpectCommit()

	repository := userRep.New(gdb)

    err = repository.CreateUser(&mockUser)
    require.NoError(t, err)
    assert.Equal(t, userId, mockUser.Id)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositoryUpdateUser(t *testing.T) {
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

    mock.ExpectBegin()

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

    mock.ExpectExec(regexp.QuoteMeta(
    `UPDATE "users" SET "first_name"=$1,"last_name"=$2,"nick_name"=$3,"avatar_img_id"=$4,`+
    `"email"=$5,"password"=$6 WHERE "id" = $7`)).WithArgs(mockUser.FirstName,
        mockUser.LastName, mockUser.NickName, mockUser.Avatar, mockUser.Email, mockUser.Password,
        mockUser.Id).WillReturnResult(sqlmock.NewResult(int64(mockUser.Id), 1))

    mock.ExpectCommit()

	repository := userRep.New(gdb)

    err = repository.UpdateUser(mockUser)
    require.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}
func TestRepositorySelectUserById(t *testing.T) {
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

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

    mock.ExpectQuery(regexp.QuoteMeta(
    `SELECT * FROM "users" WHERE id = $1 LIMIT 1`)).WithArgs(mockUser.Id).
    WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name",
    "nick_name", "avatar_img_id", "email", "password"}).AddRow(mockUser.Id, mockUser.FirstName,
        mockUser.LastName, mockUser.NickName, mockUser.Avatar, mockUser.Email, mockUser.Password))

	repository := userRep.New(gdb)

    actualUser, err := repository.SelectUserById(mockUser.Id)
    require.NoError(t, err)
    assert.Equal(t, mockUser, *actualUser)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositorySelectUserByNickName(t *testing.T) {
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

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

    mock.ExpectQuery(regexp.QuoteMeta(
    `SELECT * FROM "users" WHERE nick_name = $1 LIMIT 1`)).WithArgs(mockUser.NickName).
    WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name",
    "nick_name", "avatar_img_id", "email", "password"}).AddRow(mockUser.Id, mockUser.FirstName,
        mockUser.LastName, mockUser.NickName, mockUser.Avatar, mockUser.Email, mockUser.Password))

	repository := userRep.New(gdb)

    actualUser, err := repository.SelectUserByNickName(mockUser.NickName)
    require.NoError(t, err)
    assert.Equal(t, mockUser, *actualUser)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositorySelectUserByEmail(t *testing.T) {
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

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

    mock.ExpectQuery(regexp.QuoteMeta(
    `SELECT * FROM "users" WHERE email = $1 LIMIT 1`)).WithArgs(mockUser.Email).
    WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name",
    "nick_name", "avatar_img_id", "email", "password"}).AddRow(mockUser.Id, mockUser.FirstName,
        mockUser.LastName, mockUser.NickName, mockUser.Avatar, mockUser.Email, mockUser.Password))

	repository := userRep.New(gdb)

    actualUser, err := repository.SelectUserByEmail(mockUser.Email)
    require.NoError(t, err)
    assert.Equal(t, mockUser, *actualUser)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositorySelectAllUsers(t *testing.T) {
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

	mockUsers := make([]models.User, 0, 10)
	err = faker.FakeData(&mockUsers)
	assert.NoError(t, err)

    rows := sqlmock.NewRows([]string{"id",
    "first_name", "last_name", "nick_name", "avatar_img_id", "email", "password"})

    for _, mockUser := range mockUsers {
        rows.AddRow(mockUser.Id, mockUser.FirstName, mockUser.LastName, mockUser.NickName,
            mockUser.Avatar, mockUser.Email, mockUser.Password)
    }
    
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(rows)

	repository := userRep.New(gdb)

    actualUsers, err := repository.SelectAllUsers()
    require.NoError(t, err)
    assert.Equal(t, mockUsers, actualUsers)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

