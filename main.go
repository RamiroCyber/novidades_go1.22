package main

import (
	"fmt"
	"go/version"
	"slices"
)

func main() {

	// Comparando versão e checando se é valida
	old := "go1.19"
	new := "go1.22"

	for _, v := range []string{old, new} {
		if version.IsValid(v) {
			fmt.Printf("%s is valid\n", v)
		} else {
			fmt.Printf("%s is invalid\n", v)

		}
	}

	switch version.Compare(old, new) {
	case 0:
		fmt.Printf("%s and %s are equal\n", old, new)
	case 1:
		fmt.Printf("%s is greater than %s\n", old, new)
	case -1:
		fmt.Printf("%s is less than %s\n", old, new)

	}

	//Diferente da versão inferior agora o comprimido apenas com as posições restantes
	sliceA := []uint{1, 5, 10}
	sliceB := []uint{4, 20, 16}

	join := slices.Concat(sliceA, sliceB)
	fmt.Printf("%v joined slice\n", join)

	join = slices.Delete(join, 0, 2)
	fmt.Printf("%v deleted slice\n", join)

}
