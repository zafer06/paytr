package paytr

import (
	"os"
	"testing"
	"time"
)

var merchant = Merchant{
	ID:   os.Getenv("PAYTR_MERCHANT_ID"),
	Key:  os.Getenv("PAYTR_MERCHANT_KEY"),
	Salt: os.Getenv("PAYTR_MERCHANT_SALT"),
}

func TestGetInstallmentRates(t *testing.T) {
	var requestID = time.Now().Format("20060102T150405")

	var installments = GetInstallmentRates(requestID, merchant)

	if installments.Status != "success" {
		t.Errorf("Installments fonksiyonunda hatalar var. Hata: %s", installments.ErrMsg)
	}
}

func TestGetBasket(t *testing.T) {
	var productList = []Product{
		{"Örnek ürün 1", "18.00", 1},
		{"Örnek ürün 2", "33.25", 2},
		{"Örnek ürün 3", "45.42", 1},
	}

	var basket = GetBasket(productList)

	if len(basket) == 0 {
		t.Errorf("Basket formatı hatalı")
	}
}

func TestBinNumber(t *testing.T) {
	var binNumber = GetBinNumber("440274", merchant)

	if binNumber.Status != "success" {
		t.Errorf("BinNumber fonksiyonunda hatalar var. Hata: %s", binNumber.ErrMsg)
	}
}
