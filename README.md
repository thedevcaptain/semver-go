# SemVer

This is a library to check and compare two strings following the [Semantic Versioning](https://semver.org/) specification.

## Installation

```bash
go get github.com/thedevcaptain/semver-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/thedevcaptain/semver-go"
)

func main(){
    v1 := "1.0.0"
    v2 := "1.2.1"
    fmt.Println(semver.Compare(v1, v2)) // -1
    fmt.Println(semver.Compare(v2, v1)) // 1
    fmt.Println(semver.Compare(v2, v2)) // 0
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
