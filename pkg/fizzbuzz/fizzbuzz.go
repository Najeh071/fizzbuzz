package fizzbuzz

import (
	"errors"
	"fmt"
	"strconv"
)

// Parametre d'appel pour api FizzBuzz
type Request struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// Core logique de fizzbuzz
func Compute(req Request) ([]string, error) {

	//Input validation
	if req.Int1 <= 0 || req.Int2 <= 0 {
		return nil, errors.New("int1 and int2 must be greater than 0 please.")
	}
	
	if req.Limit <= 0 {
		return nil, errors.New("limit must be greater than 0 please.")
	}

	if req.Limit > 100 {
		return nil, fmt.Errorf("limit %d is too high (max 100)", req.Limit)
	}

	//construction de résultat et slices de la liste
	res := make([]string, 0, req.Limit)

	return res, nil
}