package components

import (
	"github.com/fogleman/gg"
)

func DrawCell(c *gg.Context, text Text, x float64, y float64, w float64, h float64) {
	// Draw cell
	c.DrawRectangle(x, y, w, h)
	c.SetRGB(0, 0,0) // Rect color
	c.Fill()

	// Draw text
	c.SetFontFace(text.font)
	c.SetColor(text.color)
	c.DrawStringWrapped(text.text, x + (w / 2), y + (h / 2), 0.5, 0.5, w, 1, gg.AlignCenter)
}

func DrawRow(c *gg.Context, row []Text, x float64, y float64, w float64, h float64, padding float64) {
	cellWidth := w / float64(len(row))

	for _, item := range row {
		DrawCell(c, item, x + padding, y, cellWidth - (padding * 2), h - padding)
		x += cellWidth
	}
}

func DrawTable(c *gg.Context, rows [][]Text, x float64, y float64, w float64, h float64, padding float64) {
	rowHeight := h / float64(len(rows))

	for _, row := range rows {
		DrawRow(c, row, x, y + padding, w, rowHeight - (padding * 2), padding)
		y += rowHeight
	}
}
