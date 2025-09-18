package intermediate

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}

	defer logger.Sync() // flushes buffer, if any

	logger.Info("This is an info message")

	logger.Info("User logged in", zap.String("username", "johndoe"), zap.String("method", "GET"), zap.Int("userID", 42))

}
