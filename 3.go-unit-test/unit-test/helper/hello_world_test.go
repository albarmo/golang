package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"runtime"
)

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Albar")
	if result != "Hi Albar" {
		t.Fail()
	}
	fmt.Println("Test 1 Done!")
}

func TestHelloWorld2(t *testing.T){
	result := HelloWorld("Moerhamsa")
	if result != "Hello Moerhamsa" {
		t.FailNow()
	}
	// if using t.FailNow() program dibawahnya ga dieksekusi
	fmt.Println("Test 2 Done!")
}

func TestHelloWorld3(t *testing.T){
	result := HelloWorld("Moerhamsa")
	if result != "Hi Moerhamsa" {
		t.Error("Result Must be Hi Moerhamsa")
	}
	// if using t.FailNow() program dibawahnya ga dieksekusi
	fmt.Println("Test 3 Done!")
}

func TestHelloWorld4(t *testing.T){
	result := HelloWorld("Moerhamsa")
	if result != "Hi Moerhamsa" {
		t.Fatal("Result Must be Hi Moerhamsa")
	}
	// if using t.FailNow() program dibawahnya ga dieksekusi
	fmt.Println("Test 3 Done!")
}

func TestHelloWorldAssertion (t *testing.T){
	result:=  HelloWorld("Albar")
	assert.Equal(t,"Hi Albar", result, "Result must be 'Hi Albar'")
	fmt.Println("Executed! Test with assertion")
}

func TestHelloWorldAssertionNotEqual (t *testing.T){
	result:=  HelloWorld("Albar")
	// using assert => t.Fail()
	assert.NotEqual(t,"Hallo Albar", result, "Result must be 'Hi Albar'")
	fmt.Println("Executed! Test with assertion")
}

func TestHelloWorldRequire (t *testing.T){
	result := HelloWorld("Albar")
	// using require => t.FailNow()
	require.Equal(t,"Hi Albar", result, "Result must be 'Hi Albar'")
	fmt.Println("Test HelloWorld with require done!")
}

func TestSkip (t *testing.T){
	fmt.Println("Runtime : ", runtime.GOOS)
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak jalan di mac os!")
	}
	result := HelloWorld("Albar")
	require.Equal(t,"Hi Albar", result, "Result must be 'Hi Albar'")
}


// use m to managing test
func TestMain (m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")
	fmt.Println("1. Connecting to database...")
	m.Run()
	// after
	fmt.Println("2. Closing conncetion to database...")
	fmt.Println("AFTER UNIT TEST")
	fmt.Println("=== TESTING SUCCESS! ===")
}

func TestSubTest (t *testing.T){
	t.Run("Albar", func(t * testing.T){
		result := HelloWorld("Albar")
		require.Equal(t, "Hi Albar", result, "Result must be 'Hi Albar'")
	})
	t.Run("Moerhamsa", func(t * testing.T){
		result := HelloWorld("Moerhamsa")
		require.Equal(t, "Hi Moerhamsa", result, "Result must be 'Hi Moerhamsa'")
	})
	// to running specific sub-test => `go test -v -run=TestSubTest/Albar`
}

func TestHelloWorldTable (t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Test Table Iteration 1",
			request: "Albar",
			expected : "Hi Albar",
		},
		{
			name: "Test Table Iteration 2",
			request: "Moerhamsa",
			expected : "Hi Moerhamsa",
		},
		{
			name: "Test Table Iteration 3",
			request: "Joko",
			expected : "Hi Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name , func (t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}