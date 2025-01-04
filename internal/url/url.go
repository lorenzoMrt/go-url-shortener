package url

import (
	"crypto/sha256"
	"encoding/base64"
)

func Shorten(originalUrl string) string {
	hash := sha256.New()
	hash.Write([]byte(originalUrl))
	shortUrl := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	shortUrl = shortUrl[:8]
	return shortUrl
}
