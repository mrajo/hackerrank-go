package main

import (
	"fmt"
	"strconv"
)

func readPeople() []string {
	var num_people, topics int
	var known_topics string
	var people []string

	fmt.Scanf("%v %v", &num_people, &topics)

	for i := 0; i < num_people; i++ {
		fmt.Scan(&known_topics)
		people = append(people, known_topics)
	}

	return people
}

func hammingWeight(i int64) int64 {
    i = i - ((i >> 1) & 0x55555555)
    i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
    return (((i + (i >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24
}

func findMaxKnownTopics(people []string) (int64, int64) {
	var hw, max_hamming_weight, pairs int64

	max_hamming_weight = 0
	pairs = 0

	for i := 0; i < len(people); i++ {
		for j := i + 1; j < len(people); j++ {
			known_topics1, _ := strconv.ParseInt(people[i], 2, 0)
			known_topics2, _ := strconv.ParseInt(people[j], 2, 0)
			hw = hammingWeight(known_topics1 | known_topics2)

			if hw > max_hamming_weight {
				max_hamming_weight = hw
				pairs = 1
			} else if hw == max_hamming_weight {
				pairs++
			}
		}
	}

	return max_hamming_weight, pairs
}

func main() {
	max_known, num_pairs := findMaxKnownTopics(readPeople())

	fmt.Println(max_known)
	fmt.Println(num_pairs)
}