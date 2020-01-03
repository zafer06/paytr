package paytr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "https://www.paytr.com/"

type (
	// Merchant model
	Merchant struct {
		ID   string
		Key  string
		Salt string
	}

	// Product model
	Product struct {
		Name     string
		Price    string
		Quantity int
	}

	// BinNumber model
	BinNumber struct {
		Status       string `json:"status"`
		Brand        string `json:"brand"`
		CardType     string `json:"cardType"`
		BusinessCard string `json:"businessCard"`
		Bank         string `json:"bank"`
		Schema       string `json:"schema"`
		ErrMsg       string `json:"err_msg"`
	}
)

// GetBinNumber function
func GetBinNumber(binNumber string, merchant Merchant) BinNumber {
	var hashStr = binNumber + merchant.ID + merchant.Salt
	var token = tokenForBinNumber(hashStr, merchant.Key)

	var values = url.Values{
		"merchant_id": {merchant.ID},
		"bin_number":  {binNumber},
		"paytr_token": {token},
	}

	var res = connect(values)

	var bin BinNumber
	err := json.Unmarshal(res, &bin)
	if err != nil {
		log.Println("PayTR istek durumu: ", bin.Status)
		log.Println("PayTR token istek sonucu: ", err)
	}

	return bin
}

// TokenForPayment Kimlik dogrulama icin gerekli token dizesini oluşturur.
func TokenForPayment(hashStr string, merchantSalt string, merchantKey string) string {
	h := hmac.New(sha256.New, []byte(merchantKey))
	h.Write([]byte(hashStr))
	h.Write([]byte(merchantSalt))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// CheckHash paytr tarafından gönderilen bildirim mesajını doğrulamak icin hash doğrulaması yapar.
func CheckHash(hashData map[string]string, merchantKey string, merchantSalt string) (bool, string) {
	h := hmac.New(sha256.New, []byte(merchantKey))
	h.Write([]byte(hashData["merchant_oid"]))
	h.Write([]byte(merchantSalt))
	h.Write([]byte(hashData["status"]))
	h.Write([]byte(hashData["total_amount"]))
	var hash = base64.StdEncoding.EncodeToString(h.Sum(nil))

	var check = false
	if hash == hashData["hash"] {
		check = true
	}

	return check, hash
}

// tokenForBinNumber BinNumber sorgusu için gerekli kimlik dogrulama token dizesini oluşturur.
func tokenForBinNumber(hashStr string, merchantKey string) string {
	h := hmac.New(sha256.New, []byte(merchantKey))
	h.Write([]byte(hashStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// PayTR api uzerinden bir istek gonderir.
func connect(values url.Values) []byte {
	res, err := http.PostForm(baseURL+"odeme/api/bin-detail", values)
	if err != nil {
		fmt.Println("--> ", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}
