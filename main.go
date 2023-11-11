package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func aliasExists(alias, data string) bool {
	// Check if the alias already exists in the zshrc file
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "alias "+alias+"=") {
			return true
		}
	}
	return false
}

func updateZshrc(alias, value, zshrcPath string) error {
	// Read the content of the zshrc file
	data, err := os.ReadFile(zshrcPath)
	if err != nil {
		return fmt.Errorf("error reading %s: %v", zshrcPath, err)
	}

	// Check if the alias already exists in the zshrc file
	if aliasExists(alias, string(data)) {
		fmt.Printf("Alias '%s' already exists in %s\n", alias, zshrcPath)
		return nil
	}

	// Add the new alias to the zshrc file
	newContent := fmt.Sprintf("\n# Added by CLIAS CLI tool\nalias %s='%s'", alias, value)
	err = os.WriteFile(zshrcPath, append(data, []byte(newContent)...), os.ModeAppend)
	if err != nil {
		return fmt.Errorf("error updating %s: %v", zshrcPath, err)
	}

	fmt.Printf("Alias '%s' added to %s\n", alias, zshrcPath)
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get alias name from user input
	fmt.Print("Enter alias name: ")
	alias, _ := reader.ReadString('\n')
	alias = strings.TrimSpace(alias)

	// Get alias value from user input
	fmt.Print("Enter alias value: ")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	// Get zshrc file path from user input
	fmt.Print("Enter zshrc file path (or press Enter to use default): ")
	zshrcPath, _ := reader.ReadString('\n')
	zshrcPath = strings.TrimSpace(zshrcPath)

	// Use the default zshrc file path if not provided
	if zshrcPath == "" {
		zshrcPath = filepath.Join(os.Getenv("HOME"), ".zshrc")
	}

	// Update the zshrc file with the new alias
	err := updateZshrc(alias, value, zshrcPath)
	if err != nil {
		fmt.Println(err)
	}
}

