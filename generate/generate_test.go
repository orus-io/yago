package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFkDef(t *testing.T) {
	for _, tt := range []struct {
		fkDef    string
		fk       string
		onUpdate string
		onDelete string
	}{
		{"Struct ONDELETE CASCADE ONUPDATE CASCADE",
			"Struct", "CASCADE", "CASCADE"},
		{"Struct ONDELETE SET NULL",
			"Struct", "", "SET NULL"},
	} {
		fk, onUpdate, onDelete := parseFkDef(tt.fkDef)
		assert.Equal(t, tt.fk, fk)
		assert.Equal(t, tt.onUpdate, onUpdate)
		assert.Equal(t, tt.onDelete, onDelete)
	}
}
