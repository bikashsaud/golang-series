### GO

#### Introduction

> code -> compile --> executable file[.exe]

>compile-time, not runtime

> Go organizes code into packages.--> models or libraries.

> Functions are the building blocks of Go programs. 

#### What is MISSING?
> TRY Catch
> 

#### Is OOP Language ?
> YES/No

#### Advantages
> Automatically import packages...


###  Installation...

> Note: Here 1.x --> current version or preferred version of GO.

[] wget https://dl.google.com/go/go1.x.linux-amd64.tar.gz

[] sudo tar -C /usr/local -xzf go1.x.linux-amd64.tar.gz

>echo 'export GOROOT=/usr/local/go' >> ~/.profile
echo 'export PATH=$PATH:$GOROOT/bin' >> ~/.profile
echo 'export GOPATH=$HOME/go' >> ~/.profile
source ~/.profile

[x] go version

#### Simple Example
> package main
import "fmt"
func main() {
	fmt.Println("Hello from bikash")
}
