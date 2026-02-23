package utils

import (
	"math/rand"
	"strings"
	"time"
)

type WARequest struct {
	APIKey    string `json:"api_key"`
	NumberKey string `json:"number_key"`
	PhoneNo   string `json:"phone_no"`
	Message   string `json:"message"`
}

type WAImageRequest struct {
	APIKey          string `json:"api_key"`
	NumberKey       string `json:"number_key"`
	PhoneNo         string `json:"phone_no"`
	URL             string `json:"url"`
	Message         string `json:"message"`
	SeparateCaption int    `json:"separate_caption"`
}

type WAFileRequest struct {
	APIKey    string `json:"api_key"`
	NumberKey string `json:"number_key"`
	PhoneNo   string `json:"phone_no"`
	URL       string `json:"url"`
}

// WAOTPResponse adalah response dari API wa
type WAResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	WorkerBy string `json:"worker_by"`
	Ack      string `json:"ack"`
}

func GenerateNumericOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	numbers := "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(otp)
}

func FormatPhoneNumber(phone string) string {
	// Hilangkan spasi dan simbol non-digit
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Jika diawali dengan "0", ganti dengan "62"
	if strings.HasPrefix(phone, "0") {
		phone = "62" + phone[1:]
	} else if strings.HasPrefix(phone, "+62") {
		phone = "62" + phone[3:]
	}

	return phone
}
