package funcfilter

import "fmt"

type User struct {
	age  int
	name string
}

func RunFilterDemo() {
	fmt.Println("------Filter Users by Age Demo------")

	allUsers := []User{
		User{
			age:  20,
			name: "User20",
		},
		User{
			age:  40,
			name: "User40",
		},
		User{
			age:  60,
			name: "User60",
		},
		User{
			age:  80,
			name: "User80",
		},
	}
	const MIN_AGE = 50
	filteredUsers := Filter[User](allUsers, func(u User) bool { return u.age > MIN_AGE })
	fmt.Println("all users     :", allUsers)
	fmt.Println("filtered users: ", filteredUsers)
	fmt.Println("")
}

// Filter returns a new slice of elements of Type T for which the filterFunc evaluates to True
func Filter[T any](unfiltered []T, filterFunc func(T) bool) []T {
	filtered := make([]T, 0)
	for _, e := range unfiltered {
		if filterFunc(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
