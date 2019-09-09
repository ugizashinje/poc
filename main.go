package main

import (
	"context"
	"fmt"

	"github.com/ugizashinje/epoc/failures"
)

func Parent(ctx context.Context, name string) (string, failures.SuperError) {
	fmt.Println("Parrent called")
	res, err := Child(ctx, name)

	if err != nil {
		return "", err
	}
	fmt.Println("Parent done")
	return res + "ic", nil
}

func Child(ctx context.Context, name string) (string, failures.SuperError) {
	fmt.Println("Child called")

	fmt.Println("Name is ", name)

	fmt.Println("Child done")
	return "", failures.SUBSCRIBER_DOES_NOT_EXSIS().WithInfo("DOB", "1990/12")
}

func main() {
	fmt.Println("Start")

	res, err := Parent(context.Background(), "Nikol")

	if err != nil {
		fmt.Println("AAA PANIC ATTACK ", err.Error())
	} else {
		fmt.Println("END ", res)
	}
}
