package main

func Convert(input [][][]byte) (output []Piece) {
	convert := convertToBool(input)
	if convert == nil {
		return nil
	}
	for _, piece := range convert {
		tmp := GetPiece(piece)
		if tmp == nil {
			return nil
		}
		output = append(output, tmp)
	}
	return
}

func convertToBool(input [][][]byte) (convertToBool [][][]bool) {
	for i, piece := range input {
		convertToBool = append(convertToBool, make([][]bool, 0))
		for j, line := range piece {
			convertToBool[i] = append(convertToBool[i], make([]bool, 0))
			for _, b := range line {
				if b == TILE {
					convertToBool[i][j] = append(convertToBool[i][j], true)
				} else if b == SPACE {
					convertToBool[i][j] = append(convertToBool[i][j], false)
				} else {
					return nil
				}
			}
		}
	}
	return
}
