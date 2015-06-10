package main

import (
	"fmt"
	"bufio"
	"os"
)

// quick stack for DFS
type Stack []int

func (s *Stack) Push(i int) {
	*s = append(*s, i);
}

// -1 for empty stack because it's not a valid value
func (s *Stack) Pop() int {
	len := s.Len();
	if len == 0 {
		return -1;
	}
	r := (*s)[len - 1];
	*s = (*s)[:len - 1];
	return r;
}

func (s *Stack) Len() int {
	return len(*s);
}

func InitializeAdjacencyMatrix(vertices int) [][]bool {
	// vertices + 1 to avoid all the x-1 notations from input values
	matrix := make([][]bool, vertices + 1);
	for i := range matrix {
		matrix[i] = make([]bool, vertices + 1);
	}
	return matrix;
}

func ReadTreeDimensions(scanner *bufio.Reader) (int, int) {
	var vertices, edges int;
	fmt.Fscanf(scanner, "%d %d\n", &vertices, &edges);
	return vertices, edges;
}

func ReadVertices(scanner *bufio.Reader, edges int, matrix [][]bool) [][]bool {
	var x, y int;
	for i := 0; i < edges; i++ {
		fmt.Fscanf(scanner, "%d %d\n", &x, &y);
		matrix[x][y] = true;
		matrix[y][x] = true;
	}
	return matrix;
}

func DoProblem(matrix [][]bool, vertex int) int {
	size := len(matrix[0]);
	visited := make([]bool, size);
	children := male([]int, size);
	traverse := Stack{};
	//remove := 0;

	traverse.Push(vertex);

	for traverse.Len() > 0 {
		current := traverse.Pop();
		fmt.Println(current);
		if !visited[current] {
			visited[current] = true;
			for edge, exists := range matrix[current] {
				if exists && !visited[edge] {
					traverse.Push(edge);
				}
			}
		}
	}

	return 1;
}

// for debugging
func PrintMatrix(matrix [][]bool) {
	for i := range matrix {
		s := ""
		for j := range matrix[i] {
			if matrix[i][j] {
				s += "1 ";
			} else {
				s += "0 ";
			}
		}
		fmt.Println(s);
	}
}

func main() {
	scanner := bufio.NewReader(os.Stdin);

	vertices, edges := ReadTreeDimensions(scanner);
	adjacency := InitializeAdjacencyMatrix(vertices);
	adjacency = ReadVertices(scanner, edges, adjacency);

	PrintMatrix(adjacency);
	DoProblem(adjacency, 1);
}