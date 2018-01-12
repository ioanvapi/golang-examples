package main

// /* Required, but no C code needed. */
import "C"

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

//This function uses raw data (json.RawMessage is []byte).
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	var args map[string]interface{}
	if err := json.Unmarshal(evt, &args); err != nil {
		return nil, err
	}

	//I can see these logs in the main execution aws windows
	log.Printf("logged arguments: %v", args)

	//I can see this in CloudWatch only
	fmt.Printf("Arguments: %v", args)

	if _, ok := args["panic"]; ok {
		//cause lambda to fail its execution
		panic("panic from lambda")
	}

	if _, ok := args["error"]; ok {
		//cause lambda to fail its execution but the function returns a json
		return nil, fmt.Errorf("return error from lambda")
	}

	return "Well done", nil
}

type Proba struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Result struct {
	Signature string `json:"signature"`
}

func Handle2(evt Proba, ctx *runtime.Context) (Result, error) {
	return Result{
		Signature: fmt.Sprintf("Name %s, Age %d", evt.Name, evt.Age),
	}, nil
}
