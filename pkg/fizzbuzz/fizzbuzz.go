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

	for i := 1; i <= req.Limit; i++ {
		multiple1 := i%req.Int1 == 0
		multiple2 := i%req.Int2 == 0

		if multiple1 && multiple2 {
			res = append(res, req.Str1+req.Str2)
		} else if multiple1 {
			res = append(res, req.Str1)
		} else if multiple2 {
			res = append(res, req.Str2)
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}

	return res, nil
}