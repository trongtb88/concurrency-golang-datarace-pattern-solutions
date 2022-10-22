package main

func foo() (string, error) {
	return "", nil
}

func bar() (string, error) {
	return "", nil
}

func baz() (string, error) {
	return "", nil
}

// ORIGINAL
// re-use the err insdie the go routine will cause data race
/*
func main() {
	_, err := foo()
	if err != nil {
		// do something
	}

	go func() {
		_, err = bar()
		if err != nil {
			// do something
		}
	}()

	_, err = baz()
	if err != nil {
		// do something
	}
}
*/

// FIX
func main() {
	_, err := foo()
	if err != nil {
		// do something
	}

	go func() {
		_, err := bar()
		if err != nil {
			// do something
		}
	}()

	_, err = baz()
	if err != nil {
		// do something
	}
}
