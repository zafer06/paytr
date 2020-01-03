package paytr

import (
	"fmt"
	"os"
	"testing"
)

var merchant = Merchant{
	ID:   os.Getenv("PAYTR_MERCHANT_ID"),
	Key:  os.Getenv("PAYTR_MERCHANT_KEY"),
	Salt: os.Getenv("PAYTR_MERCHANT_SALT"),
}

func TestGetBasket(t *testing.T) {
	var productList = []Product{
		{"Örnek ürün 1", "18.00", 1},
		{"Örnek ürün 2", "33.25", 2},
		{"Örnek ürün 3", "45.42", 1},
	}

	var basket = GetBasket(productList)

	fmt.Println("--> ", basket)

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
