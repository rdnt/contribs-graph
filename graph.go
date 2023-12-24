// Package graph allows rendering an isometric projection of a GitHub user's contribution history.
package graph

import (
	"io"
	"math"

	svg "github.com/ajstarks/svgo/float"
	"github.com/akrennmair/slice"
)

// Graph is a graph instance that is used to render a contributions graph.
type Graph struct {
	cols []col
}

type ContributionDay struct {
	Count int
	Color string
}

// NewGraph creates a graph instance with the contributions-per-day that are
// passed. Graph.Render can then be used to
// render the graph.
func NewGraph(contribs []ContributionDay) *Graph {
	var cols []col
	var total int

	for k, entry := range contribs {
		total += entry.Count

		z := float64(entry.Count*2 + 2)

		i := float64(k / 7)
		j := float64(k % 7)

		size := float64(10)
		space := float64(12)

		var pts []vector3
		// top
		pts = append(pts, vector3{X: space * i, Y: space * j, Z: -z})
		pts = append(pts, vector3{X: space*i + size, Y: space * j, Z: -z})
		pts = append(pts, vector3{X: space*i + size, Y: space*j + size, Z: -z})
		pts = append(pts, vector3{X: space * i, Y: space*j + size, Z: -z})
		// bottom
		pts = append(pts, vector3{X: space * i, Y: space * j, Z: 0})
		pts = append(pts, vector3{X: space*i + size, Y: space * j, Z: 0})
		pts = append(pts, vector3{X: space*i + size, Y: space*j + size, Z: 0})
		pts = append(pts, vector3{X: space * i, Y: space*j + size, Z: 0})

		cols = append(cols, col{pts: pts, count: entry.Count, color: entry.Color})
	}

	return &Graph{
		cols: cols,
	}
}

// Render writes the generated SVG to the given io.Writer in the appropriate Theme.
func (g *Graph) Render(f io.WriteCloser, theme Theme, isHalloween bool) error {
	canvas := svg.New(f)

	canvas.Start(840, 400)

	for _, c := range g.cols {
		// each column has 3 visible faces.
		// p1 is the outline of the column (all 3 visible faces), p2 is the bottom 2 faces and 3 is the bottom right face
		// they are rendered that way to avoid weird spacing artifacts between the faces
		p1 := []vector3{c.pts[0], c.pts[1], c.pts[5], c.pts[6], c.pts[7], c.pts[3]}
		p2 := []vector3{c.pts[3], c.pts[2], c.pts[1], c.pts[5], c.pts[6], c.pts[7]}
		p3 := []vector3{c.pts[2], c.pts[1], c.pts[5], c.pts[6]}

		// project the 3d coordinates to 2d space using an isometric projection
		p1iso := isometricProjection(p1)
		p2iso := isometricProjection(p2)
		p3iso := isometricProjection(p3)

		// horizontal & vertical offsets to center the whole chart
		h, v := float64(180), float64(25)

		// colors used for the 3 visible faces
		c1, c2, c3 := faceColors(c.color, theme, isHalloween)

		xs := slice.Map(p1iso, func(vec vector2) float64 { return vec.X + h })
		ys := slice.Map(p1iso, func(vec vector2) float64 { return vec.Y + v })
		canvas.Polygon(xs, ys, "fill:"+c1)

		xs = slice.Map(p2iso, func(vec vector2) float64 { return vec.X + h })
		ys = slice.Map(p2iso, func(vec vector2) float64 { return vec.Y + v })
		canvas.Polygon(xs, ys, "fill:"+c2)

		xs = slice.Map(p3iso, func(vec vector2) float64 { return vec.X + h })
		ys = slice.Map(p3iso, func(vec vector2) float64 { return vec.Y + v })
		canvas.Polygon(xs, ys, "fill:"+c3)
	}

	canvas.End()

	return f.Close()
}

func isometricProjection(v []vector3) []vector2 {
	return slice.Map(v, func(v vector3) vector2 {
		x, y := spaceToIso(v.X, v.Y, v.Z)

		return vector2{
			X: x,
			Y: y,
		}
	})
}

func spaceToIso(x, y, z float64) (h, v float64) {
	x, y = x+z, y+z

	h = (x - y) * math.Sqrt(3) / 2
	v = (x + y) / 2

	return h, v
}

const (
	color0 = "#ebedf0"
	color1 = "#9be9a8"
	color2 = "#40c463"
	color3 = "#30a14e"
	color4 = "#216e39"

	halloweenColor0 = "#ebedf0"
	halloweenColor1 = "#ffee4a"
	halloweenColor2 = "#ffc501"
	halloweenColor3 = "#fe9600"
	halloweenColor4 = "#03001c"
)

func faceColors(color string, theme Theme, isHalloween bool) (string, string, string) {
	switch theme {
	case Dark:
		if isHalloween {
			switch color {
			case halloweenColor0:
				return "#161b22", "#000000", "#00050c"
			case halloweenColor1:
				return "#631c03", "#390000", "#4d0800"
			case halloweenColor2:
				return "#bd561d", "#942d00", "#a84109"
			case halloweenColor3:
				return "#fa7a18", "#d05000", "#e46403"
			case halloweenColor4:
				return "#fddf68", "#d3b53e", "#e7c952"
			}
		} else {
			switch color {
			case color0:
				return "#2d333b", "#030a12", "#171e26"
			case color1:
				return "#0e4429", "#001b00", "#002f12"
			case color2:
				return "#006d32", "#004307", "#00571b"
			case color3:
				return "#26a641", "#007d1a", "#11912e"
			case color4:
				return "#39d353", "#10a92c", "#24bd40"
			}
		}

		return "#2d333b", "#030a12", "#171e26"
	default:
		fallthrough
	case Light:
		if isHalloween {
			switch color {
			case halloweenColor0:
				return "#ebedf0", "#c2c5c8", "#d6d9dc"
			case halloweenColor1:
				return "#ffee4a", "#d5c522", "#e9d936"
			case halloweenColor2:
				return "#ffc501", "#d59d00", "#e9b100"
			case halloweenColor3:
				return "#fe9600", "#d56e00", "#e98200"
			case halloweenColor4:
				return "#03001c", "#000001", "#000007"
			}
		} else {
			switch color {
			case color0:
				return "#ebedf0", "#c2c5c8", "#d6d9dc"
			case color1:
				return "#9be9a8", "#73c080", "#87d494"
			case color2:
				return "#40c463", "#199b3c", "#2daf50"
			case color3:
				return "#30a14e", "#077725", "#1b8b39"
			case color4:
				return "#216e39", "#004410", "#0c5824"
			}
		}

		return "#ebedf0", "#c2c5c8", "#d6d9dc"
	}
}