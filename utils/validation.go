package utils

import "net/url"

// IsValidURL verifica si una URL tiene un formato válido
func IsValidURL(testURL string) bool {
    _, err := url.ParseRequestURI(testURL)
    return err == nil
}
