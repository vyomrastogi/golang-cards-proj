## Learning Go - Cards Project 

This is repository to learn [golang](https://go.dev/doc/), with follow along project from [udemy course](https://www.udemy.com/course/go-the-complete-developers-guide/)

#### Packages used in project 
- __ioutils__
    - method: `WriteFile`
        -  `func WriteFile(filename string, data []byte, perm fs.FileMode) error` 
        - [go reference](https://cs.opensource.google/go/go/+/go1.18:src/io/ioutil/ioutil.go;l=45)
    - method: `ReadFile`
        - `func ReadFile(filename string) ([]byte, error)`
        - [go reference](https://cs.opensource.google/go/go/+/go1.18:src/io/ioutil/ioutil.go;l=36)
- __strings__
    - method: `Join`
        - `func Join(elems []string, sep string) string`
        - [go reference](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/strings/strings.go)
    - method: `Split`
        - `func Split(s string, sep string) []string`
        - 
- __os__
    - method: `Exit`
    - `func Exit(code int)`
    - [go reference](https://pkg.go.dev/os#Exit)
- __math/rand__
    - method: `Intn`
    - ` func Intn(n int) int ` : this function always uses same seed value to generate random number
    - [go reference](https://pkg.go.dev/math/rand@go1.18#Intn)
- __time__
    - method: `Now().UnixNano()`


