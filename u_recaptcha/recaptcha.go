package recaptcha

import (
	"context"
	"encoding/json"
	"errors"
	http_client "github.com/tikivn/ultrago/u_http_client"
	logaff "github.com/tikivn/ultrago/u_logaff"
	"net/http"
	"time"
)

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func ValidateRecaptcha(ctx context.Context, secret string, captcha string) error {
	logger := logaff.GetNewLogger()

	if secret == "" || captcha == "" {
		return errors.New("missing captcha configuration")
	}

	params := map[string]string{
		"secret":   secret,
		"response": captcha,
	}

	httpClient := http_client.NewHttpClient(5*time.Second).WithUrl(siteVerifyURL, params)
	resp, err := httpClient.Do(ctx, http.MethodPost)
	if err != nil {
		return err
	}

	var body SiteVerifyResponse
	if err := json.Unmarshal(resp, &body); err != nil {
		return err
	}

	// Check recaptcha verification success.
	if !body.Success {
		return errors.New("unsuccessful recaptcha verify request")
	}

	// Check response score.
	if body.Score < 0.3 {
		logger.Errorf("lower received score than expected: %f", body.Score)
		return errors.New("lower received score than expected")
	}

	// Check response action.
	// if body.Action != "login" {
	// 	return errors.New("mismatched recaptcha action")
	// }

	return nil
}
