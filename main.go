package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Program for finding the longest sequence of fragments.
The algorithm works as follows:
1. Reads fragments from file
2. Builds a directed graph where vertices are fragments and edges connect fragments
   if the last two characters of one match the first two characters of another
3. Finds the longest path in the graph using a recursive algorithm
4. Outputs the result
*/

// Node represents a single fragment in the graph
type Node struct {
	Value string
	Edges []*Node
}

// FindLongestPath recursively finds the longest path in the graph
func FindLongestPath(node *Node, visited map[string]bool, currentPath []string) []string {
	visited[node.Value] = true
	currentPath = append(currentPath, node.Value)

	longestPath := make([]string, len(currentPath))
	copy(longestPath, currentPath)

	for _, neighbor := range node.Edges {
		if !visited[neighbor.Value] {
			path := FindLongestPath(neighbor, visited, currentPath)
			if len(path) > len(longestPath) {
				longestPath = path
			}
		}
	}

	visited[node.Value] = false
	return longestPath
}

func main() {
	// Step 1: Read data from the file
	file, err := os.Open("fragments.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var fragments []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			fragments = append(fragments, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 2: Build graph
	nodes := make(map[string]*Node)
	for _, fragment := range fragments {
		nodes[fragment] = &Node{Value: fragment}
	}

	for _, fragment := range fragments {
		for _, otherFragment := range fragments {
			if fragment != otherFragment && strings.HasSuffix(fragment, otherFragment[:2]) {
				nodes[fragment].Edges = append(nodes[fragment].Edges, nodes[otherFragment])
			}
		}
	}

	// Step 3: Find the longest path
	longestPath := []string{}
	for _, node := range nodes {
		visited := make(map[string]bool)
		path := FindLongestPath(node, visited, []string{})
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}

	// Step 4: Output the result
	fmt.Printf("Longest sequence: %s\n", strings.Join(longestPath, ""))
	fmt.Printf("Fragments used: %d out of %d\n", len(longestPath), len(fragments))
}
