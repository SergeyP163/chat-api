package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ParseChatID(r *http.Request) (uint, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		return 0, errors.New("invalid path")
	}
	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
