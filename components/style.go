package components

import (
	"golang.org/x/image/font"
	"image/color"
)

type Text struct {
	text string
	color color.Color
	font font.Face
}

func NewText(t string, c color.Color, f font.Face) Text {
	return Text{
		text: t,
		color: c,
		font: f,
	}
}
