package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/huabtc/maroto/v2/internal/providers/gofpdf"

	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	// Act
	sut := gofpdf.NewLine(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.line", fmt.Sprintf("%T", sut))
}
