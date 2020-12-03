package main

import "fmt"

type treeMap struct {
	mapData [][]bool
}

func CountTrees() {
	data := ReadData("day_3.txt")
	treeMap := parseData(data)

	trees := countTreesWithSlope(treeMap, 1, 1)
	trees = trees * countTreesWithSlope(treeMap, 3, 1)
	trees = trees * countTreesWithSlope(treeMap, 5, 1)
	trees = trees * countTreesWithSlope(treeMap, 7, 1)
	trees = trees * countTreesWithSlope(treeMap, 1, 2)

	fmt.Printf("Multiplied together: %v\n", trees)
}

func countTreesWithSlope(data treeMap, right, down int) int {
	trees := 0

	j := 0
	for i := 0; i < len(data.mapData); i = i + down {
		if data.get(i, j) {
			trees++
		}
		j = j + right
	}

	fmt.Printf("Encountered Trees for %v, %v: %v\n", right, down, trees)

	return trees
}

func (m treeMap) get(i, j int) bool {
	return m.mapData[i][j % len(m.mapData[i])]
}

func parseData(input []string) treeMap {
	var mapData = make([][]bool, len(input))

	for i, line := range input {
		mapData[i] = make([]bool, len(input[i]))

		for j, tree := range line {
			if tree == '.' {
				mapData[i][j] = false
				continue
			}
			mapData[i][j] = true
		}
	}

	return treeMap{mapData: mapData}
}
