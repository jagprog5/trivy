package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aquasecurity/trivy/pkg/iac/types"
)

func Test_ValueAsInt(t *testing.T) {
	val := NewValue(int64(10), types.NewTestMetadata())
	assert.Equal(t, 10, val.AsInt())
}
