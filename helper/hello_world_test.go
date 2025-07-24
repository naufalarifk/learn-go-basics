package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Nau",
			request: "Nau",
		},
		{
			name:    "Arif",
			request: "Arif Kurniawan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkSub(b *testing.B) {
	b.Run("Nau", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nau")
		}
	})
	b.Run("Fal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fal")
		}
	})
}

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nau")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{{
		name:     "Nau",
		request:  "Nau",
		expected: "Hello Nau",
	},
		{
			name:     "Fal",
			request:  "Fal",
			expected: "Hello Fal",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := HelloWorld(test.request)
			require.Equal(t, test.expected, res)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Nau", func(t *testing.T) {
		res := HelloWorld(("Nau"))
		require.Equal(t, "Hello Nau", res, "res Must be 'Hello Nau'")
	})
	t.Run("Fal", func(t *testing.T) {
		res := HelloWorld(("Fal"))
		require.Equal(t, "Hello Fal", res, "res Must be 'Hello Fal'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Running Tests...")
	m.Run()
	fmt.Println("All Tests Done!")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nau")
	if result != "Hello Nau" {
		t.Error("Should BE Nau")
	}
}

func TestHelloWorldZamn(t *testing.T) {
	res := HelloWorld("zamn")
	if res != "Hello zamn" {
		t.Fatal("zamn boi")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	res := HelloWorld("Require")
	require.Equal(t, "Hello Require", res, "Res has to be Hello Assert")
}

func TestHelloWorldAssertion(t *testing.T) {
	res := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", res, "Res has to be Hello Assert")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}

}
