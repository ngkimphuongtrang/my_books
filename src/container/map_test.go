package container

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestMap_JSON(t *testing.T) {
	t.Parallel()

	body := Map{"book_id": 1}
	bodyJSON, err := body.JSON()
	require.NoError(t, err)
	log.Println(bodyJSON)
}
