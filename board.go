package main

type SmartBool bool

const (
	T SmartBool = true
	F SmartBool = false
)

const TILE = '#'
const SPACE = '.'

type Board struct {
	Field [][]byte
}

type Info struct {
	Board    Board
	AnsBoard Board
	Length   int
	Pieces   []Piece // 10
	Used     []bool
}

/*
type Tetrimino struct {
	piece Piece
	x int
	y int
}
*/

type Piece [][]SmartBool

var (
	TETRIMINO1  Piece   = [][]SmartBool{{T, T, T, T}}
	TETRIMINO2  Piece   = [][]SmartBool{{T, T, T}, {T, F, F}}
	TETRIMINO3  Piece   = [][]SmartBool{{T, T, T}, {F, T, F}}
	TETRIMINO4  Piece   = [][]SmartBool{{T, T, T}, {F, F, T}}
	TETRIMINO5  Piece   = [][]SmartBool{{T, T}, {T, T}}
	TETRIMINO6  Piece   = [][]SmartBool{{T, T, F}, {F, T, T}}
	TETRIMINO7  Piece   = [][]SmartBool{{T, T}, {T, F}, {T, F}}
	TETRIMINO8  Piece   = [][]SmartBool{{T, T}, {F, T}, {F, T}}
	TETRIMINO9  Piece   = [][]SmartBool{{F, T, T}, {T, T, F}}
	TETRIMINO10 Piece   = [][]SmartBool{{T, F, F}, {T, T, T}}
	TETRIMINO11 Piece   = [][]SmartBool{{F, T, F}, {T, T, T}}
	TETRIMINO12 Piece   = [][]SmartBool{{F, F, T}, {T, T, T}}
	TETRIMINO13 Piece   = [][]SmartBool{{T, F}, {T, T}, {F, T}}
	TETRIMINO14 Piece   = [][]SmartBool{{F, T}, {T, T}, {T, F}}
	TETRIMINO15 Piece   = [][]SmartBool{{T, F}, {T, T}, {T, F}}
	TETRIMINO16 Piece   = [][]SmartBool{{F, T}, {T, T}, {F, T}}
	TETRIMINO17 Piece   = [][]SmartBool{{T, F}, {T, F}, {T, T}}
	TETRIMINO18 Piece   = [][]SmartBool{{F, T}, {F, T}, {T, T}}
	TETRIMINO19 Piece   = [][]SmartBool{{T}, {T}, {T}, {T}}
	TETRIMINOS  []Piece = []Piece{TETRIMINO1, TETRIMINO2, TETRIMINO3, TETRIMINO4, TETRIMINO5,
		TETRIMINO6, TETRIMINO7, TETRIMINO8, TETRIMINO9, TETRIMINO10, TETRIMINO11, TETRIMINO12,
		TETRIMINO13, TETRIMINO14, TETRIMINO15, TETRIMINO16, TETRIMINO17, TETRIMINO18, TETRIMINO19}
)

// ファイルから読み込んで、Piece型に変換
func GetPiece(input [][]bool) Piece {
	for i := range input {
		for j := range input[i] {
			for _, tetrimino := range TETRIMINOS {
				if i+len(tetrimino) > len(input) || j+len(tetrimino[0]) > len(input[i]) {
					continue
				}
				for k := range tetrimino {
					for l := range tetrimino[k] {
						if input[i+k][j+l] != bool(tetrimino[k][l]) {
							goto NG
						}
					}
				}
				return tetrimino
			NG:
			}
		}
	}
	return nil
}

/*
func (board *Board) Set(Field Board, length int) {
	board = Field
}
*/

func InfoNew(length int, pieces []Piece) *Info {
	info := &Info{}
	info.Pieces = pieces
	info.Board.New(length)
	// TODO AnsBoard
	info.Used = make([]bool, len(pieces))
	info.Length = length
	return info
}

func (board *Board) New(length int) bool {
	board.Field = make([][]byte, length)
	if board.Field == nil {
		return false
	}
	for i := range board.Field {
		board.Field[i] = make([]byte, length)
		if board.Field[i] == nil {
			return false
		}
		for j := range board.Field[i] {
			board.Field[i][j] = SPACE
		}
	}
	return true
}

//&
func (board *Board) IsFit(piece Piece, y int, x int) bool {
	if len(piece) == 0 || len(board.Field) == 0 {
		return false
	}
	if y+len(piece) > len(board.Field) || x+len(piece[0]) > len(board.Field[0]) {
		return false
	}
	for i := range piece {
		for j := range piece[i] {
			if board.Field[y+i][x+j] != SPACE && bool(piece[i][j]) {
				return false
			}
		}
	}
	return true
}

//+
func (board *Board) Set(piece Piece, c byte, y int, x int) {
	if len(piece) == 0 || len(board.Field) == 0 {
		return
	}
	if piece[0][0] == F {
		if piece[0][1] == F {
			// 2マスpieceを左にずらす
		} else {
			// 1マス左にずらす
		}
	}
	for i := range piece {
		for j := range piece[i] {
			if piece[i][j] {
				board.Field[y+i][x+j] = c
			}
		}
	}
}

func (board *Board) Dup() Board {
	dst := Board{}
	dst.Field = make([][]byte, len(board.Field))
	for i, v := range board.Field {
		dst.Field[i] = make([]byte, len(v)) 
		copy(dst.Field[i], v)
	}
	return dst
}

//-
// func (board *Board) Remove(piece Piece, y int, x int) {
// 	if len(piece) == 0 || len(board.Field) == 0 {
// 		return
// 	}
// 	for i := range piece {
// 		for j := range piece[i] {
// 			if piece[i][j] {
// 				board.Field[y+i][x+j] = SPACE
// 			}
// 		}
// 	}
// }

//<<
//func

/*
	TETRIMINO1 Piece = [][]SmartBool{{T, T, T, T}, {F, F, F, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO2 Piece = [][]SmartBool{{T, T, T, F}, {T, F, F, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO3 Piece = [][]SmartBool{{T, T, T, F}, {F, T, F, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO4 Piece = [][]SmartBool{{T, T, T, F}, {F, F, T, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO5 Piece = [][]SmartBool{{T, T, F, F}, {T, T, F, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO6 Piece = [][]SmartBool{{T, T, F, F}, {F, T, T, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO7 Piece = [][]SmartBool{{T, T, F, F}, {T, F, F, F}, {T, F, F, F}, {F, F, F, F}}
	TETRIMINO8 Piece = [][]SmartBool{{T, T, F, F}, {F, T, F, F}, {F, T, F, F}, {F, F, F, F}}
	TETRIMINO9 Piece = [][]SmartBool{{F, T, T, F}, {T, T, F, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO10 Piece = [][]SmartBool{{T, F, F, F}, {T, T, T, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO11 Piece = [][]SmartBool{{F, T, F, F}, {T, T, T, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO12 Piece = [][]SmartBool{{F, F, T, F}, {T, T, T, F}, {F, F, F, F}, {F, F, F, F}}
	TETRIMINO13 Piece = [][]SmartBool{{T, F, F, F}, {T, T, F, F}, {F, T, F, F}, {F, F, F, F}}
	TETRIMINO14 Piece = [][]SmartBool{{F, T, F, F}, {T, T, F, F}, {T, F, F, F}, {F, F, F, F}}
	TETRIMINO15 Piece = [][]SmartBool{{T, F, F, F}, {T, T, F, F}, {T, F, F, F}, {F, F, F, F}}
	TETRIMINO16 Piece = [][]SmartBool{{F, T, F, F}, {T, T, F, F}, {F, T, F, F}, {F, F, F, F}}
	TETRIMINO17 Piece = [][]SmartBool{{T, F, F, F}, {T, F, F, F}, {T, T, F, F}, {F, F, F, F}}
	TETRIMINO18 Piece = [][]SmartBool{{F, T, F, F}, {F, T, F, F}, {T, T, F, F}, {F, F, F, F}}
	TETRIMINO19 Piece = [][]SmartBool{{T, F, F, F}, {T, F, F, F}, {T, F, F, F}, {T, F, F, F}}
*/
