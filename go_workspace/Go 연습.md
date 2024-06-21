#go 

# 기초

```go title:메서드연습_1
package main

import (

"fmt"

"time"

)

func main() {

var now time.Time = time.Now()

var year int = now.Year()

fmt.Println(year)

}
```
```go title:메서드연습_2
package main

  

import (

"fmt"

"strings"

)

  

func main() {

broken := "G# r#cks!"

replacer := strings.NewReplacer("#", "o")

fixed := replacer.Replace(broken)

fmt.Println(fixed)

}
```
