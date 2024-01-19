package main

import "fmt"

type device interface {
	turnOn() string
}

type iphone struct {
	name  string
	model string
}

type imac struct {
	name  string
	model string
}

func (phone iphone) turnOn() string {
	return "iOS starting up..."
}

func (mac imac) turnOn() string {
	return "macOS starting up..."
}

func main() {
	dev1 := iphone{"iPhone", "11 Pro"}
	dev2 := imac{"iMac", "27 5k Retina"}

	devices := []device{dev1, dev2}

	for _, dev := range devices {
		fmt.Println(dev.turnOn())
	}
}

/*
An interface in Go is a collection of one or more method signatures that collectively describe the behavior of a type that implements it.
By their very definition, interfaces are to be considered abstract, meaning you cannot create instances directly from them.
Instead, what you can do is declare variables with the interface type and then assign them to instances of types that happen to implement
all the methods declared within the interface.

Now, one of the really interesting things that Go does when working with interfaces, is that it allows types to implicitly implement
an interface if that type implements the full range of method signatures within the interface.
Let's see this in action. In the following example, a custom device interface is declared containing a single method named turnOn.
I then declare two custom struct types, an iphone struct and an imac struct.

Both structs contain the two fields name and model. Now, on lines 19 to 21 and lines 23 to 25, I implement two methods with
different receiver arguments, one for the iphone struct and one for the imac struct, respectively.
Both methods have the same turnOn name, which is also importantly the method name declared within the device interface, this
being the turnOn method. The method implementations simply return a string indicating the flavor of the operating system that boots up when started.

In this case, the receiver arguments are declared as non-pointer types since neither method is updating any of the data on the receiver argument within the body.
Now, further down within the main function, I create instances of both the iphone and imac types on lines 28 and 29.
Now, what follows is really the magic of working with interfaces, and demonstrates the deeper concept of polymorphism.
Line 31 creates and initializes the device's slice, which importantly here is typed using the device interface.
I then add both the iphone and imac struct instances to the device's slice. Finally, on lines 33, 34 and 35, I iterate over the items
in the device's slice using the range keyword, and then for each item call its turnOn method.

Again, keep in mind that both the iphone and imac items are actually different in terms of their types, but importantly they share the
same common device interface since both implement the same method signature as declared within the device interface.
This particular implicit typing system is an extremely useful tool and abstraction mechanism that aids the overall development
experience when working with large code bases within Go.
Let's now run this example and examine the resulting output. Here we can see the different boot up messages for each of the two different devices.

In summary, you have just observed how to create custom interface types, and how having the capability to abstract
behaviors into interfaces, as just demonstrated, helps to create reusable and maintainable code, particularly, again, in large code bases.
*/
