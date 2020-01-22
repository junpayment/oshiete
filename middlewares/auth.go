package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	slackSigningSecret = "SLACK_SIGNING_SECRETE"
	slackVersion       = "v0"
)

// see https://api.slack.com/docs/verifying-requests-from-slack
func Auth(c *gin.Context) {
	if gin.Mode() == gin.ReleaseMode {
		signingSecret := os.Getenv(slackSigningSecret)
		timestamp := c.Request.Header.Get("X-Slack-Request-Timestamp")
		body := c.Request.Body
		baseString, err := ioutil.ReadAll(body)
		// see https://stackoverflow.com/questions/33532374/in-go-how-can-i-reuse-a-readcloser
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(baseString))
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		hash := makeHMAC(strings.Join([]string{slackVersion, timestamp, string(baseString)}, ":"), signingSecret)
		if slackVersion+"="+hash != c.Request.Header.Get("X-Slack-Signature") {
			_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		}
	}
	c.Next()
}

// see https://cipepser.hatenablog.com/entry/2017/05/27/100516
func makeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
