# go-pypi

A PyPi API client enabling Go project to interact with PyPi.

## Usage

### Installation

```console
go get github.com/roaldnefs/go-pypi
```

### Importing

```go
import "github.com/roaldnefs/go-pypi"
```

## Examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/roaldnefs/go-pypi"
)

func main() {
	pypi := pypi.NewClient(nil)

	// Get a project
	project, _, err := pypi.Project.GetProject("sampleproject")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(project.Info.Name)

	// Get a project
	release, _, err := pypi.Project.GetRelease("sampleproject", "1.3.0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(release.Info.Name)
}
```