/*
* 	Jayson Salkey
*	01/26/2016 02:54 UTC-5
*/ 

package kmeans

import (
	"fmt"
)

func main(){
	data := []point{}
	data = append(data, point{[]float64{1.0,3.0,5.0,2.0}})
	data = append(data, point{[]float64{43.0,7.0,12.0,7.0}})
	data = append(data, point{[]float64{2.0,12.0,5.0,8.0}})
	data = append(data, point{[]float64{12.0,1945.0,34.0,65.0}})
	fmt.Println(kmeans(data, 2))
	
}
