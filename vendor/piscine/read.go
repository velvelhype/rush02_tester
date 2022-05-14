package piscine

import "os"

const BUFFERSIZE = 10


func Read(str string) [][][]byte {
	file, err := os.Open(str)
	if err != nil {
		return nil
	}
	defer file.Close()
	var input string
	buf := make([]byte, BUFFERSIZE)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		input += string(buf[:n])
	}
	var pieces [][][]byte
	pieces = append(pieces, make([][]byte, 0))
	flag := true
	count := 0
	i := 0
	for _, b := range []byte(input) {
		if b == '\n' {
			if flag {
				pieces = append(pieces, make([][]byte, 0))
				i = 0
				count++
			} else {
				i++
			}
			flag = true
		} else {
			if flag {
				pieces[count] = append(pieces[count], []byte{b})
			} else {
				pieces[count][i] = append(pieces[count][i], b)
			}
			flag = false
		}
	}
	return pieces
}
