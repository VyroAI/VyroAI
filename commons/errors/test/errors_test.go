package test

import (
	"fmt"
	"github.com/vyroai/VyroAI/commons/errors"
	"testing"
)

func compare(int1, int2 int) (bool, error) {
	if int1 == int2 {
		return true, nil
	} else {
		return false, fmt.Errorf("int1 does not equal to int2")
	}
}

func TestInvalidError(t *testing.T) {
	results, err := compare(1, 2)
	if err != nil {
		err = errors.ErrInvalid.Wrap(err, "compare function failed")
		t.Error(err)
		fmt.Println(err)
		return
	}

	fmt.Println(results)
}

func domainCompareFunction(int1, int2 int) (bool, error) {
	var err error
	if int1 == 0 || int2 == 0 {
		err = errors.ErrForbidden.New("missing arguments")
		return false, err
	}

	result, err := compare(int1, int2)
	if err != nil {
		err = errors.ErrInvalid.Wrap(err, "compare function failed")
		fmt.Println(err)
		return false, err
	}
	return result, nil
}

func TestHandleInvalidError(t *testing.T) {
	//error from domain
	result, err := domainCompareFunction(1, 1)
	if err != nil {
		//print the error type
		fmt.Println(errors.GetType(err))

		//handle error on controller layer e.g http server
		switch errors.GetType(err) {
		case errors.ErrInvalid:
			fmt.Println("bad request")
		case errors.ErrForbidden:
			fmt.Println("missing token")

		default:
			fmt.Println("internal server error")
		}

	}
	t.Log(result)

}

func TestHandleError(t *testing.T) {
	//error from domain
	result, err := domainCompareFunction(0, 0)
	if err != nil {
		//print the error type
		fmt.Println(errors.GetType(err))

		//handle error on controller layer e.g http server
		switch errors.GetType(err) {
		case errors.ErrInvalid:
			fmt.Println("bad request")
			//get error message
			fmt.Println(err.Error())
		case errors.ErrForbidden:
			fmt.Println("missing token")
			//get error message
			fmt.Println(err.Error())

		default:
			fmt.Println("internal server error")
		}

	}
	t.Log(result)

}
