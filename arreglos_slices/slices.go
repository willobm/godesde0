package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 7}
var arreglo [10]int = [10]int{2, 7, 3, 9, 6, 4, 8, 1, 0}

func MostrarSlices() {
	porcion1 := arreglo[3:]  //Slices creado desde un vector, desde la posición 3
	porcion2 := arreglo[:7]  //Slices creado desde un vectro, desde la posición 0 hasta la posición 7
	porcion3 := arreglo[6:7] //Slicec creado desde un vector, desde la posición 3 hasta la posición 7

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("largo %d, Capacidad %d \n", len(nums), cap(nums))
}
