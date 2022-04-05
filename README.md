## Learning Go - Cards Project 

This is repository to learn [golang](https://go.dev/doc/), with follow along project from [udemy course](https://www.udemy.com/course/go-the-complete-developers-guide/)

#### Packages Tried 
- ioutils
    - method: `WriteFile`
        -  `func WriteFile(filename string, data []byte, perm fs.FileMode) error` 
        - [go reference](https://cs.opensource.google/go/go/+/go1.18:src/io/ioutil/ioutil.go;l=45)
    - method: `ReadFile`
        - `func ReadFile(filename string) ([]byte, error)`
        - [go reference](https://cs.opensource.google/go/go/+/go1.18:src/io/ioutil/ioutil.go;l=36)
- strings
    - method: `Join`
        - `func Join(elems []string, sep string) string`
        - [go reference](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/strings/strings.go)
    - method: `Split`
        - `func Split(s string, sep string) []string`
        - 
- os
    - method: `Exit`
    - `func Exit(code int)`
    - [go reference](https://pkg.go.dev/os#Exit)
