package postgres_test

import (
	"regexp"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/faker"
	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestRepositoryAddFriend(t *testing.T) {
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

	var mockFriends models.Friends
	err = faker.FakeData(&mockFriends)
	assert.NoError(t, err)

    mock.ExpectExec(regexp.QuoteMeta(
    `INSERT INTO "friends" ("id1","id2") VALUES ($1,$2)`)).WithArgs(mockFriends.Id1,
        mockFriends.Id2).WillReturnResult(sqlmock.NewResult(int64(1), 1))

    mock.ExpectCommit()

	repository := friendsRep.New(gdb)

    err = repository.AddFriend(mockFriends)
    require.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositoryDeleteFriend(t *testing.T) {
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

	var mockFriends models.Friends
	err = faker.FakeData(&mockFriends)
	assert.NoError(t, err)

    mock.ExpectExec(regexp.QuoteMeta(
    `DELETE FROM "friends" WHERE id1 = $1 AND id2 = $2`)).WithArgs(mockFriends.Id1,
        mockFriends.Id2).WillReturnResult(sqlmock.NewResult(int64(1), 1))

    mock.ExpectCommit()

	repository := friendsRep.New(gdb)

    err = repository.DeleteFriend(mockFriends)
    require.NoError(t, err)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}
func TestRepositoryCheckFriends(t *testing.T) {
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

	var mockFriends models.Friends
	err = faker.FakeData(&mockFriends)
	assert.NoError(t, err)

    expectedCount := 1

    mock.ExpectQuery(regexp.QuoteMeta(
    `SELECT count(*) FROM "friends" WHERE id1 = $1 AND id2 = $2`)).WithArgs(mockFriends.Id1,
        mockFriends.Id2).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

	repository := friendsRep.New(gdb)

    friendSExist, err := repository.CheckFriends(mockFriends)
    require.NoError(t, err)
    assert.True(t, friendSExist)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

func TestRepositorySelectFriends(t *testing.T) {
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

	mockFriends := make([]models.User, 0, 10)
	err = faker.FakeData(&mockFriends)
	assert.NoError(t, err)

    userId := 1

    rows := sqlmock.NewRows([]string{"id",
    "first_name", "last_name", "nick_name", "avatar_img_id", "email"})

    for _, mockFriend := range mockFriends {
        rows.AddRow(mockFriend.Id, mockFriend.FirstName, mockFriend.LastName, mockFriend.NickName,
            mockFriend.Avatar, mockFriend.Email)
    }

    for i := range mockFriends {
        mockFriends[i].Password = ""
    }
    
    mock.ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id","users"."first_name",`+
    `"users"."last_name","users"."nick_name","users"."avatar_img_id","users"."email"`+
    ` FROM "users" JOIN friends ON friends.id2 = users.id WHERE id1 = $1`)).WillReturnRows(rows)

	repository := friendsRep.New(gdb)

    actualFriends, err := repository.SelectFriends(userId)
    require.NoError(t, err)
    assert.Equal(t, mockFriends, actualFriends)

    err = mock.ExpectationsWereMet()
    assert.NoError(t, err)
}

