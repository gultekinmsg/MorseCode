package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const BlankSpace = " "
var morseMap = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"ç": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"ğ":"--.",
	"h": "....",
	"ı": "..",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"ö": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"ş": "...",
	"t": "-",
	"u": "..-",
	"ü": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	" ": "/",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		input, err := reader.ReadString('\n') // read when press enter key
		if err != nil {
			log.Fatal(err)
		}
		input = strings.ReplaceAll(input, "\r\n", "") //
		input = strings.ReplaceAll(input, "\n", "")
		output := ConvertTo(input)
		fmt.Println(output)
	}
}

func ConvertTo(s string) string {
	s = strings.ToLower(s)
	var result strings.Builder
	didItExist := false
	for i := 0; i < len(s); i++ {
		for key, value := range morseMap {
			if key == string(s[i]) {
				result.WriteString(value)
				if i != len(s)-1 {
					result.WriteString(BlankSpace)
				}
				didItExist = true
				break
			}
		}
		if !didItExist {
			log.Printf("can not find %s at morse map", string(s[i]))
		}
	}
	return result.String()
}
