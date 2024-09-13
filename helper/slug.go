package helper

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func GenerateSlug(title string) string {
	slug := strings.ToLower(title)

	rand.Seed(time.Now().UnixNano())

	slug = strings.ReplaceAll(slug, " ", fmt.Sprintf("_%d", rand.Intn(99)+1))

	reg := regexp.MustCompile(`[^\w]+`)
	slug = reg.ReplaceAllString(slug, "")

	return slug
}
