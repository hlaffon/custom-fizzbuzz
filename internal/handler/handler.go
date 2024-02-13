package handler

import (
	"fmt"
	"strconv"

	"custom-fizzbuzz/pkg/model"
)

type Handler struct {
}

func (h *Handler) PrintNumber(input *model.InputData) []string {
	var res []string
	for i := 1; i <= input.MaxNumber; i++ {
		var number string
		switch {
		case i%input.FirstMultiple == 0 && i%input.SecondMultiple == 0:
			number = fmt.Sprintf("%s%s", input.FirstAlias, input.SecondAlias)
		case i%input.FirstMultiple == 0:
			number = input.FirstAlias
		case i%input.SecondMultiple == 0:
			number = input.SecondAlias
		default:
			number = strconv.Itoa(i)
		}
		res = append(res, number)
	}
	return res
}
