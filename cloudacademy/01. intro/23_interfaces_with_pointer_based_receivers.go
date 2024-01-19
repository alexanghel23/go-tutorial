package main

import (
	"fmt"
	"strings"
)

type device interface {
	turnOn() string
	update(version float32)
}

type iphone struct {
	name    string
	model   string
	version float32
}

type imac struct {
	name    string
	model   string
	version float32
}

func (phone iphone) turnOn() string {
	return "iOS starting up..."
}

func (mac imac) turnOn() string {
	return "macOS starting up..."
}

func (phone *iphone) update(version float32) {
	phone.version = version
}

func (mac *imac) update(version float32) {
	mac.version = version
}

func main() {
	dev1 := iphone{"iPhone", "11 Pro", 13.1}
	dev2 := imac{"iMac", "27 5k Retina", 10.15}

	devices := []device{&dev1, &dev2}

	for _, dev := range devices {
		if strings.Contains(dev.turnOn(), "iOS") {
			dev.update(14.0)
		} else if strings.Contains(dev.turnOn(), "macOS") {
			dev.update(11.00)
		}
	}

	fmt.Println(dev1)
	fmt.Println(dev2)
}

/*
The following code example extends on the previous example to demonstrate Methods that have pointer based receivers, which provides
us with the ability to update the data within the underlying struct, such that the changes persist, importantly, beyond the
lifetime of the Method execution call.

For starters, the device interface has been updated to include a second method.
This one is named update and takes a single input parameter named version. Two new concrete implementations are then created on lines 31 to 33 and lines 35 to 37.
Both of these Methods have pointer based receivers. This is done to allow the updates made to the version field permanent.

You may have noticed that on lines 32 and 36, I don't use the dereferencing operator, even though the input variables are pointers.
The reason for this is that they are automatically dereferenced by the compiler since the underlying type is a struct and dot notation is being used on it.
This is really just syntactic sugar that the compiler knows about.
Down within the main function, line 43 is updated to pass in the memory addresses of the two devices.
This can be observed by the fact that both dev1 and dev2, or device1 and device2, are now prefixed with the ampersand character.

Finally, within the for loop spanning lines 45 to 51, the new update Method is called on each device.
Now, since we are using pointer receivers for the new update Methods, the updates applied to the version field will become permanent, meaning that
if we print out the whole device struct for both device1 and device2, as we do on lines 53 and 54, you will indeed see the updated version numbers.
This shows that both devices were passed by reference to the update Methods. Let's now run this example and examine the resulting output.
Here we can see that the operating system version for both devices has indeed been updated.

In summary, you have just observed how to implement Methods that use pointer based receivers, to edit and importantly persist data changes to the underlying struct type.
*/
