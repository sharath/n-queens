# n-queens
Solving n-queens problem with Genetic Programming in Go

#### Running
`go get github.com/sharath/n-queens`

`cd $GOPATH/src/github.com/sharath/n-queens`

`go build`

`./main`

#### Specifying Initial States
The positions are represented with just the row numbers of each queen.
Just modify 8queens.json or make a new json file following that format and call
`./main <file.json>`

#### Known Issues
* Doesn't support n > 9