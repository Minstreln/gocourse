package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

func main() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}

	// sort.Sort(ByAge(people))
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted by age", people)

	// numbers := []int{5, 3, 4, 1, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// stringSlice := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted strings:", stringSlice)

}
