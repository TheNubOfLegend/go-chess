package main

import "fmt"

type Player struct {
    Color
}

type Tile struct {
    Piece
    Empty bool
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

type Color bool

const (
    White Color = true
    Black Color = false
)

func main() {
    var board [64]Tile = initBoard()
    drawBoard(board)
    Black.movePiece(&board, 11, 19)
    drawBoard(board)
}

func (player Color) movePiece(board *[64]Tile, pos int, newPos int) {
    if !board[pos].Empty && board[pos].Piece.Color == player {
        if (board[pos].Piece.ValidMoves & (1 << newPos)) > 0 {
            board[newPos] = board[pos]
            board[pos].Empty = true
        } else {
            fmt.Print("1")
        }
    } else {
        fmt.Print("2")
    }
}

func findValidMoves(board [64]Tile, pos int) {
    board[pos].ValidMoves = 0
    piece := board[pos]
    switch piece.Kind {
    case Pawn:
        //TODO: passant
        var colorDir int
        if piece.Color {
            colorDir = 1
        } else {
            colorDir = -1
        }
        if piece.Color == Color.White && pos > 7 && pos < 16 && board[pos + 16].Empty {
            piece.ValidMoves |= 1 << (pos + 16)
        } else if piece.Color == Color.Black && pos > 47 && pos < 56 && board[pos - 16].Empty {
            piece.ValidMoves |= 1 << (pos - 16)
        }
        if !board[pos + colorDir * 9].Empty && (pos + colorDir * 9) % 8 != 0 && piece.Color != board[pos + colorDir * 9].Color {
            piece.ValidMoves |= 1 << (pos + colorDir * 9)
        }
        if !board[pos + colorDir * 9].Empty && pos % 8 != 0 && piece.Color != board[pos + colorDir * 7].Color {
            piece.ValidMoves |= 1 << (pos + colorDir * 7)
        }
        if board[pos + colorDir * 8].Empty {
            piece.ValidMoves |= 1 << (pos + colorDir * 8)
        }
    case Rook:
        //up
        for i := pos + 8; i < 64; i += 8 {
            if board[i].Empty {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
        //down
        for i := pos - 8; i >= 0; i -= 8 {
            if board[i].Empty {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
        //left
        for i := pos - 1; i >= 0; i -= 1 {
            if board[i].Empty {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
        //right
        for i := pos + 1; i >= 0; i += 1 {
            if board[i].Empty {
                piece.ValidMoves |= 1 << (i)
            } else if board[i].Color != piece.Color {
                piece.ValidMoves |= 1 << (i)
                break
            }
        }
    case Knight:
        if 


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
    }
}

func drawBoard(board [64]Tile) {
    for i, tile := range board {
        if i % 8 == 0 {
            fmt.Println()
        }
        if !tile.Empty {
            drawPiece(tile.Piece)
        } else {
            fmt.Print(" ")
        }
        fmt.Print(" ")
    }
}

func initBoard() [64]Tile {
    var board [64]Tile
    board[0] = Tile { Piece {Black, Rook, 0}, false }
    board[1] = Tile { Piece {Black, Knight, 0}, false }
    board[2] = Tile { Piece {Black, Bishop, 0}, false }
    board[3] = Tile { Piece {Black, Queen, 0}, false }
    board[4] = Tile { Piece {Black, King, 0}, false }
    board[5] = Tile { Piece {Black, Bishop, 0}, false }
    board[6] = Tile { Piece {Black, Knight, 0}, false }
    board[7] = Tile { Piece {Black, Rook, 0}, false }
    board[8] = Tile { Piece {Black, Pawn, 0}, false }
    board[9] = Tile { Piece {Black, Pawn, 0}, false }
    board[10] = Tile { Piece {Black, Pawn, 0}, false }
    board[11] = Tile { Piece {Black, Pawn, 0}, false }
    board[12] = Tile { Piece {Black, Pawn, 0}, false }
    board[13] = Tile { Piece {Black, Pawn, 0}, false }
    board[14] = Tile { Piece {Black, Pawn, 0}, false }
    board[15] = Tile { Piece {Black, Pawn, 0}, false }
    board[16] = Tile { Piece {Black, Pawn, 0}, true }
    board[17] = Tile { Piece {Black, Pawn, 0}, true }
    board[18] = Tile { Piece {Black, Pawn, 0}, true }
    board[19] = Tile { Piece {Black, Pawn, 0}, true }
    board[20] = Tile { Piece {Black, Pawn, 0}, true }
    board[21] = Tile { Piece {Black, Pawn, 0}, true }
    board[22] = Tile { Piece {Black, Pawn, 0}, true }
    board[23] = Tile { Piece {Black, Pawn, 0}, true }
    board[24] = Tile { Piece {Black, Pawn, 0}, true }
    board[25] = Tile { Piece {Black, Pawn, 0}, true }
    board[26] = Tile { Piece {Black, Pawn, 0}, true }
    board[27] = Tile { Piece {Black, Pawn, 0}, true }
    board[28] = Tile { Piece {Black, Pawn, 0}, true }
    board[29] = Tile { Piece {Black, Pawn, 0}, true }
    board[30] = Tile { Piece {Black, Pawn, 0}, true }
    board[31] = Tile { Piece {Black, Pawn, 0}, true }
    board[32] = Tile { Piece {Black, Pawn, 0}, true }
    board[33] = Tile { Piece {Black, Pawn, 0}, true }
    board[34] = Tile { Piece {Black, Pawn, 0}, true }
    board[35] = Tile { Piece {Black, Pawn, 0}, true }
    board[36] = Tile { Piece {Black, Pawn, 0}, true }
    board[37] = Tile { Piece {Black, Pawn, 0}, true }
    board[38] = Tile { Piece {Black, Pawn, 0}, true }
    board[39] = Tile { Piece {Black, Pawn, 0}, true }
    board[40] = Tile { Piece {Black, Pawn, 0}, true }
    board[41] = Tile { Piece {Black, Pawn, 0}, true }
    board[42] = Tile { Piece {Black, Pawn, 0}, true }
    board[43] = Tile { Piece {Black, Pawn, 0}, true }
    board[44] = Tile { Piece {Black, Pawn, 0}, true }
    board[45] = Tile { Piece {Black, Pawn, 0}, true }
    board[46] = Tile { Piece {Black, Pawn, 0}, true }
    board[47] = Tile { Piece {Black, Pawn, 0}, true }
    board[48] = Tile { Piece {White, Pawn, 0}, false }
    board[49] = Tile { Piece {White, Pawn, 0}, false }
    board[50] = Tile { Piece {White, Pawn, 0}, false }
    board[51] = Tile { Piece {White, Pawn, 0}, false }
    board[52] = Tile { Piece {White, Pawn, 0}, false }
    board[53] = Tile { Piece {White, Pawn, 0}, false }
    board[54] = Tile { Piece {White, Pawn, 0}, false }
    board[55] = Tile { Piece {White, Pawn, 0}, false }
    board[56] = Tile { Piece {White, Rook, 0}, false }
    board[57] = Tile { Piece {White, Knight, 0}, false }
    board[58] = Tile { Piece {White, Bishop, 0}, false }
    board[59] = Tile { Piece {White, King, 0}, false }
    board[60] = Tile { Piece {White, Queen, 0}, false }
    board[61] = Tile { Piece {White, Bishop, 0}, false }
    board[62] = Tile { Piece {White, Knight, 0}, false }
    board[63] = Tile { Piece {White, Rook, 0}, false }
    return board
}
