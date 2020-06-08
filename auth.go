package convai_package_sdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signatureVerificationMiddleware(key string) func(c *gin.Context) {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		sig := getSignature(body, key)

		header := c.GetHeader("X-Convai-Signature")

		if header != sig {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			return
		}

		c.Next()
	}
}

func getSignature(body []byte, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write(body)
	hash := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash)
}