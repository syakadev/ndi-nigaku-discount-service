package utils

import (
	"encoding/json"
	"fmt"
	"os"
	servicemodel "service/discount/api/model/service"
	"time"

	"github.com/valyala/fasthttp"
)

type LayananWAResponse struct {
	StatusCode int    `json:"status_code"`
	Status     bool   `json:"status"`
	Message    string `json:"message"`
}

func FetchLevelName(idLevel string) (string, error) {
	url := fmt.Sprintf(os.Getenv("URL_RBAC")+"/level/nama-level/%s", idLevel)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_KEY"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &fasthttp.Client{}
	err := client.DoTimeout(req, resp, 3*time.Second)
	if err != nil || resp.StatusCode() != fasthttp.StatusOK {
		return "", fmt.Errorf("gagal request: %v", err)
	}

	var data servicemodel.FetchLevelName
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return "", err
	}

	return data.Name, nil
}

func SendTextWAMessage(phoneNo, message string) (bool, error) {
	url := fmt.Sprintf(os.Getenv("URL_LAYANAN") + "/whatsapp/send-text/")

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// Buat body JSON
	bodyRequest := map[string]interface{}{
		"message":      message,
		"phone_no":     phoneNo,
		"nama_layanan": "akudicari",
	}

	bodyBytes, err := json.Marshal(bodyRequest)
	if err != nil {
		return false, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_KEY"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBody(bodyBytes)

	client := &fasthttp.Client{}
	err = client.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		return false, fmt.Errorf("gagal request: %v", err)
	}

	var data LayananWAResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return false, err
	}

	if data.StatusCode != 200 {
		return false, fmt.Errorf("gagal request v2: %v", data.Message)
	}

	return true, nil
}

func FetchProductName(id string) (string, error) {
	url := fmt.Sprintf(os.Getenv("URL_INVENTORY")+"/product/name/%s", id)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_KEY"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &fasthttp.Client{}
	err := client.DoTimeout(req, resp, 5*time.Second)
	if err != nil || resp.StatusCode() != fasthttp.StatusOK {
		return "", fmt.Errorf("gagal request: %v", err)
	}

	var data servicemodel.FetchName
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return "", err
	}

	return data.Name, nil
}

func FetchCustomerName(id string) (string, error) {
	url := fmt.Sprintf(os.Getenv("URL_TRANSACTION")+"/customer/%s", id)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET_KEY"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &fasthttp.Client{}
	err := client.DoTimeout(req, resp, 5*time.Second)
	if err != nil || resp.StatusCode() != fasthttp.StatusOK {
		return "", fmt.Errorf("gagal request: %v", err)
	}

	var data servicemodel.FetchName
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return "", err
	}

	return data.Name, nil
}
