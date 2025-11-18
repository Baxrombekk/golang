package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click

// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	massiv := []int{2, 1, 1, 15, 4, 8, 7, 6, 5}
	for i := 0; i < len(massiv)-1; i++ {
		for j := 0; j < len(massiv)-1-i; j++ {
			if massiv[j] > massiv[j+1] {
				massiv[j], massiv[j+1] = massiv[j+1], massiv[j]
			}
		}
	}

	fmt.Println(massiv)

}
