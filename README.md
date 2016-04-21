# led-web-interface

A web interface used to change LED lights plugged into a Raspberry Pi's GPIO ports. Needs be decoupled a bit 
better for more general use. Right now you make 3-coloured or single-coloured lights by setting their GPIO pins (examples in code).
The handling of user input is pretty specific to the lights you want to control right now. 

To install: 
```bash
# go get https://github.com/dangodai/led-web-interface
``` 

From within the directory:

```bash
# go install
# cp $GOBIN/led-web-interface .
# sudo ./led-web-interface
```
(Or just make sure static/templates are in the same directory as the binary)

Visit localhost in your browser. 
