package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var text string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio."
	fmt.Println(text)

	var character = count(text)

	sortChar := make([]string, 0, len(character))
	for key := range character {
		sortChar = append(sortChar, key)
	}
	sort.Strings(sortChar)

	for _, key := range sortChar {
		fmt.Printf("Karakter %s sebanyak %d kali\n", key, character[key])
	}

	fmt.Println()
	moveChar(text)

}

func count(text string) map[string]int {
	text = strings.ToLower(text)
	var character = map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	for key := range character {
		character[key] = strings.Count(text, key)
	}
	return character
}

func moveChar(text string) {
	for _, char := range text {
		char += 5
		if char > 90 {
			if char <= 95 {
				char = char - 90 + 64
			} else if char <= 96 {
				char -= 5
			}
		}
		if char > 122 {
			char = char - 122 + 96
		}
		if char > 127 || char < 65 {
			char -= 5
		}
		fmt.Print(string(char))
	}
}
