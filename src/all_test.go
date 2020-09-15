package main

import(
	"testing"
	"fmt"
	"math/big"
)

func TestGetAndPrintClosestPrimeNumber(t *testing.T) {
    total := GetAndPrintClosestPrimeNumber("15")
    if total != "13" {
       t.Errorf("GetAndPrintClosestPrimeNumber was incorrect, got: %s, expecting: %d.", total, 15)
    }
	total = GetAndPrintClosestPrimeNumber("9")
    if total != "7" {
       t.Errorf("GetAndPrintClosestPrimeNumber was incorrect, got: %s, expecting: %d.", total, 7)
    }
}

func TestGetPrimeLowerThan(t *testing.T) {
    var bignum, _ = new(big.Int).SetString("15", 0)
	total := fmt.Sprintf("%d", GetPrimeLowerThan(bignum))
    if total != "13" {
       t.Errorf("GetPrimeLowerThan was incorrect, got: %s, expecting: %d.", total, 15)
    }
	bignum, _ = new(big.Int).SetString("9", 0)
	total = fmt.Sprintf("%d", GetPrimeLowerThan(bignum))
    if total != "7" {
       t.Errorf("GetPrimeLowerThan was incorrect, got: %s, expecting: %d.", total, 7)
    }
}