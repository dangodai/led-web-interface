package pigpio

import (
	"os/exec"
	"fmt"
	"github.com/dangodai/led-web-interface/colour"
)

type Brightness struct{
	Brightness int
}

func (io *Brightness) SetBrightness(b int) {
	_ = exec.Command("pigs", fmt.Sprintf("p %v %v", io.Brightness, b)).Run()
}

type FixedColour struct{
	Brightness
}

type RGB struct {
	Red, Green, Blue int
}

func (io *RGB) ExecuteColour(c colour.Colour) {
	_ = exec.Command("pigs", fmt.Sprintf("p %v %v", io.Red, c.Red)).Run()
	_ = exec.Command("pigs", fmt.Sprintf("p %v %v", io.Green, c.Green)).Run()
	_ = exec.Command("pigs", fmt.Sprintf("p %v %v", io.Blue, c.Blue)).Run()
}
