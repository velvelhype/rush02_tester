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

func sqrt(square int) int {
	for i:=1; ; i++ {
		if i*i >= square {
			return i
		}
	}
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
		reader := piscine.Read(file)
		if reader == nil {
			return
		}
		pieces := Convert(reader)
		if pieces == nil {
			return
		}
		info := InfoNew(sqrt(len(pieces)*4), pieces)
		Solve(info)
		fmt.Println("-----------------")
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//とりあえず試行回数は一回で、何回やるかもとるか
func randomGeneration() {
	fmt.Printf("random generation : %s elements\n", os.Args[1])
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	// ここひ引数分ランダム駒を作っていき、solveにかける
	for j := i; j > 0; j-- {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(randomInt(1, 19))
	}

	reader := piscine.Read(os.Args[1])
	if reader == nil {
		return
	}
	pieces := Convert(reader)
	if pieces == nil {
		return
	}
	info := InfoNew(sqrt(len(pieces)*4), pieces)
	Solve(info)
	fmt.Println("-----------------")

}

func testFile() {
	fmt.Println("-----------------")
	fmt.Println(os.Args[1])
	fmt.Println("-----------------")
	reader := piscine.Read(os.Args[1])
	if reader == nil {
		return
	}
	pieces := Convert(reader)
	if pieces == nil {
		return
	}
	info := InfoNew(sqrt(len(pieces)*4), pieces)
	Solve(info)
	fmt.Println("-----------------")
}

func main() {
	if  strings.Contains(os.Args[1], ".fillit"){
		testFile()
	} else if isNumeric(os.Args[1]) {
		randomGeneration()
	} else { 
		testDirectory()
	}
}
