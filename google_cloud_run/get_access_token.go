package google_cloud_run

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/idtoken"
)

func GetAuthorizationHeader(endpoint string) (*string, error) {
	accessToken := os.Getenv("GOOGLE_CLOUD_RUN_ACCESS_TOKEN")
	if accessToken != "" {
		return &accessToken, nil
	}
		tokenSource, err := idtoken.NewTokenSource(context.Background(), endpoint)
		if err != nil {
			return nil, fmt.Errorf("idtoken.NewTokenSource error: %s", err)
		}

		token, err := tokenSource.Token()
		if err != nil {
			log.Println("tokenSource.Token error:", err)
			return nil, fmt.Errorf("tokenSource.Token error: %s", err)
		}

	return &token.AccessToken, nil
}
