package main

import (
	// "ft"
	"os"
	"piscine"
	"fmt"
	"log"
	"strings"
	"strconv"
	"time"
)

import "math/rand"

func isNumericalRune(c rune) bool {
	if	'0' <= c && c <= '9' {
		return true
	}
	return false
}

func isNumeric(s string) bool {
	ret := []rune(s)
	for _, r := range ret {
		if isNumericalRune(r) == false {
			return false
		}
	}
	return true
}

func testDirectory() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list,_ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fmt.Println("-----------------")
		fmt.Println(name)
		fmt.Println("-----------------")
		file := string(os.Args[1]) + "/" + string(name)
		pieces, err := piscine.ConvertFile(file)
		if err != nil {
			piscine.Puts(err.Error())
			return
		}
		if len(pieces) > 26 {
			piscine.Puts("Impossible")
			return
		}
		piscine.Solve(pieces)
		fmt.Println("-----------------")
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//とりあえず試行回数は一回で、何回やるかもとるか
func randomGeneration() {
	fmt.Printf("random generation : %s elements\n", os.Args[1])
	fmt.Println("-----------------")
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	j := 10
	if len(os.Args) == 3 {
		j, err = strconv.Atoi(os.Args[2])
	}
	for x := 0; x < j; x++ {
		// ここで引数分ランダム駒を作っていき、solveにかける
		pieces := []piscine.Piece{}
		// pieces = nil
	for j := i; j > 0; j-- {
		rand.Seed(time.Now().UnixNano())
		form := randomInt(1, 19)
		if form == 1{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T, piscine.T, piscine.T}})
		}
		if form == 2{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T, piscine.T}, {piscine.T, piscine.F, piscine.F}})
		}
		if form == 3{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T, piscine.T}, {piscine.F, piscine.T, piscine.F}})
		}
		if form == 4{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T, piscine.T}, {piscine.F, piscine.F, piscine.T}})
		}
		if form == 5{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T}, {piscine.T, piscine.T}})
		}
		if form == 6{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T, piscine.F}, {piscine.F, piscine.T, piscine.T}})
		}
		if form == 7{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T}, {piscine.T, piscine.F}, {piscine.T, piscine.F}})
		}
		if form == 8{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.T}, {piscine.F, piscine.T}, {piscine.F, piscine.T}})
		}
		if form == 9{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.F, piscine.T, piscine.T}, {piscine.T, piscine.T, piscine.F}})
		}
		if form == 10{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.F, piscine.F}, {piscine.T, piscine.T, piscine.T}})
		}
		if form == 11{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.F, piscine.T, piscine.F}, {piscine.T, piscine.T, piscine.T}})
		}
		if form == 12{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.F, piscine.F, piscine.T}, {piscine.T, piscine.T, piscine.T}})
		}
		if form == 13{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.F}, {piscine.T, piscine.T}, {piscine.F, piscine.T}})
		}
		if form == 14{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.F, piscine.T}, {piscine.T, piscine.T}, {piscine.T, piscine.F}})
		}
		if form == 15{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.F}, {piscine.T, piscine.T}, {piscine.T, piscine.F}})
		}
		if form == 16{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.F, piscine.T}, {piscine.T, piscine.T}, {piscine.F, piscine.T}})
		}
		if form == 17{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T, piscine.F}, {piscine.T, piscine.F}, {piscine.T, piscine.T}})
		}
		if form == 18{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.F, piscine.T}, {piscine.F, piscine.T}, {piscine.T, piscine.T}})
		}
		if form == 19{
			pieces = append(pieces, [][]piscine.SmartBool{{piscine.T}, {piscine.T}, {piscine.T}, {piscine.T}})
		}
		if len(pieces) > 26 {
			piscine.Puts("Impossible")
			return
		}
	}
	piscine.Solve(pieces)
	fmt.Println("-----------------")
	}
}

func testFile() {
	fmt.Println("-----------------")
	fmt.Println(os.Args[1])
	fmt.Println("-----------------")
	pieces, err := piscine.ConvertFile(os.Args[1])
	if err != nil {
		piscine.Puts(err.Error())
		return
	}
	if len(pieces) > 26 {
		piscine.Puts("Impossible")
		return
	}
	piscine.Solve(pieces)
	fmt.Println("-----------------")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("One arg pls")
		return
	}
	if  strings.Contains(os.Args[1], ".fillit"){
		testFile()
	} else if isNumeric(os.Args[1]) {
		randomGeneration()
	} else { 
		testDirectory()
	}
}
