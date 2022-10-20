package main

func foo() (string, error) {
	return "", nil
}

func bar() (string, error) {
	return "", nil
}

func baz(x string, val bool) (string, error) {
	return "", nil
}

func redeem() (resp string, err error) {
	defer func() {
		resp, err = foo()
	}()
	_, err = bar()
	// err check but no return

	go func() {
		baz("abcdef", err != nil)
	}()
	return // the defer function runs after here
}

func main() {
	_, _ = redeem()
}
