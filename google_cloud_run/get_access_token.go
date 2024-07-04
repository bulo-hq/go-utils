package google_cloud_run

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/idtoken"
)

func GetAccessToken(endpoint string) (*string, error) {
	idToken := os.Getenv("GOOGLE_CLOUD_IDENTITY_TOKEN")
	if idToken != "" {
		return &idToken, nil
	}

	tokenSource, err := idtoken.NewTokenSource(context.Background(), endpoint)
	if err != nil {
		return nil, fmt.Errorf("idtoken.NewTokenSource error: %w", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("tokenSource.Token error: %w", err)
	}

	return &token.AccessToken, nil
}
