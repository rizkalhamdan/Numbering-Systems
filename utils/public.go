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
