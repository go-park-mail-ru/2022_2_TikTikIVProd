package postgres_test

import (
	"regexp"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/faker"
	attachmentRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/repository/postgres"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestRepositoryGetPostAttachments(t *testing.T) {
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

	mockAttachments := make([]*models.Attachment, 0, 10)
	err = faker.FakeData(&mockAttachments)
	assert.NoError(t, err)

	for idx := range mockAttachments {
		mockAttachments[idx].Type = "image"
	}

	var postId uint64 = 1

	rows := sqlmock.NewRows([]string{"id", "ttype", "att_link"})

	for _, mockAttachment := range mockAttachments {
		rows.AddRow(mockAttachment.ID, 0, mockAttachment.AttLink)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "attachments"."id","attachments"."att_link",` + 
	`"attachments"."ttype" FROM "attachments" JOIN user_posts_attachments upi ON ` +
	`upi.att_id = attachments.id AND upi.user_post_id = $1`)).WillReturnRows(rows)

	repository := attachmentRep.NewAttachmentRepository(gdb)

	actualImgs, err := repository.GetPostAttachments(postId)
	require.NoError(t, err)
	assert.Equal(t, mockAttachments, actualImgs)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositoryGetAttachment(t *testing.T) {
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

	var mockAttachment models.Attachment
	err = faker.FakeData(&mockAttachment)
	assert.NoError(t, err)
	mockAttachment.Type = "image"

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "attachments" WHERE "attachments"."id" = $1 ORDER BY "attachments"."id" LIMIT 1`)).
		WithArgs(mockAttachment.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "att_link"}).
		AddRow(mockAttachment.ID, mockAttachment.AttLink))

	repository := attachmentRep.NewAttachmentRepository(gdb)

	att, err := repository.GetAttachment(mockAttachment.ID)
	require.NoError(t, err)
	assert.Equal(t, &mockAttachment, att)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositoryCreateAttachment(t *testing.T) {
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

	var mockAttachment models.Attachment
	err = faker.FakeData(&mockAttachment)
	assert.NoError(t, err)

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "attachments" ("att_link","ttype","id") VALUES ($1,$2,$3) RETURNING "id"`)).WithArgs(mockAttachment.AttLink, 0, mockAttachment.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockAttachment.ID))

	mock.ExpectCommit()

	repository := attachmentRep.NewAttachmentRepository(gdb)

	err = repository.CreateAttachment(&mockAttachment)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepositoryGetMessageAttachments(t *testing.T) {
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

	mockAttachments := make([]*models.Attachment, 0, 10)
	err = faker.FakeData(&mockAttachments)
	assert.NoError(t, err)

	for idx := range mockAttachments {
		mockAttachments[idx].Type = "image"
	}

	var postId uint64 = 1

	rows := sqlmock.NewRows([]string{"id", "ttype", "att_link"})

	for _, mockAttachment := range mockAttachments {
		rows.AddRow(mockAttachment.ID, 0, mockAttachment.AttLink)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "attachments"."id","attachments"."att_link",` + 
	`"attachments"."ttype" FROM "attachments" JOIN message_attachments upi ON upi.att_id = ` +
	`attachments.id AND upi.message_id = $1`)).WillReturnRows(rows)

	repository := attachmentRep.NewAttachmentRepository(gdb)

	actualImgs, err := repository.GetMessageAttachments(postId)
	require.NoError(t, err)
	assert.Equal(t, mockAttachments, actualImgs)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
