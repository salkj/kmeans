/*
* 	Jayson Salkey
*	01/26/2016 02:54 UTC-5
*/ 

package kmeans

import (
	"math"
	"math/rand"
)

const DELTA = 0.001 
 
type point struct{
    entry[] float64
}
 
type centroid struct{
    center point
    points []point
}
 
func (p_1 point) distanceTo(p_2 point) float64{
	sum := float64(0)
	for e:=0;e<len(p_1.entry);e++{
		sum += math.Pow((p_1.entry[e] - p_2.entry[e]),2)
	}
    return math.Sqrt(sum)
}
 
func (c_1 *centroid) reCenter() float64{
    new_centroid := make([]float64,len(c_1.center.entry))
    for _, e := range c_1.points{
    	for r:=0;r<len(new_centroid);r++{
    		new_centroid[r] += e.entry[r]
    	}
    }
    for r:=0;r<len(new_centroid);r++{
    	new_centroid[r] /= float64(len(c_1.points))
    }
    old_center := c_1.center
    c_1.center = point{new_centroid}
    return old_center.distanceTo(c_1.center)
}
 
func kmeans(data []point, k uint64) (centroids []centroid){
    for i:=uint64(0); i<k; i++{
        centroids = append(centroids,centroid{center: data[rand.Intn(len(data))]})
    }
    
    converged := false
    for !converged {
        for i:= range data{
            min_distance := 9999999.0
            z := 0
            for v,e:=range centroids {
                distance := data[i].distanceTo(e.center)
                if distance < min_distance{
                    min_distance = distance;
                    z = v
                }
            }
            centroids[z].points = append(centroids[z].points, data[i])
        }
        max_delta := -9999999.0
        for i:= range centroids {
            movement := centroids[i].reCenter()
            if movement > max_delta{
                max_delta = movement
            }
        }
        if DELTA >= max_delta{
            converged = true
            return
        }
        for i:= range centroids {
            centroids[i].points = nil
        }
    }
    return centroids
}