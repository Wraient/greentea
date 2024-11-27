package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/exec"
	"math/rand"
	"time"
)

func searchAndCopy(searchTerm string, previousWords map[string]bool) bool {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		fmt.Println("Error opening dictionary:", err)
		return false
	}
	defer file.Close()

	// Collect all matching words
	var matches []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		word := strings.Fields(line)[0]
		
		if strings.Contains(strings.ToLower(word), strings.ToLower(searchTerm)) && !previousWords[word] {
			matches = append(matches, word)
		}
	}

	if len(matches) == 0 {
		fmt.Println("No new matching words found")
		return false
	}

	// Pick a random match
	randomIndex := rand.Intn(len(matches))
	selectedWord := matches[randomIndex]
	previousWords[selectedWord] = true

	// Copy to clipboard
	cmd := exec.Command("wl-copy", selectedWord)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		return false
	}
	fmt.Printf("Found and copied: %s (%d more matches available)\n", selectedWord, len(matches)-1)
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	previousWords := make(map[string]bool)
	
	for {
		fmt.Print("Enter substring (or 'quit' to exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "quit" {
			break
		}
		
		if input == "clear" {
			previousWords = make(map[string]bool)
			fmt.Println("History cleared!")
			continue
		}
		
		if input == "" {
			continue
		}
		
		searchAndCopy(input, previousWords)
	}
}