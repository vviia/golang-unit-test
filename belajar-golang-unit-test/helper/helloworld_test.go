package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"  //saat gagal akan memanggil fail
	"github.com/stretchr/testify/require" // saat gagal akan memanggil fail now
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Via",
			request: "Via",
		},
		{
			name:    "Octavia",
			request: "Octavia",
		},
		{
			name:    "Aulia Octavia",
			request: "Aulia Octavia vea",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
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
	b.Run("via", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("via")
		}
	})
	b.Run("octavia", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("octavia")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("via")
	}
}

func BenchmarkHelloWorldoctavia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("octavia")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "via",
			request:  "via",
			expected: "Hello via",
		},
		{
			name:     "octavia",
			request:  "octavia",
			expected: "Hello octavia",
		},
		{
			name:     "vea",
			request:  "vea",
			expected: "Hello vea",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("via", func(t *testing.T) {
		result := HelloWorld("via")
		require.Equal(t, "Hello via", result, "Result must be 'Hello via'")
	})
	t.Run("octavia", func(t *testing.T) {
		result := HelloWorld("octavia")
		require.Equal(t, "Hello octavia", result, "Result must be 'Hello octavia'")
	})
	t.Run("vea", func(t *testing.T) {
		result := HelloWorld("vea")
		require.Equal(t, "Hello vea", result, "Result must be 'Hello vea'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("via")
	require.Equal(t, "Hello via", result, "Result must be 'Hello via'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("via")
	require.Equal(t, "Hello via", result, "Result must be 'Hello via'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("via")
	assert.Equal(t, "Hello via", result, "Result must be 'Hello via'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldVia(t *testing.T) {
	result := HelloWorld("via")

	if result != "Hello via" {
		// error
		t.Error("Result must be 'Hello via'")
	}

	fmt.Println("TestHelloWorldVia Done")
}

func TestHelloWorldVea(t *testing.T) {
	result := HelloWorld("vea")

	if result != "Hello vea" {
		// error
		t.Fatal("Result must be 'Hello vea'")
	}

	fmt.Println("TestHelloWorldVea Done")
}

func TestHelloWorldoctavia(t *testing.T) {
	result := HelloWorld("octavia")

	if result != "Hello octavia" {
		// error
		t.Fatal("Result must be 'Hello octavia'")
	}

	fmt.Println("TestHelloWorldoctavia Done")
}
