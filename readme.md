
## Mock usage and example

this is an example of Mock testing in GO-Lang, common usage is to simulate querying data(?)

but anyway, this is how it works,


first you might wanna make an entity folder (or you can check in the repo /entity). make a category.go (here in my case, since we are querying category)

```go
package entity

type Category struct {
	Id   string
	Name string
}
```

then you can make a folder /repository or whatever based on your current case.

then you need to make it look like this:

```go

package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entity.Category)
		return &category
	}

}


```

as you can see here i am importing, my entity folder.

```go
import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

```

then init the struct and the func

```go

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entity.Category)
		return &category
	}

}

```

the CategoryRepositoryMock contains Mock field which contains the Mock struct from the package. 


then we bind the func FindById to the CategoryRepositoryMock.

```go

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
}

```

here to enable us to call, 

```go

args := repository.Mock.Called(id)

```

then, if index[0] we get equals to nil we return nil, else the method returns category. 

why index[0]? because we just want to simulate.

```go

		category := args.Get(0).(entity.Category)
		return &category

```

this method will return a pointer to *entity.Category.


after this you will want to create a /repository folder containing category_repository.go and category_repository_mock.go

we edit category_repository.go first, ignore category_repository_mock.go for now.


```go

package repository

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}


```



then you will want to create a /service folder containing category_service.go and category_service_test.go.


in category_service.go it will look like this,


```go

package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}


```


in category_service.go you will see this code:

```go

package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}


```


we import several packages from our own codebase like entity and repository.

first init, the struct:

```go

type CategoryService struct {
	Repository repository.CategoryRepository
}

```




while in category_service_test.go, you will find code like this.


```go

package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	cat := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(cat)

	res, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, cat.Id, res.Id)
	assert.Equal(t, cat.Id, res.Id)
	assert.Equal(t, cat.Name, res.Name)

}


```

declare these two, make reference to CategoryRepositoryMock from the folder repository.
then also declare categoryService

```go

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

```

here we have both cases where it is not found and success, 


```go

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

```

see how we declare

```go

	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")


```

while categoryRepository needs to run the mock on FindById, which declared previously and now we insert args "1". here we simulate what happens if we can't find the item.


while, 

```go

func TestCategoryService_GetSuccess(t *testing.T) {
	cat := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(cat)

	res, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, cat.Id, res.Id)
	assert.Equal(t, cat.Id, res.Id)
	assert.Equal(t, cat.Name, res.Name)

}

```

does the otherwise.


## Benchmark usage and example

this is an example of benchmark testing in GO-Lang, common usage is with benchmark table, but the code contains all testing type

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

```$ go test -v -run=NoTest -bench=BenchmarkSub/Fal```

this will run the second 

```go 

	b.Run("Fal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fal")
		}
	})

```








Benchmark [API Reference](https://pkg.go.dev/testing#hdr-Benchmarks)



