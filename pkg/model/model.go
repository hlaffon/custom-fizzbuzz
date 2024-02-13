package model

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"custom-fizzbuzz/pkg/pb"
)

const (
	defaultMaxNumber      = 200
	defaultFirstMultiple  = 3
	defaultSecondMultiple = 5
	defaultFirstAlias     = "Fizz"
	defaultSecondAlias    = "Buzz"
)

type RequestData struct {
	MaxNumber      *int    `json:"maxNumber"`
	FirstMultiple  *int    `json:"firstMultiple"`
	SecondMultiple *int    `json:"secondMultiple"`
	FirstAlias     *string `json:"firstAlias"`
	SecondAlias    *string `json:"secondAlias"`
}

type InputData struct {
	MaxNumber      int
	FirstMultiple  int
	SecondMultiple int
	FirstAlias     string
	SecondAlias    string
}

func FromRequestData(request *RequestData) *InputData {
	data := defaultInput()
	if request == nil {
		return data
	}
	if request.MaxNumber != nil {
		data.MaxNumber = *request.MaxNumber
	}
	if request.FirstMultiple != nil {
		data.FirstMultiple = *request.FirstMultiple
	}
	if request.SecondMultiple != nil {
		data.SecondMultiple = *request.SecondMultiple
	}
	if request.FirstAlias != nil {
		data.FirstAlias = *request.FirstAlias
	}
	if request.SecondAlias != nil {
		data.SecondAlias = *request.SecondAlias
	}
	return data
}

func FromProto(request *pb.Request) *InputData {
	data := defaultInput()
	if request == nil {
		return data
	}

	if request.GetMaxNumber() != nil {
		data.MaxNumber = int32ValueToInt(request.GetMaxNumber())
	}
	if request.GetFirstMultiple() != nil {
		data.FirstMultiple = int32ValueToInt(request.GetFirstMultiple())
	}
	if request.GetSecondMultiple() != nil {
		data.SecondMultiple = int32ValueToInt(request.GetSecondMultiple())
	}
	if request.GetFirstAlias() != nil {
		data.FirstAlias = stringValueToString(request.GetFirstAlias())
	}
	if request.GetSecondAlias() != nil {
		data.SecondAlias = stringValueToString(request.GetSecondAlias())
	}
	return data
}

func ToProto(data []string) *pb.Response {
	return &pb.Response{Numbers: data}
}

func (i *InputData) Validate() error {
	switch {
	case i.MaxNumber < 1:
		return fmt.Errorf("max number must be greater than zero")
	case i.FirstMultiple == 0:
		return fmt.Errorf("first multiple cannot be zero")
	case i.SecondMultiple == 0:
		return fmt.Errorf("second multiple cannot be zero")
	default:
		return nil
	}
}

func defaultInput() *InputData {
	return &InputData{
		MaxNumber:      defaultMaxNumber,
		FirstMultiple:  defaultFirstMultiple,
		SecondMultiple: defaultSecondMultiple,
		FirstAlias:     defaultFirstAlias,
		SecondAlias:    defaultSecondAlias,
	}
}

func int32ValueToInt(val *wrappers.Int32Value) int {
	return int(val.GetValue())
}

func stringValueToString(val *wrappers.StringValue) string {
	return val.GetValue()
}
