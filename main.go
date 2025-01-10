package capture

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// BufferMethod defines the approach used to handle captured data.
type BufferMethod int

const (
	// PipeDirectly uses the default pipe-based capturing (may block on large output).
	PipeDirectly BufferMethod = iota

	// PipeWithGoroutine uses a goroutine to read and buffer data, avoiding blocking.
	PipeWithGoroutine
)

// Result holds the Result Result.
type Result struct {
	Value string
}

// Capture is used to configure and manage the capturing of output streams like os.Stdout and os.Stderr.
type Capture struct {
	captureStdout bool
	captureStderr bool
	method        BufferMethod
}

// UseMethod initializes a new Capture instance with the specified BufferMethod.
func UseMethod(method BufferMethod) *Capture {
	capture := &Capture{method: method}
	return capture
}

// Stdout captures stdout.
func (c *Capture) Stdout(f func()) Result {
	c.captureStdout = true
	return Result{Value: c.capture(f)}
}

// Stderr captures stderr.
func (c *Capture) Stderr(f func()) Result {
	c.captureStderr = true
	return Result{Value: c.capture(f)}
}

// Output captures stdout and stderr.
func (c *Capture) Output(f func()) Result {
	c.captureStdout = true
	c.captureStderr = true
	return Result{Value: c.capture(f)}
}

// Stdout captures stdout.
func Stdout(f func()) Result {
	capture := &Capture{captureStdout: true}
	return Result{Value: capture.capture(f)}
}

// Stderr captures stderr.
func Stderr(f func()) Result {
	capture := &Capture{captureStderr: true}
	return Result{Value: capture.capture(f)}
}

// Output captures stdout and stderr.
func Output(f func()) Result {
	capture := &Capture{captureStdout: true, captureStderr: true}
	return Result{Value: capture.capture(f)}
}

func (c *Capture) capture(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defer r.Close()

	var buf bytes.Buffer

	if c.method == PipeWithGoroutine {
		// Use a goroutine to read data from the pipe.
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = io.Copy(&buf, r)
		}()

		c.redirectAndExecute(w, f)
		w.Close()
		wg.Wait() // Wait for the goroutine to finish reading.
	} else {
		// Use direct pipe reading (may block if buffer is full).
		c.redirectAndExecute(w, f)
		w.Close()
		_, err = io.Copy(&buf, r)
		if err != nil {
			panic(err)
		}
	}

	return buf.String()
}

func (c *Capture) redirectAndExecute(w *os.File, f func()) {
	if c.captureStdout {
		stdout := os.Stdout
		os.Stdout = w
		defer func() {
			os.Stdout = stdout
		}()
	}

	if c.captureStderr {
		stderr := os.Stderr
		os.Stderr = w
		defer func() {
			os.Stderr = stderr
		}()
	}

	f()
}

// AsBool converts the Result Result to a bool.
func (o Result) AsBool() (bool, error) {
	result, err := strconv.ParseBool(o.Value)
	return result, err
}

// AsString converts the Result Result to a string.
func (o Result) AsString() string {
	return o.Value
}

// AsInt converts the Result Result to an int.
func (o Result) AsInt() (int, error) {
	result, err := strconv.Atoi(o.Value)
	return result, err
}

// AsInt8 converts the Result Result to an int8.
func (o Result) AsInt8() (int8, error) {
	result, err := strconv.ParseInt(o.Value, 10, 8)
	return int8(result), err
}

// AsInt16 converts the Result Result to an int16.
func (o Result) AsInt16() (int16, error) {
	result, err := strconv.ParseInt(o.Value, 10, 16)
	return int16(result), err
}

// AsInt32 converts the Result Result to an int32.
func (o Result) AsInt32() (int32, error) {
	result, err := strconv.ParseInt(o.Value, 10, 32)
	return int32(result), err
}

// AsInt64 converts the Result Result to an int64.
func (o Result) AsInt64() (int64, error) {
	result, err := strconv.ParseInt(o.Value, 10, 64)
	return result, err
}

// AsUint converts the Result Result to an uint.
func (o Result) AsUint() (uint, error) {
	result, err := strconv.ParseUint(o.Value, 10, 64)
	return uint(result), err
}

// AsUint8 converts the Result Result to an uint8.
func (o Result) AsUint8() (uint8, error) {
	result, err := strconv.ParseUint(o.Value, 10, 8)
	return uint8(result), err
}

// AsUint16 converts the Result Result to an uint16.
func (o Result) AsUint16() (uint16, error) {
	result, err := strconv.ParseUint(o.Value, 10, 16)
	return uint16(result), err
}

// AsUint32 converts the Result Result to an uint32.
func (o Result) AsUint32() (uint32, error) {
	result, err := strconv.ParseUint(o.Value, 10, 32)
	return uint32(result), err
}

// AsUint64 converts the Result Result to an uint64.
func (o Result) AsUint64() (uint64, error) {
	result, err := strconv.ParseUint(o.Value, 10, 64)
	return result, err
}

// AsUintptr converts the Result Result to an uintptr.
func (o Result) AsUintptr() (uintptr, error) {
	result, err := strconv.ParseUint(o.Value, 10, 64)
	return uintptr(result), err
}

// AsByte converts the Result Result to an byte.
func (o Result) AsByte() (byte, error) {
	result, err := strconv.ParseUint(o.Value, 10, 8)
	return byte(result), err
}

// AsRune converts the Result Result to an rune.
func (o Result) AsRune() (rune, error) {
	result, err := strconv.ParseInt(o.Value, 10, 32)
	return rune(result), err
}

// AsFloat32 converts the Result Result to a float32.
func (o Result) AsFloat32() (float32, error) {
	result, err := strconv.ParseFloat(o.Value, 32)
	return float32(result), err
}

// AsFloat64 converts the Result Result to a float64.
func (o Result) AsFloat64() (float64, error) {
	result, err := strconv.ParseFloat(o.Value, 64)
	return result, err
}

// AsComplex64 converts the Result Result to a complex64.
func (o Result) AsComplex64() (complex64, error) {
	result, err := strconv.ParseComplex(o.Value, 64)
	return complex64(result), err
}

// AsComplex128 converts the Result Result to a complex128.
func (o Result) AsComplex128() (complex128, error) {
	result, err := strconv.ParseComplex(o.Value, 128)
	return result, err
}

func cleanInput(input string) string {
	// Remove square brackets and trim any surrounding whitespace.
	cleaned := strings.TrimSpace(input)
	cleaned = strings.Trim(cleaned, "[]")

	// Handle space-separated input by replacing spaces with commas.
	if strings.Contains(cleaned, " ") && !strings.Contains(cleaned, ",") {
		cleaned = strings.ReplaceAll(cleaned, " ", ",")
	}

	// Use a regex to clean up multiple commas or trailing commas.
	re := regexp.MustCompile(`,\s*,|,\s*$`)
	cleaned = re.ReplaceAllString(cleaned, ",")

	// Ensure the result is enclosed in square brackets.
	return "[" + cleaned + "]"
}

// AsSliceInt converts the Result Result to a slice of integers.
func (o Result) AsSliceInt() ([]int, error) {
	value := cleanInput(o.Value)

	var result []int
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceInt8 converts the Result Result to a slice of int8.
func (o Result) AsSliceInt8() ([]int8, error) {
	value := cleanInput(o.Value)

	var result []int8
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceInt16 converts the Result Result to a slice of int16.
func (o Result) AsSliceInt16() ([]int16, error) {
	value := cleanInput(o.Value)

	var result []int16
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceInt32 converts the Result Result to a slice of int32.
func (o Result) AsSliceInt32() ([]int32, error) {
	value := cleanInput(o.Value)

	var result []int32
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceInt64 converts the Result Result to a slice of int64.
func (o Result) AsSliceInt64() ([]int64, error) {
	value := cleanInput(o.Value)

	var result []int64
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceUint converts the Result Result to a slice of uint.
func (o Result) AsSliceUint() ([]uint, error) {
	value := cleanInput(o.Value)

	var result []uint
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceUint8 converts the Result Result to a slice of uint8.
func (o Result) AsSliceUint8() ([]uint8, error) {
	value := cleanInput(o.Value)

	var result []uint8
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceUint16 converts the Result Result to a slice of uint16.
func (o Result) AsSliceUint16() ([]uint16, error) {
	value := cleanInput(o.Value)

	var result []uint16
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceUint32 converts the Result Result to a slice of uint32.
func (o Result) AsSliceUint32() ([]uint32, error) {
	value := cleanInput(o.Value)

	var result []uint32
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceUint64 converts the Result Result to a slice of uint64.
func (o Result) AsSliceUint64() ([]uint64, error) {
	value := cleanInput(o.Value)

	var result []uint64
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceUintptr converts the Result Result to a slice of uintptr.
func (o Result) AsSliceUintptr() ([]uintptr, error) {
	value := cleanInput(o.Value)

	var result []uintptr
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceByte converts the Result Result to a slice of byte.
func (o Result) AsSliceByte() ([]byte, error) {
	value := cleanInput(o.Value)

	var result []byte
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceRune converts the Result Result to a slice of rune.
func (o Result) AsSliceRune() ([]rune, error) {
	value := cleanInput(o.Value)

	var result []rune
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceFloat32 converts the Result Result to a slice of float32.
func (o Result) AsSliceFloat32() ([]float32, error) {
	value := cleanInput(o.Value)

	var result []float32
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceFloat64 converts the Result Result to a slice of float64.
func (o Result) AsSliceFloat64() ([]float64, error) {
	value := cleanInput(o.Value)

	var result []float64
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceComplex64 converts the Result Result to a slice of complex64.
func (o Result) AsSliceComplex64() ([]complex64, error) {
	value := cleanInput(o.Value)

	var result []complex64
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceComplex128 converts the Result Result to a slice of complex128.
func (o Result) AsSliceComplex128() ([]complex128, error) {
	value := cleanInput(o.Value)

	var result []complex128
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceBool converts the Result Result to a slice of bools.
func (o Result) AsSliceBool() ([]bool, error) {
	value := cleanInput(o.Value)

	var result []bool
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}

// AsSliceString converts the Result Result to a slice of strings.
func (o Result) AsSliceString() ([]string, error) {
	value := cleanInput(o.Value)

	var result []string
	err := json.Unmarshal([]byte(value), &result)
	return result, err
}
