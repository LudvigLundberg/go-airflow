# go-airflow
> A small go wrapper for Apache Airflow commands

## Structure
    ├── cli                     # cli wrapper library 
    ├── main.go                 
    └── README.md

## Build
A simple executable can be built by running

```sh
go build
```

## Usage example

Using the library:

```go
import (
	"fmt"
  
	"github.com/LudvigLundberg/go-airflow/cli"
)

func listDags() {
	dags, err := cli.ListDags()

	if err != nil {
		panic(err)
	}

	for _, dag := range dags {
		fmt.Println(dag)
	}
}
```


Running the executeable (requires Airflow):
```sh
./go-airflow list_dags
```


## Release History

* 0.0.1
    * Work in progress - list_dags implemented

## Contributing

1. Fork it (<https://github.com/yourname/yourproject/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
