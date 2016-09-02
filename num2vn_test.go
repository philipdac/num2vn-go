package num2vn

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"testing"
	"strings"
	"strconv"
)

const (
	TEST_DATA_FILE = "num2vn_test.json"
)

type TestElement struct {
	Number  string  // Thousand separator can be space or comma
	VnStr   string  // Expected outcome in Vietnamese
}

type TestData struct {
	TestData []TestElement
}

func TestNum2Vn(t *testing.T) {

	// Read the json file
	rawData, err := ioutil.ReadFile(TEST_DATA_FILE)
	if err != nil {
		println("ERR! Reading file.", TEST_DATA_FILE, err.Error())
		return
	}

	// Decode the json content
	var testData TestData
	err = json.Unmarshal(rawData, &testData)
	if err != nil {
		println("ERR! Unmarshal resource.", err.Error())
		return
	}

	// Remove thousand separator " " or ","
	r := strings.NewReplacer(" ", "", ",", "")

	// Test and Error counter
	testCount := 0
	errCount := 0

	convStr := ""

	for _, test := range testData.TestData {

		numStr := r.Replace(test.Number)

		if strings.Index(numStr, ".") == -1 {
			// Convert an integer number
			numInt, err := strconv.ParseInt(numStr, 10, 64)

			if err != nil {
				fmt.Printf("ERR! Invalid number %s %s\n", test.Number, err.Error())
				continue
			}

			testCount++
			convStr = Int2Vn(numInt)

		} else {
			// Convert a float number
			numFloat, err := strconv.ParseFloat(numStr, 64)

			if err != nil {
				fmt.Printf("ERR! Invalid number %s %s\n", test.Number, err.Error())
				continue
			}

			testCount++
			convStr = Float2Vn(numFloat)
		}

		if strings.Compare(test.VnStr, convStr) != 0 {
			errCount++

			fmt.Printf("\n")
			fmt.Printf("Number:    %s\n", test.Number)
			fmt.Printf("Expected:  %s\n", test.VnStr)
			fmt.Printf("vs Output: %s\n", convStr)
		}
	}

	if errCount == 0 {
		fmt.Printf("\nSuccessfully tested %v %s\n", testCount, "number(s)")
	} else {
		fmt.Printf("\n%v %s\n", errCount, "error(s) found")
	}
}

func BenchmarkInt2Vn(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Int2Vn(9223372036854775807)
	}
}
