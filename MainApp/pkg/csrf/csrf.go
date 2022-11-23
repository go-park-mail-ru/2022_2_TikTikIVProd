package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const secretKey = "ws_secret"

func CreateCSRF(session string) (string, error) {
	h := hmac.New(sha256.New, []byte(secretKey))
	tokenExpTime := time.Now().Add(time.Hour).Unix()
	data := fmt.Sprintf("%s:%d", session, tokenExpTime)
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	return token, nil
}

func CheckCSRF(session string, inputToken string) (bool, error) {
	tokenData := strings.Split(inputToken, ":")
	if len(tokenData) != 2 {
		return false, errors.New("bad token data")
	}

	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false, errors.New("bad token time")
	}

	if tokenExp < time.Now().Unix() {
		return false, errors.New("token expired")
	}

	h := hmac.New(sha256.New, []byte(secretKey))
	data := fmt.Sprintf("%s:%d", session, tokenExp)
	h.Write([]byte(data))
	if err != nil {
		return false, err
	}
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false, err
	}

	return hmac.Equal(messageMAC, expectedMAC), nil
}