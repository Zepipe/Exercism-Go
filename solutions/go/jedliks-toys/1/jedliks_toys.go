package jedlik

import ("fmt"
        "strconv"
        )

func (c *Car) Drive() {
    if c.battery >= c.batteryDrain {
    	c.battery -= c.batteryDrain
    	c.distance += c.speed
    }
}

func (c Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
    return "Battery at " + strconv.Itoa(c.battery) + "%"
}

func (c Car) CanFinish(trackDistance int) bool {
    return c.battery * c.speed >= trackDistance * c.batteryDrain
}