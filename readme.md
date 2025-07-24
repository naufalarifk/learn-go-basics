
# Unit Test, Benchmark & Mock

this section here will explain based on my current understanding of how unit tests in GO-Lang work.



## Benchmark Usage and example


this is an example of unit testing in GO-Lang, common usage is with benchmark table, but the code contains all testing type

```go
package helper

import (
	"fmt"
	"runtime"
	"testing"

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

```

benchmark has to start with the keyword 'Benchmark' at the start of the func for example like the above code it starts with Benchmark, thus the name BenchmarkTable.

'Benchmark' + your custom additional func name

benchmark table iterates the benchmark struct like the example above, for loop in the range of benchmark. use b.Run("name here", func here). b.N is the default number of the benchmarks provided by the package itself.


how to run benchmark

```go test -v -run=NoTest -bench=BenchmarkTable```

-run=NoTest is when you dont want to run any tests, -bench as for which benchmark func you want to run




you can also run sub benchmark, same naming rules. only difference here is that you have multiple benchmarks to run on a single benchmark func.

```go

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


```

how to run sub-benchmark 

``` go test -v -run=NoTest -bench=BenchmarkSub/Fal```

this will run the second b.Run()




Benchmark [API Reference](https://pkg.go.dev/testing#hdr-Benchmarks)



