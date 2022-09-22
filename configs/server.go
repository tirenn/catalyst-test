package configs

import (
	"fmt"
	"os"
)

// GetAddress will return formatted address
func GetAddress() string {
	host := os.Getenv("HOST")
	app_port := os.Getenv("PORT")

	return fmt.Sprintf("%s:%s", host, app_port)
}
