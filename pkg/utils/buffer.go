package utils

import (
	"bufio"
	"os"
)

// Helper function to clear the input buffer in case of read errors
func ClearBuffer() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
