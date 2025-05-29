package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

const size = 3

func main() {
	a := app.New()
	w := a.NewWindow("三目並べ（Tic Tac Toe）")
	w.Resize(fyne.NewSize(300, 350))

	board := [size][size]string{}
	turn := "O"
	buttons := [size][size]*widget.Button{}
	texts := [size][size]*canvas.Text{}
	info := widget.NewLabel("先手: O")

	resetBoard := func() {
		for i := range board {
			for j := range board[i] {
				board[i][j] = ""
				texts[i][j].Text = ""
				texts[i][j].Refresh()
				buttons[i][j].Enable()
			}
		}
		turn = "O"
		info.SetText("先手: O")
	}

	checkWinner := func() string {
		for i := 0; i < size; i++ {
			if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
				return board[i][0]
			}
			if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
				return board[0][i]
			}
		}
		if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
			return board[0][0]
		}
		if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
			return board[0][2]
		}
		return ""
	}

	isFull := func() bool {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if board[i][j] == "" {
					return false
				}
			}
		}
		return true
	}

	grid := container.NewGridWithRows(size)
	for i := 0; i < size; i++ {
		row := container.NewGridWithColumns(size)
		for j := 0; j < size; j++ {
			x, y := i, j
			text := canvas.NewText("", color.White)
			text.Alignment = fyne.TextAlignCenter
			text.TextSize = 48

			btn := widget.NewButton("", func() {
				if board[x][y] == "" {
					board[x][y] = turn
					texts[x][y].Text = turn
					texts[x][y].Refresh()
					buttons[x][y].Disable()

					winner := checkWinner()
					if winner != "" {
						info.SetText(fmt.Sprintf("勝者: %s", winner))
						for ii := 0; ii < size; ii++ {
							for jj := 0; jj < size; jj++ {
								buttons[ii][jj].Disable()
							}
						}
						return
					}
					if isFull() {
						info.SetText("引き分け！")
						return
					}
					if turn == "O" {
						turn = "X"
					} else {
						turn = "O"
					}
					info.SetText(fmt.Sprintf("次の手: %s", turn))
				}
			})
			buttons[i][j] = btn
			texts[i][j] = text

			// ボタン＋白色テキストを重ねて配置
			cell := container.NewMax(btn, text)
			row.Add(cell)
		}
		grid.Add(row)
	}

	resetBtn := widget.NewButton("リセット", func() {
		resetBoard()
	})

	content := container.NewVBox(
		info,
		grid,
		resetBtn,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
