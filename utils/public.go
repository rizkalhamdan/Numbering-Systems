package utils

import (
	"os"
	"os/exec"
)

func Code() int {
	return 94
}

func DO(com string) bool {
	cmd := exec.Command(com)

	cmd.Stdout = os.Stdout

	cmd.Run()

	return true
}

func Welcome() {
	println(" ## Number System ##")

	println("Good morning!")

	println()

	println("##################")
	println("##################")
}

func PrintTypes() {
	println()

	println("Numbering Systems")

	println("\t1. Binary")
	println("\t2. Octa")
	println("\t3. Decimal")
	println("\t4. Hexadecimal")

	println()
}




// ...existing code...
func GetRealNumber(num string) []int {

    var new_arr []int

    for _, r := range num {
        switch {
        case r >= '0' && r <= '9':
            new_arr = append(new_arr, int(r-'0'))
        case r >= 'a' && r <= 'f':
            new_arr = append(new_arr, int(r-'a'+10))
        case r >= 'A' && r <= 'F':
            new_arr = append(new_arr, int(r-'A'+10))
        default:
            // ignore non-hex characters
        }
    }

    return new_arr

}
// ...existing code...





// func GetRealNumber(num string) []int {

// 	num = strings.ToLower(num)
// 	arr := strings.Split(num, "")
// 	var new_arr []int

// 	for _, val := range arr {

// 		switch val {
// 		case "a":
// 			new_arr = append(new_arr, 10)
// 		case "b":
// 			new_arr = append(new_arr, 11)
// 		case "c":
// 			new_arr = append(new_arr, 12)
// 		case "d":
// 			new_arr = append(new_arr, 13)
// 		case "e":
// 			new_arr = append(new_arr, 14)
// 		case "f":
// 			new_arr = append(new_arr, 15)
// 		default:
// 			nve, err := strconv.Atoi(val)

// 			if err == nil {
// 				break
// 			}

// 			new_arr = append(new_arr, nve)
// 		}
// 	}

// 	return new_arr

// }
