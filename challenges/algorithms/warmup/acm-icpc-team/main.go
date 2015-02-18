package main

import "fmt"

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

func knowsTopic(t uint8) int {
	if (t == 48) {
		return 0
	}
	return 1
}

func findMaxKnownTopics(people []string) (int, int) {
	var max_known, num_pairs, known_topics int

	max_known = 0
	num_pairs = 0

	for i := 0; i < len(people); i++ {
		for j := i + 1; j < len(people); j++ {
			topics1 := people[i]
			topics2 := people[j]
			known_topics = 0

			for k := 0; k < len(topics1); k++ {
				known_topics += knowsTopic(topics1[k]) | knowsTopic(topics2[k])
			}

			if known_topics > max_known {
				max_known = known_topics
				num_pairs = 1
			} else if known_topics == max_known {
				num_pairs++
			}
		}
	}

	return max_known, num_pairs
}

func main() {
	max_known, num_pairs := findMaxKnownTopics(readPeople())

	fmt.Println(max_known)
	fmt.Println(num_pairs)
}