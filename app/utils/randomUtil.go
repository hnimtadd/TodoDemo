package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
func GetRandomContentBatch(ctx context.Context, batchsize int, randommerKey string) ([]string, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", "https://randommer.io/api/Text/LoremIpsum?loremtype=business&type=paragraphs&number="+fmt.Sprint(batchsize), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-API-KEY", randommerKey)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	var result string
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	contents := strings.Split(result, "<br>")
	return contents, nil
}
