package asteroids

import (
	"fmt"
	"unicode"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type (
	Score struct {
		score int
		level int

		text   *text.Text
		bounds *pixel.Rect
	}

	Drawable interface {
		Update(delta float64)
		Draw(imd *imdraw.IMDraw)
	}
)

func NewScore(bounds *pixel.Rect) Score {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII, text.RangeTable(unicode.Latin))
	position := pixel.V(0, bounds.H()-atlas.LineHeight()*1.5)

	return Score{
		bounds: bounds,
		text:   text.New(position, atlas),
	}
}

func (s *Score) Update(delta float64) {
	fmt.Fprintf(s.text, "Hello World")
}

func (s *Score) Draw(imd *imdraw.IMDraw) {
	s.text.Color = colornames.White
	//s.text.Draw(imd, pixel.IM.Scaled(s.text.Orig, 4.0))
}
