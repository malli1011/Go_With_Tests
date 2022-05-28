Go has really taken off lately. More and more companies are adopting it, and developers are generally accepting it, because it is easy to learn, yet powerful to use.

If you are currently at the beginning of your Go journey, here are 10 useful CLI commands that you will probably be using every day while working with Golang.

Before you can do anything with Go on your local machine, you need to install the Go compiler. You can check if it is already installed by running:

`go version`

If it’s already installed you should something like this:

go version go1.17.6 linux/amd64

If it is not installed, you can install it from here.

Next, you might want to check some environment variables related to Go, that might be useful such as GOROOT or GOPATH. You can see them by running:

`go env`

Now you can actually start writing code. The first command you should run at the beginning of every Go project is:

`go mod init <project name>`

It will initialize a go.mod file, which is something similar to a pom.xml if you are coming from Java, or package.json if you are coming from JavaScript.

Next, you want to install a certain third party library or framework which you will use in your project:

`go get <package name>`

You then write some Go code, and want to run it locally to see if everything works as expected:

`go run .`

One of the neatest things about Go is that you can compile your entire code into a single binary. The command for doing that is:

`go build -o <name of binary> <packages>`

The compile time is also extremely fast compared to other compiled languages like Java or C++ .

The code builds and you are happy. Before pushing it to Github, you will want to format it, so your PR doesn’t get a please format your code before merging comment:

`go fmt`

After experimenting with a couple of frameworks, you decide to drop all of them from your code base and write everything yourself. But, your go.mod file still contains these unused modules. We can fix this by running:

`go mod tidy`

This command will remove all unused modules from your go.mod file, so you don’t have to manually edit this file.

Next, you write a bunch of unit tests, so that you can easily check if any updates to your code will break existing features. In order to run your tests in a given directory simply execute:

`go test`

Run all tests under a given directory including sub directories. 

`go test ./...`

Finally, an interesting command you may want to run is:

`go vet`

As stated in the official Go docs here:

Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.