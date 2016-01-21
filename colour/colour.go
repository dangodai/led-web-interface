package colour

import(
	"fmt"
)

type Colour struct {
	Red, Green, Blue string
}

func (c *Colour) String() string {
	return fmt.Sprintf("%v %v %v", c.Red, c.Green, c.Blue)
}