//Package keyboard reads user input from the keyboard.
//Doc comments should begin with “Package” followed by the package name.

package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Function comments should begin with the name of the function they describe:
// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
