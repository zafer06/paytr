package paytr

import (
	"os"
	"testing"
)

var merchant = Merchant{
	ID:   os.Getenv("PAYTR_MERCHANT_ID"),
	Key:  os.Getenv("PAYTR_MERCHANT_KEY"),
	Salt: os.Getenv("PAYTR_MERCHANT_SALT"),
}

func TestBinNumber(t *testing.T) {
	var binNumber = GetBinNumber("440274", merchant)

	if binNumber.Status != "success" {
		t.Errorf("BinNumber fonksiyonunda hatalar var. Hata: %s", binNumber.ErrMsg)
	}
}
