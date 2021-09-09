package u_recaptcha

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/tikivn/ultrago/u_env"
	"github.com/tikivn/ultrago/u_http_client"
	"github.com/tikivn/ultrago/u_logger"
)

const (
	RECAPTCHA_SECRET_KEY string = "RECAPTCHA_SECRET_KEY"
	RECAPTCHA_THRESHOLD  string = "RECAPTCHA_THRESHOLD"
)

func NewRecaptcha(httpExecutor u_http_client.HttpExecutor) (*Recaptcha, error) {
	secret := u_env.GetString(RECAPTCHA_SECRET_KEY, "")
	if secret == "" {
		return nil, errors.New("missing env RECAPTCHA_SECRET")
	}
	return &Recaptcha{
		verifyURL:    "https://www.google.com/recaptcha/api/siteverify",
		secret:       secret,
		threshold:    u_env.GetFloat(RECAPTCHA_THRESHOLD, 0.3),
		httpExecutor: httpExecutor,
	}, nil
}

type Recaptcha struct {
	verifyURL string
	secret    string
	threshold float64

	httpExecutor u_http_client.HttpExecutor
}

func (c Recaptcha) VerifyCaptcha(ctx context.Context, captcha string) error {
	ctx, logger := u_logger.GetLogger(ctx)
	if c.secret == "" {
		return errors.New("missing recaptcha secret")
	}
	if captcha == "" {
		return errors.New("missing captcha input")
	}

	params := map[string]string{
		"secret":   c.secret,
		"response": captcha,
	}

	resp, err := u_http_client.NewHttpClient(c.httpExecutor, 5*time.Second).
		WithUrl(c.verifyURL, params).
		Do(ctx, http.MethodPost)
	if err != nil {
		return err
	}

	var res SiteVerifyResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		logger.Errorf("parse SiteVerifyResponse failed: %v", err)
		return err
	} else if err = res.Validate(c.threshold); err != nil {
		logger.Errorf("validate SiteVerifyResponse failed: %v", err)
		return err
	}
	return nil
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func (res SiteVerifyResponse) Validate(threshold float64) error {
	if !res.Success {
		return fmt.Errorf("verify recaptcha failed with scores=%f", res.Score)
	}
	if res.Score < threshold {
		return fmt.Errorf("verify recaptcha with score=%f lower than threshold=%f", res.Score, threshold)
	}
	return nil
}
