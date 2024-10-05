package shorturl

import (
	"strings"

	"github.com/samber/lo"
)

var (
	chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	div   = uint64(len(chars)) // 62
)

func ShortURLByID(id uint64) string {
	shortURL := make([]string, 0, 16)
	for {
		shortURL = append(shortURL, string(chars[id%div]))
		id = id / div

		if id == 0 {
			break
		}
	}

	shortURL = lo.Reverse(shortURL)
	return strings.Join(shortURL, "")
}

func IdByShortURL(shortURL string) uint64 {
	id := uint64(0)
	r := []rune(shortURL)

	for i := 0; i < len(shortURL); i++ {
		if 'a' <= r[i] && r[i] <= 'z' {
			id = id*div + uint64(r[i]) - 'a'
		}
		if 'A' <= r[i] && r[i] <= 'Z' {
			id = id*div + uint64(r[i]) - 'A' + 26
		}
		if '0' <= r[i] && r[i] <= '9' {
			id = id*div + uint64(r[i]) - '0' + 52
		}
	}

	return id
}
