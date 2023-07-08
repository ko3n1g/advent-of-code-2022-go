package marker

import (
	"errors"
	"strings"
)

func GetIdxOfMarker(message string, messageSize int) (int, error) {
	if len(message) == 0 {
		return 0, errors.New("Message empty, not happy with that.")
	}

	for i := messageSize; i < len(message); i++ {
		var uniqueChars map[string]bool = make(map[string]bool)
		for _, char := range strings.Split(message[i-messageSize:i], "") {
			uniqueChars[char] = true
		}

		if len(uniqueChars) == messageSize {
			return i, nil
		}
	}

	return 0, errors.New("No marker found in message.")
}
