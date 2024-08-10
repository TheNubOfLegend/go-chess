package main

import "fmt"

type Player struct {
    Color
}

type Piece struct {
    Color
    Kind
    ValidMoves uint64
}

type Kind int

const (
    Pawn Kind = iota
    Rook
    Knight
    Bishop
    Queen
    King
)

type Color int

const (
    White Color = iota
    Black
    None
)

func main() {
    var board [64]Piece = initBoard()
    drawBoard(board)
    Black.movePiece(&board, 11, 19)
    drawBoard(board)
}

func (player Color) movePiece(board *[64]Piece, pos int, newPos int) {
    if board[pos].Color != None && board[pos].Color == player {
        if (board[pos].ValidMoves & (1 << newPos)) > 0 {
            board[newPos] = board[pos]
            board[pos].Color = None
        } else {
            fmt.Print("1")
        }
    } else {
        fmt.Print("2")
    }
}

//must check that tile is non-empty before use
func findValidMoves(board [64]Piece, pos int) {
    board[pos].ValidMoves = 0
    piece := board[pos]
    switch piece.Kind {
    case Pawn:
        //TODO: passant
        var colorDir int
        if piece.Color == White {
            colorDir = 1
        } else if piece.Color == Black {
            colorDir = -1
        }
        if piece.Color == White && pos > 7 && pos < 16 && board[pos + 16].Color == None {
            piece.ValidMoves |= 1 << (pos + 16)
        } else if piece.Color == Black && pos > 47 && pos < 56 && board[pos - 16].Color == None {
            piece.ValidMoves |= 1 << (pos - 16)
        }
        if board[pos + colorDir * 9].Color != None && (pos + colorDir * 9) % 8 != 0 && piece.Color != board[pos + colorDir * 9].Color {
            piece.ValidMoves |= 1 << (pos + colorDir * 9)
        }
        if board[pos + colorDir * 9].Color != None && pos % 8 != 0 && piece.Color != board[pos + colorDir * 7].Color {
            piece.ValidMoves |= 1 << (pos + colorDir * 7)
        }
        if board[pos + colorDir * 8].Color == None {
            piece.ValidMoves |= 1 << (pos + colorDir * 8)
        }
    case Rook:
        //up
        for i := pos + 8; i < 64; i += 8 {
            if board[i].Color == None {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
        //down
        for i := pos - 8; i >= 0; i -= 8 {
            if board[i].Color == None {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
        //left
        for i := pos - 1; i >= 0; i -= 1 {
            if board[i].Color == None {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
        //right
        for i := pos + 1; i >= 0; i += 1 {
            if board[i].Color == None {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
    case Knight:
        if pos +  {}
    }
}


func drawPiece(piece Piece) {
    switch piece.Kind {
    case Pawn:
        if piece.Color == White {
            fmt.Print("\u2659")
        } else {
            fmt.Print("\u265f")
        }
    case Rook:
        if piece.Color == White {
            fmt.Print("\u2656")
        } else {
            fmt.Print("\u265c")
        }
    case Knight:
        if piece.Color == White {
            fmt.Print("\u2658")
        } else {
            fmt.Print("\u265e")
        }
    case Bishop:
        if piece.Color == White {
            fmt.Print("\u2657")
        } else {
            fmt.Print("\u265d")
        }
    case Queen:
        if piece.Color == White {
            fmt.Print("\u2655")
        } else {
            fmt.Print("\u265b")
        }
    case King:
        if piece.Color == White {
            fmt.Print("\u2654")
        } else {
            fmt.Print("\u265a")
        }
    default: {}
    }
}

func drawBoard(board [64]Piece) {
    for i, piece := range board {
        if i % 8 == 0 {
            fmt.Println()
        }
        if piece.Color != None {
            drawPiece(piece)
        } else {
            fmt.Print(" ")
        }
        fmt.Print(" ")
    }
}

func initBoard() [64]Piece {
    var board [64]Piece
    board[0] = Piece {Black, Rook, 0}
    board[1] = Piece {Black, Knight, 0}
    board[2] = Piece {Black, Bishop, 0}
    board[3] = Piece {Black, Queen, 0}
    board[4] = Piece {Black, King, 0}
    board[5] = Piece {Black, Bishop, 0}
    board[6] = Piece {Black, Knight, 0}
    board[7] = Piece {Black, Rook, 0}
    board[8] = Piece {Black, Pawn, 0}
    board[9] = Piece {Black, Pawn, 0}
    board[10] = Piece {Black, Pawn, 0}
    board[11] = Piece {Black, Pawn, 0}
    board[12] = Piece {Black, Pawn, 0}
    board[13] = Piece {Black, Pawn, 0}
    board[14] = Piece {Black, Pawn, 0}
    board[15] = Piece {Black, Pawn, 0}
    board[16] = Piece {None, Pawn, 0}
    board[17] = Piece {None, Pawn, 0}
    board[18] = Piece {None, Pawn, 0}
    board[19] = Piece {None, Pawn, 0}
    board[20] = Piece {None, Pawn, 0}
    board[21] = Piece {None, Pawn, 0}
    board[22] = Piece {None, Pawn, 0}
    board[23] = Piece {None, Pawn, 0}
    board[24] = Piece {None, Pawn, 0}
    board[25] = Piece {None, Pawn, 0}
    board[26] = Piece {None, Pawn, 0}
    board[27] = Piece {None, Pawn, 0}
    board[28] = Piece {None, Pawn, 0}
    board[29] = Piece {None, Pawn, 0}
    board[30] = Piece {None, Pawn, 0}
    board[31] = Piece {None, Pawn, 0}
    board[32] = Piece {None, Pawn, 0}
    board[33] = Piece {None, Pawn, 0}
    board[34] = Piece {None, Pawn, 0}
    board[35] = Piece {None, Pawn, 0}
    board[36] = Piece {None, Pawn, 0}
    board[37] = Piece {None, Pawn, 0}
    board[38] = Piece {None, Pawn, 0}
    board[39] = Piece {None, Pawn, 0}
    board[40] = Piece {None, Pawn, 0}
    board[41] = Piece {None, Pawn, 0}
    board[42] = Piece {None, Pawn, 0}
    board[43] = Piece {None, Pawn, 0}
    board[44] = Piece {None, Pawn, 0}
    board[45] = Piece {None, Pawn, 0}
    board[46] = Piece {None, Pawn, 0}
    board[47] = Piece {None, Pawn, 0}
    board[48] = Piece {White, Pawn, 0}
    board[49] = Piece {White, Pawn, 0}
    board[50] = Piece {White, Pawn, 0}
    board[51] = Piece {White, Pawn, 0}
    board[52] = Piece {White, Pawn, 0}
    board[53] = Piece {White, Pawn, 0}
    board[54] = Piece {White, Pawn, 0}
    board[55] = Piece {White, Pawn, 0}
    board[56] = Piece {White, Rook, 0}
    board[57] = Piece {White, Knight, 0}
    board[58] = Piece {White, Bishop, 0}
    board[59] = Piece {White, King, 0}
    board[60] = Piece {White, Queen, 0}
    board[61] = Piece {White, Bishop, 0}
    board[62] = Piece {White, Knight, 0}
    board[63] = Piece {White, Rook, 0}
    return board
}
