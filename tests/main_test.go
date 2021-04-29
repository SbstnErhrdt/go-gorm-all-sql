package tests

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	ass := assert.New(t)
	client, err := sql.ConnectToDatabase()
	ass.NoError(err)
	ass.NotNil(client)
}
