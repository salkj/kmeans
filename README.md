# k-means.go
a ready-to-use naive kmeans package for Go


```
package main

import (
	"fmt"
	"github.com/salkj/kmeans"
)

func main() {
	fmt.Println("Hello World!")
	data := []kmeans.Point{}
    data = append(data, kmeans.Point{[]float64{1.0,3.0,5.0,2.0}})
    data = append(data, kmeans.Point{[]float64{43.0,7.0,12.0,7.0}})
    data = append(data, kmeans.Point{[]float64{2.0,12.0,5.0,8.0}})
    data = append(data, kmeans.Point{[]float64{12.0,1945.0,34.0,65.0}})
    fmt.Println(kmeans.KMEANS(data, 2))
}
```
