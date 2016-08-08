package main

import (
	"fmt"
	"github.com/dhconnelly/rtreego"
	"strconv"
)

func main() {
	myExample2()
}

func example() {
	// rt := rtreego.NewTree(2, 25, 50)
	// bb, _ := rtreego.NewRect(rtreego.Point{1.7, -3.4}, []float64{3.2, 1.9})

	// // Get a slice of the objects in rt that intersect bb:
	// // results := rt.SearchIntersect(bb)
	// // fmt.Println(results)

	// // Get a slice of the objects in rt that are contained inside bb:
	// // results = rt.SearchContained(bb)

	// q := rtreego.Point{6.5, -2.47}
	// k := 5

	// // Get a slice of the k objects in rt closest to q:
	// rt.SearchIntersectWithLimit(k, bb)
	// results := rt.SearchNearestNeighbors(q, k)
	// fmt.Println(results)

	rt := rtreego.NewTree(2, 3, 3)

	thing1 := mustRect(rtreego.Point{1, 1}, []float64{1, 1})
	thing2 := mustRect(rtreego.Point{1, 3}, []float64{1, 1})
	thing3 := mustRect(rtreego.Point{3, 2}, []float64{1, 1})
	thing4 := mustRect(rtreego.Point{-7, -7}, []float64{1, 1})
	thing5 := mustRect(rtreego.Point{7, 7}, []float64{1, 1})
	thing6 := mustRect(rtreego.Point{10, 2}, []float64{1, 1})

	// 	rt.Insert(&Thing{r1, "foo"})
	// 	rt.Insert(&Thing{r2, "bar"})
	rt.Insert(&Thing{thing1, "thing1"})
	rt.Insert(&Thing{thing2, "thing2"})
	rt.Insert(&Thing{thing3, "thing3"})
	rt.Insert(&Thing{thing4, "thing4"})
	rt.Insert(&Thing{thing5, "thing5"})
	rt.Insert(&Thing{thing6, "thing6"})

	obj1 := rt.NearestNeighbor(rtreego.Point{0.5, 0.5})
	// obj2 := rt.NearestNeighbor(rtreego.Point{1.5, 4.5})
	// obj3 := rt.NearestNeighbor(rtreego.Point{5, 2.5})
	// obj4 := rt.NearestNeighbor(rtreego.Point{3.5, 2.5})

	fmt.Println(obj1)

	// if obj1 != things[0] || obj2 != things[1] || obj3 != things[2] || obj4 != things[2] {
	// 	t.Errorf("NearestNeighbor failed")
	// }
}
func mustRect(p rtreego.Point, widths []float64) *rtreego.Rect {
	r, err := rtreego.NewRect(p, widths)
	if err != nil {
		panic(err)
	}
	return r
}

//北京经度范围  115.862734 - 116.855577
//北京纬度范围  40.107910 - 39.754257
func myExample() {
	rt := rtreego.NewTree(2, 3, 3)

	thing1 := mustRect(rtreego.Point{115.862734, 40.10791}, []float64{1, 1})
	thing2 := mustRect(rtreego.Point{115.862734, 39.754257}, []float64{1, 1})
	thing3 := mustRect(rtreego.Point{116.855577, 40.10791}, []float64{1, 1})
	thing4 := mustRect(rtreego.Point{116.855577, 39.754257}, []float64{1, 1})
	thing5 := mustRect(rtreego.Point{116.249042, 39.988167}, []float64{1, 1})
	thing6 := mustRect(rtreego.Point{116.290435, 39.898774}, []float64{1, 1})

	rt.Insert(&Thing{thing1, "thing1"})
	rt.Insert(&Thing{thing2, "thing2"})
	rt.Insert(&Thing{thing3, "thing3"})
	rt.Insert(&Thing{thing4, "thing4"})
	rt.Insert(&Thing{thing5, "thing5"})
	rt.Insert(&Thing{thing6, "thing6"})

	obj1 := rt.NearestNeighbor(rtreego.Point{116.761866, 40.086717})
	// obj2 := rt.NearestNeighbor(rtreego.Point{1.5, 4.5})
	// obj3 := rt.NearestNeighbor(rtreego.Point{5, 2.5})
	// obj4 := rt.NearestNeighbor(rtreego.Point{3.5, 2.5})

	fmt.Println(obj1)
}

func myExample2() {
	rt := rtreego.NewTree(2, 3, 3)

	for x := 0.0; x < 10.0; x++ {
		for y := 0.0; y < 10.0; y++ {
			thing1 := mustRect(rtreego.Point{x, y}, []float64{1, 1})
			rt.Insert(&Thing{thing1, strconv.FormatFloat(x, 'f', 6, 64) + strconv.FormatFloat(y, 'f', 6, 64)})
		}
	}

	obj1 := rt.NearestNeighbor(rtreego.Point{5.9, 6.0})

	fmt.Println(obj1)
}

// func example() {
// 	rt := rtreego.NewTree(2, 25, 50)

// 	p1 := rtreego.Point{0.4, 0.5}
// 	p2 := rtreego.Point{6.2, -3.4}

// 	r1, _ := rtreego.NewRect(p1, []float64{1, 2})
// 	r2, _ := rtreego.NewRect(p2, []float64{1.7, 2.7})

// 	rt.Insert(&Thing{r1, "foo"})
// 	rt.Insert(&Thing{r2, "bar"})

// 	size := rt.Size() // returns 2

// 	rt.Delete(thing2)
// 	// do some stuff...
// 	rt.Insert(anotherThing)
// }

type Thing struct {
	where *rtreego.Rect
	name  string
}

func (t *Thing) Bounds() *rtreego.Rect {
	return t.where
}
