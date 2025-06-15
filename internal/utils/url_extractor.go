package utils

import (
	"net/url"
	"path"
)

// func ExtractS3KeyFromURL(url, baseURL string) string {
// 	return strings.TrimPrefix(url, baseURL+"/")
// }

func ExtractFileNameFromURL(fullURL string) string {
	parsed, err := url.Parse(fullURL)
	if err != nil {
		return ""
	}
	return path.Base(parsed.Path)
}
