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
	"strconv"
)

const baseURL = "https://www.paytr.com/odeme/"

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

	// Installment model
	Installment struct {
		Status    string `json:"status"`
		RequestID string `json:"request_id"`
		ErrMsg    string `json:"err_msg"`
		Rates     struct {
			World struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"world"`
			Axess struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"axess"`
			Maximum struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"maximum"`
			Cardfinans struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"cardfinans"`
			Paraf struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"paraf"`
			Advantage struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"advantage"`
			Combo struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"combo"`
			Bonus struct {
				Taksit2  float64 `json:"taksit_2"`
				Taksit3  float64 `json:"taksit_3"`
				Taksit4  float64 `json:"taksit_4"`
				Taksit5  float64 `json:"taksit_5"`
				Taksit6  float64 `json:"taksit_6"`
				Taksit7  float64 `json:"taksit_7"`
				Taksit8  float64 `json:"taksit_8"`
				Taksit9  float64 `json:"taksit_9"`
				Taksit10 float64 `json:"taksit_10"`
				Taksit11 float64 `json:"taksit_11"`
				Taksit12 float64 `json:"taksit_12"`
			} `json:"bonus"`
		} `json:"oranlar"`
	}
)

// GetInstallmentRates function
func GetInstallmentRates(requestID string, merchant Merchant) Installment {
	var hashStr = merchant.ID + requestID + merchant.Salt
	var token = getToken(hashStr, merchant.Key)

	var values = url.Values{
		"merchant_id": {merchant.ID},
		"request_id":  {requestID},
		"paytr_token": {token},
	}

	var res = connect(baseURL+"taksit-oranlari", values)

	var installment Installment
	err := json.Unmarshal(res, &installment)
	if err != nil {
		log.Println("PayTR istek durumu: ", installment.Status)
		log.Println("PayTR token istek sonucu: ", err)
	}

	return installment
}

// GetBinNumber function
func GetBinNumber(binNumber string, merchant Merchant) BinNumber {
	var hashStr = binNumber + merchant.ID + merchant.Salt
	var token = getToken(hashStr, merchant.Key)

	var values = url.Values{
		"merchant_id": {merchant.ID},
		"bin_number":  {binNumber},
		"paytr_token": {token},
	}

	var res = connect(baseURL+"api/bin-detail", values)

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

// GetBasket ürünleri uygun formatta bir dizge haline getirir.
func GetBasket(products []Product) string {
	var list string
	for _, p := range products {
		list += getEncodedProduct(p.Name, p.Price, p.Quantity)
		list += ","
	}
	list = list[0 : len(list)-1]

	return "[" + list + "]"
}

// Ürün bilglerindeki Türkçe karakterleri encode eder.
func getEncodedProduct(name string, price string, quantity int) string {
	return "[" +
		strconv.QuoteToASCII(name) + "," +
		strconv.QuoteToASCII(price) + "," +
		strconv.Itoa(quantity) + "]"
}

// getToken BinNumber sorgusu için gerekli kimlik dogrulama token dizesini oluşturur.
func getToken(hashStr string, merchantKey string) string {
	h := hmac.New(sha256.New, []byte(merchantKey))
	h.Write([]byte(hashStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// PayTR api uzerinden bir istek gonderir.
func connect(url string, values url.Values) []byte {
	res, err := http.PostForm(url, values)
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
