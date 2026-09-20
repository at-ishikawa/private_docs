package internal

type node struct {
	children map[byte]*node
	isEnd    bool
}

var (
	directions = [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	foundWords map[string]bool
	visited    [][]bool
)

func (n node) searchBoard(
	board [][]byte,
	row int,
	column int,
	currentWord string,
) {
	if row < 0 || len(board) <= row || column < 0 || len(board[0]) <= column {
		return
	}
	if visited[row][column] {
		return
	}

	child, ok := n.children[board[row][column]]
	if !ok {
		return
	}

	visited[row][column] = true
	nextWord := currentWord + string(board[row][column])

	if child.isEnd {
		foundWords[nextWord] = true
	}
	for _, direction := range directions {
		child.searchBoard(board, row+direction[0], column+direction[1], nextWord)
	}
	visited[row][column] = false
}

func findWords(board [][]byte, words []string) []string {
	root := &node{
		children: make(map[byte]*node),
	}
	for _, word := range words {
		current := root
		for _, letter := range []byte(word) {
			next, ok := current.children[letter]
			if !ok {
				next = &node{
					children: make(map[byte]*node),
				}
				current.children[letter] = next
			}
			current = next
		}
		current.isEnd = true
	}
	// fmt.Printf("%#v, %#v\n", root, root.children['a'])

	foundWords = make(map[string]bool)
	visited = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

LOOP:
	for rowIndex, row := range board {
		for columnIndex := range row {
			root.searchBoard(
				board,
				rowIndex,
				columnIndex,
				"",
			)
			if len(foundWords) == len(words) {
				break LOOP
			}
		}
	}

	result := make([]string, 0, len(foundWords))
	for word := range foundWords {
		result = append(result, word)
	}
	return result
}
