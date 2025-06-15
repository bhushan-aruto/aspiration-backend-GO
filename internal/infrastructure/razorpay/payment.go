package razorpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type RazorPayRepo struct {
	Url        string
	KeyId      string
	KeySecrete string
}

func NewRazorPayRepo(url, keyId, keySecrete string) *RazorPayRepo {
	return &RazorPayRepo{
		Url:        url,
		KeyId:      keyId,
		KeySecrete: keySecrete,
	}
}

func (r *RazorPayRepo) CreateOrder(amount int32) ([]byte, error) {
	amountPaise := amount * 100

	receiptId := uuid.New().String()

	payload := map[string]interface{}{
		"amount":          amountPaise,
		"currency":        "INR",
		"receipt":         receiptId,
		"payment_capture": 1,
	}

	jsonPayload, _ := json.Marshal(payload)

	httpClient := &http.Client{}

	httpRequest, _ := http.NewRequest("POST", r.Url, bytes.NewBuffer(jsonPayload))
	httpRequest.SetBasicAuth(r.KeyId, r.KeySecrete)
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	return respBody, nil

}

func (r *RazorPayRepo) VerifyOrder(orderID, paymentID, razorpaySignature, secret string) bool {
	data := orderID + "|" + paymentID
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(expectedSignature), []byte(razorpaySignature))
}
