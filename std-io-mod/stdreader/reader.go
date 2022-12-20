// Package stdreader provides functionality for reading from the standard input.
package stdreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetDataFromUser(promptInformation string, delimiter byte) (string, error) {
	fmt.Print(promptInformation)
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString(delimiter)
	if err != nil {
		return "", err
	}

	data = strings.TrimSpace(data)

	return data, nil
}
