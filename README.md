## Library for working with sorted unidirectional linked list (integer data type)

## Install

`go get github.com/AnonymFromInternet/BSight/sortedList`

## Example

```go
import (
    "github.com/AnonymFromInternet/BSight/sortedList"
    "fmt"
    )

func main() {
    sl := sortedList.New()
    sl.Insert(5)

    res := sl.Delete(5)

    sl.Insert(8)
    isExists := sl.IsExists(8)

    node := sl.Search(8)
}
```