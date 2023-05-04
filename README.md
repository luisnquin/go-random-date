# randatetime - A package to generate date-time objects between two years

## Install

```bash
go get github.com/luisnquin/randatetime
```

## Use

```go
package main

import (
 "time"
 "github.com/luisnquin/randatetime"
)
 
func main() { 
    dt := randatetime.BetweenYears(2015, 2020)

    //...
}
```
