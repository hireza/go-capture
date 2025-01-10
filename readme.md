# Go-Capture

[![License](https://img.shields.io/github/license/hireza/go-capture)](https://github.com/hireza/go-capture/blob/master/LICENSE)
[![CI](https://github.com/hireza/go-capture/workflows/CI/badge.svg)](https://github.com/hireza/go-capture/actions?query=workflow%3ACI)
[![CodeQL](https://github.com/hireza/go-capture/workflows/CodeQL/badge.svg)](https://github.com/hireza/go-capture/actions?query=workflow%3ACodeQL)  
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/fbacd3660b784ccfb6cda13e2c7c5029)](https://app.codacy.com/gh/hireza/go-capture/dashboard)
[![codecov](https://codecov.io/github/hireza/go-capture/graph/badge.svg?token=04DLRik7Qn)](https://codecov.io/github/hireza/go-capture)
[![Coverage Status](https://coveralls.io/repos/github/hireza/go-capture/badge.svg)](https://coveralls.io/github/hireza/go-capture)  
[![Go](https://img.shields.io/github/go-mod/go-version/hireza/go-capture)](https://github.com/hireza/go-capture)
[![Lines](https://img.shields.io/endpoint?url=https%3A%2F%2Fghloc.vercel.app%2Fapi%2Fhireza%2Fgo-capture%2Fbadge&color=blue)](https://github.com/hireza/go-capture)
[![Visitor](https://badges.pufler.dev/visits/hireza/go-capture)](https://github.com/hireza/go-capture)

**Go-Capture** is a lightweight Go package designed to capture `os.Stdout` and/or `os.Stderr`. This package is particularly useful for writing tests that involve outputs printed using the `fmt` package.

## 🔌 Installation

To install **Go-Capture**, use `go get` command below:

```bash
go get github.com/hireza/go-capture
```

## 💡 How To Use

Here is a basic example to use **Go-Capture**:

```go
package main

import (
	"fmt"
	"github.com/hireza/go-capture"
)

func printSomethingString() {
	fmt.Println("Hello, Go-Capture!")
}

func printSomethingSliceInt() {
	fmt.Println([]int{1, 2, 3})
}

func main() {
	// Example using method PipeDirectly
	// PipeDirectly uses the default pipe-based capturing (may block on large output).
	output := capture.UseMethod(capture.PipeDirectly).Output(func() {
		 printSomethingString()
	})
	fmt.Println("Captured Output as String:", output.AsString())

	// Example using method PipeWithGoroutine
	// PipeWithGoroutine uses a goroutine to read and buffer data, avoiding blocking.
	output = capture.UseMethod(capture.PipeWithGoroutine).Output(func() {
		fmt.Println("Hello from Goroutine!")
	})
	fmt.Println("Captured Output as String:", output.AsString())

	// Simplified example without specifying a method (default: PipeDirectly)
	output = capture.Output(func() {
		fmt.Println("Simplified Capture Example")
	})
	fmt.Println("Captured Output as String:", output.AsString())

	// You can convert it to the other data type
	// Use .AsBool(), .AsInt(), .AsByte(), etc...
	// Check the complete method on main.go
	output = capture.Output(func() {
		printSomethingSliceInt()
	})

	res, err := output.AsSliceInt()
	if err != nil {
		fmt.Errorf("failed to convert output to slice of integers: %v", err)
		return
	}
	fmt.Println("Captured Output as Slice of Integers:", res)
}
```

## 🧪 Running Tests

To ensure the solutions are correct, the repository includes a comprehensive test suite. Run the tests using the following command:

```bash
go test ./... -v
```

## 🛠️ Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Write your code and add tests.
4. Submit a pull request with a clear description of your changes.

### Coding Guidelines

- Ensure your code is properly formatted using `gofmt`.
- Write clear, concise, and well-documented code.
- Maintain a high level of test coverage.

Feel free to open issues or start discussions to improve the library further!

## 📧 Contact

For questions or suggestions, feel free to reach out:

- **LinkedIn**: [@hireza](https://www.linkedin.com/in/hireza)
- **Email**: [say@hireza.top](mailto:say@hireza.top)

---

### ⭐ If you find this repository helpful, please give it a star
