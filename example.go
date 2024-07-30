package internal_package_example

// example.go:3:8: use of internal package github.com/learning-go-book/internal_example/foo/internal not allowed
import "github.com/learning-go-book/internal_example/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}
