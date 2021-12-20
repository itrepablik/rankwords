package envs

import "os"

const (
	SERVER_PORT = "8080"
)

// SetEnv sets the environment variables with string values
func SetEnv(env map[string]string) {
	for k, v := range env {
		os.Setenv(k, v)
	}
}

// InitEnvVars initialize environment variables
func InitEnvVars() error {
	// Set environment variables with string values
	SetEnv(map[string]string{
		"PORT": SERVER_PORT,
	})
	return nil
}
