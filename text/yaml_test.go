package main

import (
	"fmt"
	"testing"
)


func TestYaml(t *testing.T)   {

	map1 := make( map[string]map[string]string)

	map0:=make( map[string]string)
	map0["1"]="1a"
	map0["2"]="2a"
	map0["3"]="3a"
	map1["A"]= map0
	fmt.Println(

		map1,
		)

}
