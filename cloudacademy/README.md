# Introduction to Go

Go was first released in open source to the community in 2009 by Google and has grown in popularity year on year.

Go excels at systems programming and concurrency requirements. When you consider the success of projects like Docker, Kubernetes, Terraform, Etcd, Istio, InfluxDB, the list goes on, all of these projects were built using Go. The success and wide adoption of each of these projects is testament to Go's ability as a language that is productive, performant, and stable.

The original motives for designing Go were primarily to build a lightweight language that had strong support for concurrency which, once compiled, could make better use of multi-core systems. This was to address Google's continued thirst for processing large amounts of data as quickly as possible and as often as possible. 

While setting out to achieve this they also decided that they wanted the new language to combine the best traits of other popular languages and, therefore, designed it to have the performance and safety of compiled languages such as C and C++. 

## Language design

The Go language syntax is simple and very minimalistic by design. This has the effect of making it perhaps less expressive when compared with other languages such as Ruby and Python but it does make the code easier to read, maintain, and learn. 

The language looks and feels a bit like C. It is typed and comes with a primitive set of data types, all of which you would expect.

The language is opinionated which helps you to write code in a best practice manner. 

Go isn't strictly object orientated. For example, the language lacks the concept of inheritance. Regardless, it still has object orientated-like features such as custom struck types, methods, and dot notation for accessing. The language is designed in such a way as to preference composition over inheritance which many believe to be a better way of structuring applications

## Static typing

The Go language is statically typed meaning the Go compiler checks for and catches type mismanagement errors at compile time. The compiler checks for incorrect type conversions and casts that would fail, as well as many other issues. 

The compiler is also designed to discover and error out on any unused variables that may have been declared but are never actually used. This may not sound like a big deal but when your applications are being built to run mission critical workloads at scale and to be hardly performant then having the knowledge that your code compiles safely and is optimized will give you a level of comfort not afforded to other dynamic and interpreted languages. 

## Static analysis tools

The Go Toolchain provides a number of useful static analysis tools which, when used, helps to ensure that your Go application remains maintainable and performs as desired at runtime. Example static analysis tools are:
- gofmt - used to reformat source code layout according to Go's opinionated best practices
- golint - used to report on and enforce semantical and syntactical practices 
- go vet - used to report on suspicious or abnormal or just nonsensical code in your application

As good practice these static analysis tools can and should be integrated into your Go CICD pipelines so that they are rerun over the code each and every time a build is performed. 

## Standard library

Go ships with and installs a full featured Standard Library which provides all of your basic and midrange requirements for data processing, network connectivity, security, cryptography, IO, and the list goes on. The Standard Library provides many useful built-in functions that help you to manipulate the primitive data types. 

The Standard Library even provides web server-like support that makes it trivial to stand up an HTTP web service within your Go applications. Working with data formats, such as JSON and or XML again is trivial with the serialization and deserialization support already packaged within the Standard Library for these formats. 

```bash
go list std
```

## Module management

Go module functionality has been integrated directly into the Go Toolchain making the developer's workflow involving third party open source projects much simpler and more robust than in the past. The Go Toolchain will automatically check and detect import statements when a build is performed, and if required, download the module making it available locally so that compilation succeeds. Semantic Versioning is used to tag modules ensuring that your code uses the correct version when multiple versions of the dependent module exist.

```bash
go mod init
go get
go mod verify
go list
go mod why
```

## Performance

When it comes to assessing performance Go is right up there with the most performant languages such as C. 

Go's language from the outset was designed to exploit multi-core machines and therefore brings a lot to the table when building the applications. Building applications that run concurrently across multiple cores and multiple machines is a sweet spot for Go

## Fast compilation

Go's compiler is also renowned for being speedy with build times completing in seconds rather than minutes often being the case. Again, this was by intent and, in fact, was one of the very early motives for creating Go. Go's simple language design allows the compilation phases to be quick which benefits developers two-fold, less of the language to learn and understand, and faster and rapid build times. 

```bash
time -p go build
```

## Binaries

Go's compiler produces a single binary with all runtime dependencies compiled and packaged within it. The beauty of this is that it removes the need for you to manage and install external requirements to make it run and unlike Java it doesn't need a runtime engine installed. 

From the perspective of DevOps dealing with and maintaining a single binary is a walk in the park compared to alternatives. Go also has the cool ability to compile the same source code without modification for various operating systems and process architectures including Linux, Windows, and MacOS.

## Concurrency

The Go language really differentiates itself from competitor languages when it comes to implementing concurrency. 

The Go language introduces the concepts of go routines and channels which, when combined, makes it incredibly straightforward and almost trivial to build applications that have multiple forks of asynchronous work being performed. Go applications target multi-core systems and its ability to do so helps application scale.

## Garbage collection

At runtime, Go uses garbage collection so that you don't have to implement memory management yourself. This keeps your code clean and safe and keeps you focused on building business value rather than being distracted with memory management issues. The garbage collector is fast and keeps getting faster and more optimized with each new release of Go. 

## Unit test support

The Go Standard Library ships with a testing package which provides really excellent support for building and writing unit tests. Go encourages writing unit test files in the same location as the code it is testing using the file naming convention of name of the file under test underscore test dot go. The Go test runner will then find these files and execute them as unit tests.

## Go playground 

Although, not strictly part of the Go language, nor the runtime, the Go Playground provides an extremely useful service to those who are just starting out on their Go adventure. The Go Playground is a web based application located at https://play.golang.org. You can use it to test out and prototype various code snippets without having to install or set up the Go Toolchain locally.

# Toolchain commands

``` bash
go env # executed to display information about the current Go operating environment
go env GOROOT
go env GOPATH
go help env

go doc # consult the standard library documentation for a particular package
go doc math 
go doc math.Pow10

go run main.go  # convenient fast way to both compile and then execute immediately, allowing you to observe and use the running application
go build # performs just the compilation phase, generating a binary as its output if the compilation succeeds; will create an executable binary if the source code contains a main package with a main function declared within it
go build -o cloudacademy-demo # Typically, you'll call "go build" with the -o parameter which is used to explicitly name the executable. 


# The "go build" command provides two very cool features. The first being that it compiles down everything into a single executable binary. And the second is that it can perform cross-compilation. By default, go build will compile for the current operating system and architecture. Using various flags you can recompile the same source code for other operating system and architecture combinations

GOOS=linux 
GOARCH=amd64
go build -o godemo-linux-amd64 .

GOOS=windows
GOARCH=amd64
go build -o godemo-win-amd64.exe .

go tool dist list # full list of supported operating systems and architectures

# The "go test" command can be invoked with various options to either run tests in the current directory, in a specific directory, or across all directories recursively
go test .
go test ./...
go test ./util

# you can use the run flag with a regular expression to target matching named unit tests
go test -v -run=^TestGetX$ ./...

# Third party go modules are automatically pulled down whenever you execute the go run or go build command if they are not already local.
# Sometimes however, you may want to retrieve a particular version
go get github.com/google/uuid@v1.1.0

# see all of the dependencies that the current project is configured with. This will also now be reflected with a go.mod file
go list -m all

# The go mod command is used to manage and maintain modules. Modules are used to collect multiple related packages together
# Creating your own module can be accomplished with the command "go mod init"
go mod init
go mod init github.com/cloudacademy/godemo
go get # re-pull down any module dependencies
ls -la $(go env GOPATH)/pkg/mod/github.com/google
go list -m all # display all of the dependencies for this project

go mod verify # verify the integrity of downloaded modules ensuring that they haven't been modified. This consults the go.sum file which contains cryptographic checksums of all downloaded modules
go list -u -all # used to list out all dependent modules and whether an updated version exists and if so, what the latest semantic version number is
go mod tidy # used to prune out any required module declarations from the go.mod file if the module dependency is no longer used within the project source code anywhere

gofmt -s -d main.go  # This will reformat the code showing just the differences but not save the changes; The -s flag performs simplifications on the code. The -d flag performs a diff to highlight before and after. 

gofmt -s -d -w main.go #Extending the same command with the -w flag will cause the changes to be saved

go fmt ./... # shortcut wrapper over the gofmt tool, and will perform the same reformatting recursively across the entire project, saving all updates to the disk in one hit

go vet # used to perform a static analysis across your source code. It can be used to detect and report on code that although compiles may be considered problematic
```