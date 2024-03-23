package algohash

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	source := "中文的也行吗"
	target := "中文的也行吗"
	fmt.Println(validAnagram(source, target))
	s1 := "中文English"
	t1 := "中文English"
	fmt.Println(validAnagram(s1, t1))
}
