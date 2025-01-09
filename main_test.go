package capture_test

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"

	"github.com/hireza/go-capture"
)

func TestCaptureMethodStdout(t *testing.T) {
	tests := []struct {
		name           string
		method         capture.BufferMethod
		expectedOutput string
	}{
		{
			name:           "Test PipeDirectly",
			method:         capture.PipeDirectly,
			expectedOutput: "Hello, direct capture!\n",
		},
		{
			name:           "Test PipeWithGoroutine",
			method:         capture.PipeWithGoroutine,
			expectedOutput: "Hello, goroutine capture!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of the function execution
			output := capture.UseMethod(tt.method).Stdout(func() {
				// Generate the output for testing
				if tt.method == capture.PipeDirectly {
					fmt.Println("Hello, direct capture!")
				} else if tt.method == capture.PipeWithGoroutine {
					fmt.Println("Hello, goroutine capture!")
				}
			})

			// Compare the captured output with the expected output
			if output.Value != tt.expectedOutput {
				t.Errorf("Expected %q but got %q", tt.expectedOutput, output.Value)
			}
		})
	}
}

func TestCaptureMethodStderr(t *testing.T) {
	tests := []struct {
		name           string
		method         capture.BufferMethod
		expectedOutput string
	}{
		{
			name:           "Test PipeDirectly",
			method:         capture.PipeDirectly,
			expectedOutput: "Hello, direct capture!\n",
		},
		{
			name:           "Test PipeWithGoroutine",
			method:         capture.PipeWithGoroutine,
			expectedOutput: "Hello, goroutine capture!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of the function execution
			output := capture.UseMethod(tt.method).Stderr(func() {
				// Generate the output for testing
				if tt.method == capture.PipeDirectly {
					fmt.Fprintln(os.Stderr, "Hello, direct capture!")
				} else if tt.method == capture.PipeWithGoroutine {
					fmt.Fprintln(os.Stderr, "Hello, goroutine capture!")
				}
			})

			// Compare the captured output with the expected output
			if output.Value != tt.expectedOutput {
				t.Errorf("Expected %q but got %q", tt.expectedOutput, output.Value)
			}
		})
	}
}

func TestCaptureMethodOutput(t *testing.T) {
	tests := []struct {
		name           string
		method         capture.BufferMethod
		expectedOutput string
	}{
		{
			name:           "Test PipeDirectly",
			method:         capture.PipeDirectly,
			expectedOutput: "Hello, direct capture!\n",
		},
		{
			name:           "Test PipeWithGoroutine",
			method:         capture.PipeWithGoroutine,
			expectedOutput: "Hello, goroutine capture!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of the function execution
			output := capture.UseMethod(tt.method).Output(func() {
				// Generate the output for testing
				if tt.method == capture.PipeDirectly {
					fmt.Println("Hello, direct capture!")
				} else if tt.method == capture.PipeWithGoroutine {
					fmt.Println("Hello, goroutine capture!")
				}
			})

			// Compare the captured output with the expected output
			if output.Value != tt.expectedOutput {
				t.Errorf("Expected %q but got %q", tt.expectedOutput, output.Value)
			}
		})
	}
}

func TestCaptureStdout(t *testing.T) {
	output := capture.Stdout(func() {
		fmt.Println("Hello, stdout!")
	})

	expected := "Hello, stdout!\n"
	if output.Value != expected {
		t.Errorf("Expected output to be %q, but got %q", expected, output.Value)
	}
}

func TestCaptureStderr(t *testing.T) {
	output := capture.Stderr(func() {
		fmt.Fprintln(os.Stderr, "Hello, stderr!")
	})

	expected := "Hello, stderr!\n"
	if output.Value != expected {
		t.Errorf("Expected output to be %q, but got %q", expected, output.Value)
	}
}

func TestCaptureOutput(t *testing.T) {
	output := capture.Output(func() {
		fmt.Println("Hello, stdout!")
		fmt.Fprintln(os.Stderr, "Hello, stderr!")
	})

	expected := "Hello, stdout!\nHello, stderr!\n"
	if output.Value != expected {
		t.Errorf("Expected output to be %q, but got %q", expected, output.Value)
	}
}

func TestCapturedOutputAsBool(t *testing.T) {
	tests := []struct {
		input   string
		want    bool
		wantErr bool
	}{
		{"true", true, false},
		{"false", false, false},
		{"TRUE", true, false},
		{"FALSE", false, false},
		{"True", true, false},
		{"False", false, false},
		{"1", true, false},
		{"t", true, false},
		{"T", true, false},
		{"0", false, false},
		{"f", false, false},
		{"F", false, false},
		{"tRuE", false, true},
		{"fAlSe", false, true},
		{"invalid", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsString(t *testing.T) {
	input := "Hello, world!"
	output := capture.Result{Value: input}
	if output.AsString() != input {
		t.Errorf("AsString() got = %q, want %q", output.AsString(), input)
	}
}

func TestCapturedOutputAsInt(t *testing.T) {
	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{"42", 42, false},
		{"-10", -10, false},
		{"0", 0, false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsInt8(t *testing.T) {
	tests := []struct {
		input   string
		want    int8
		wantErr bool
	}{
		{"127", 127, false},
		{"-128", -128, false},
		{"128", 127, true},
		{"-129", -128, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsInt8()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsInt8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsInt16(t *testing.T) {
	tests := []struct {
		input   string
		want    int16
		wantErr bool
	}{
		{"32767", 32767, false},
		{"-32768", -32768, false},
		{"32768", 32767, true},
		{"-32769", -32768, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsInt16()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsInt32(t *testing.T) {
	tests := []struct {
		input   string
		want    int32
		wantErr bool
	}{
		{"2147483647", 2147483647, false},
		{"-2147483648", -2147483648, false},
		{"2147483648", 2147483647, true},
		{"-2147483649", -2147483648, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsInt32()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsInt64(t *testing.T) {
	tests := []struct {
		input   string
		want    int64
		wantErr bool
	}{
		{"9223372036854775807", 9223372036854775807, false},
		{"-9223372036854775808", -9223372036854775808, false},
		{"9223372036854775808", 9223372036854775807, true},
		{"-9223372036854775809", -9223372036854775808, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsInt64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsUint(t *testing.T) {
	tests := []struct {
		input   string
		want    uint
		wantErr bool
	}{
		{"42", 42, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"18446744073709551615", 18446744073709551615, false},
		{"18446744073709551616", 18446744073709551615, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsUint()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsUint8(t *testing.T) {
	tests := []struct {
		input   string
		want    uint8
		wantErr bool
	}{
		{"255", 255, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"256", 255, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsUint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsUint16(t *testing.T) {
	tests := []struct {
		input   string
		want    uint16
		wantErr bool
	}{
		{"65535", 65535, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"65536", 65535, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsUint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsUint32(t *testing.T) {
	tests := []struct {
		input   string
		want    uint32
		wantErr bool
	}{
		{"4294967295", 4294967295, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"4294967296", 4294967295, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsUint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsUint64(t *testing.T) {
	tests := []struct {
		input   string
		want    uint64
		wantErr bool
	}{
		{"18446744073709551615", 18446744073709551615, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"18446744073709551616", 18446744073709551615, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsUint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsUintptr(t *testing.T) {
	tests := []struct {
		input   string
		want    uintptr
		wantErr bool
	}{
		{"42", uintptr(42), false},
		{"0", uintptr(0), false},
		{"-1", uintptr(0), true},
		{"18446744073709551615", uintptr(18446744073709551615), false},
		{"invalid", uintptr(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsUintptr()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsUintptr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsByte(t *testing.T) {
	tests := []struct {
		input   string
		want    byte
		wantErr bool
	}{
		{"255", 255, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"256", 255, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsRune(t *testing.T) {
	tests := []struct {
		input   string
		want    rune
		wantErr bool
	}{
		{"65", 'A', false},
		{"0", 0, false},
		{"-1", -1, false},
		{"1114111", 1114111, false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsRune()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsRune() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsFloat32(t *testing.T) {
	tests := []struct {
		input   string
		want    float32
		wantErr bool
	}{
		{"3.14159", 3.14159, false},
		{"-2.71828", -2.71828, false},
		{"0.0", 0.0, false},
		{"inf", float32(math.Inf(1)), false},
		{"-inf", float32(math.Inf(-1)), false},
		{"NaN", float32(math.NaN()), false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsFloat32()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if math.IsNaN(float64(got)) && math.IsNaN(float64(tt.want)) {
				return
			}

			if math.Abs(float64(got)-float64(tt.want)) > 1e-6 {
				t.Errorf("AsFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsFloat64(t *testing.T) {
	tests := []struct {
		input   string
		want    float64
		wantErr bool
	}{
		{"3.141592653589793", 3.141592653589793, false},
		{"-2.718281828459045", -2.718281828459045, false},
		{"0.0", 0.0, false},
		{"inf", math.Inf(1), false},
		{"-inf", math.Inf(-1), false},
		{"NaN", math.NaN(), false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsFloat64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !math.IsNaN(tt.want) && !math.IsNaN(got) && got != tt.want {
				t.Errorf("AsFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsComplex64(t *testing.T) {
	tests := []struct {
		input   string
		want    complex64
		wantErr bool
	}{
		{"3+4i", complex(3, 4), false},
		{"-2-1.5i", complex(-2, -1.5), false},
		{"0+0i", 0, false},
		{"1i", complex(0, 1), false},
		{"-2i", complex(0, -2), false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsComplex64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsComplex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsComplex64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsComplex128(t *testing.T) {
	tests := []struct {
		input   string
		want    complex128
		wantErr bool
	}{
		{"3+4i", complex(3, 4), false},
		{"-2-1.5i", complex(-2, -1.5), false},
		{"0+0i", 0, false},
		{"1i", complex(0, 1), false},
		{"-2i", complex(0, -2), false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsComplex128()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsComplex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsComplex128() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceInt(t *testing.T) {
	tests := []struct {
		input   string
		want    []int
		wantErr bool
	}{
		{"[1, 2, 3]", []int{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []int{0, 0, 1}, true},
		{"[]", []int{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceInt8(t *testing.T) {
	tests := []struct {
		input   string
		want    []int8
		wantErr bool
	}{
		{"[1, 2, 3]", []int8{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []int8{0, 0, 1}, true},
		{"[]", []int8{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceInt8()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceInt8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceInt16(t *testing.T) {
	tests := []struct {
		input   string
		want    []int16
		wantErr bool
	}{
		{"[1, 2, 3]", []int16{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []int16{0, 0, 1}, true},
		{"[]", []int16{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceInt16()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceInt32(t *testing.T) {
	tests := []struct {
		input   string
		want    []int32
		wantErr bool
	}{
		{"[1, 2, 3]", []int32{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []int32{0, 0, 1}, true},
		{"[]", []int32{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceInt32()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceInt64(t *testing.T) {
	tests := []struct {
		input   string
		want    []int64
		wantErr bool
	}{
		{"[1, 2, 3]", []int64{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []int64{0, 0, 1}, true},
		{"[]", []int64{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceInt64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceUint(t *testing.T) {
	tests := []struct {
		input   string
		want    []uint
		wantErr bool
	}{
		{"[1, 2, 3]", []uint{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []uint{0, 0, 1}, true},
		{"[]", []uint{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceUint()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceUint8(t *testing.T) {
	tests := []struct {
		input   string
		want    []uint8
		wantErr bool
	}{
		{"[1, 2, 3]", []uint8{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []uint8{0, 0, 1}, true},
		{"[]", []uint8{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceUint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceUint16(t *testing.T) {
	tests := []struct {
		input   string
		want    []uint16
		wantErr bool
	}{
		{"[1, 2, 3]", []uint16{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []uint16{0, 0, 1}, true},
		{"[]", []uint16{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceUint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceUint32(t *testing.T) {
	tests := []struct {
		input   string
		want    []uint32
		wantErr bool
	}{
		{"[1, 2, 3]", []uint32{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []uint32{0, 0, 1}, true},
		{"[]", []uint32{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceUint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceUint64(t *testing.T) {
	tests := []struct {
		input   string
		want    []uint64
		wantErr bool
	}{
		{"[1, 2, 3]", []uint64{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []uint64{0, 0, 1}, true},
		{"[]", []uint64{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceUint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceUintptr(t *testing.T) {
	tests := []struct {
		input   string
		want    []uintptr
		wantErr bool
	}{
		{"[1, 2, 3]", []uintptr{1, 2, 3}, false},
		{"[\"hello\", \"world\", 1]", []uintptr{0, 0, 1}, true},
		{"[]", []uintptr{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceUintptr()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceUintptr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceByte(t *testing.T) {
	tests := []struct {
		input   string
		want    []byte
		wantErr bool
	}{
		{"[97, 98, 99]", []byte{'a', 'b', 'c'}, false},
		{"[\"hello\", \"world\", 1]", []byte{0, 0, 1}, true},
		{"[]", []byte{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceRune(t *testing.T) {
	tests := []struct {
		input   string
		want    []rune
		wantErr bool
	}{
		{"[\"a\", \"b\", \"c\"]", []rune{0, 0, 0}, true},
		{"[\"hello\", \"world\", 1]", []rune{0, 0, 1}, true},
		{"[]", []rune{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceRune()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceRune() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceFloat32(t *testing.T) {
	tests := []struct {
		input   string
		want    []float32
		wantErr bool
	}{
		{"[1.0, 2.5, 3.14]", []float32{1.0, 2.5, 3.14}, false},
		{"[\"hello\", \"world\", 1]", []float32{0, 0, 1}, true},
		{"[]", []float32{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceFloat32()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("AsSliceFloat32() got = %v, want %v", got, tt.want)
				return
			}

			for i := range got {
				diff := math.Abs(float64(got[i]) - float64(tt.want[i]))
				if diff > 1e-6 {
					break
				}
			}
		})
	}
}

func TestCapturedOutputAsSliceFloat64(t *testing.T) {
	tests := []struct {
		input   string
		want    []float64
		wantErr bool
	}{
		{"[1.0, 2.5, 3.14]", []float64{1.0, 2.5, 3.14}, false},
		{"[\"hello\", \"world\", 1]", []float64{0, 0, 1}, true},
		{"[]", []float64{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceFloat64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range got {
				if !math.IsNaN(tt.want[i]) && !math.IsNaN(got[i]) && got[i] != tt.want[i] {
					t.Errorf("AsSliceFloat64() got = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestCapturedOutputAsSliceComplex64(t *testing.T) {
	tests := []struct {
		input   string
		want    []complex64
		wantErr bool
	}{
		{"[]", []complex64{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceComplex64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceComplex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceComplex64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceComplex128(t *testing.T) {
	tests := []struct {
		input   string
		want    []complex128
		wantErr bool
	}{
		{"[]", []complex128{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceComplex128()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceComplex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceComplex128() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceBool(t *testing.T) {
	tests := []struct {
		input   string
		want    []bool
		wantErr bool
	}{
		{"[true, false, true]", []bool{true, false, true}, false},
		{"[]", []bool{}, false},
		{"invalid", nil, true},
		{"[true, false, true]", []bool{true, false, true}, false},
		{"[]", []bool{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapturedOutputAsSliceString(t *testing.T) {
	tests := []struct {
		input   string
		want    []string
		wantErr bool
	}{
		{"[\"hello\", \"world\"]", []string{"hello", "world"}, false},
		{"[\"hello\", \"world\", \"2\"]", []string{"hello", "world", "2"}, false},
		{"[]", []string{}, false},
		{"invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := capture.Result{Value: tt.input}
			got, err := output.AsSliceString()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsSliceString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsSliceString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
