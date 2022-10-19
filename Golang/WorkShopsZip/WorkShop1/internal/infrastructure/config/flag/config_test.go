package flag_test

import (
	flag "corp/fif/inte/customers/internal/infrastructure/config/flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFlagConfig(t *testing.T) {
	f := flag.NewFlagConfig()
	assert.Equal(t, f.HTTPPort, "8080")
}

func TestEnvValue(t *testing.T) {
	envName := "PORT"
	defaultValue := "50001"

	resp1 := flag.EnvValue(envName, defaultValue)
	assert.Equal(t, resp1, defaultValue)

	envValue := "100"
	os.Setenv("PORT", envValue)
	resp2 := flag.EnvValue(envName, defaultValue)
	assert.Equal(t, resp2, envValue)
}
