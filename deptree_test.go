package deptree

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("test parent child", t, func() {
		d := NewTree()
		d.Add("grandparent", "parent")
		d.Add("parent", "me")
		fmt.Println(d.Solve("me"))
		fmt.Println("**")

		e := NewTree()
		e.Add("grandparent", "parent1")
		e.Add("grandparent", "parent2")
		e.Add("parent1", "me")
		e.Add("parent2", "me")
		fmt.Println(e.Solve("parent1"))
		fmt.Println(e.Solve("parent2"))
		fmt.Println(e.Solve("me"))
	})
}
