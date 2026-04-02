To get started, download the installation files from **golang.org/dl/** and follow the instructions for your operating system. You can verify the installation by running the version command in your terminal.

* **Check Installation:** `go version`
* **IDE Setup:** Download VS Code, install the "Go" extension by Google, and use the Command Palette to run "Go: Install/Update Tools."
* **Initialization:** Run `go mod init example.com/hello` in your terminal to prepare your project.



Create a file named **helloworld.go** and paste the following code:

```go
package main
import ("fmt")

func main() { 
    fmt.Println("Hello World!") 
} 
```

Once the file is saved, use the terminal to execute or compile your program:

* **To Run:** `go run .\helloworld.go` (Output: Hello World!)
* **To Build:** `go build .\helloworld.go` (Saves as an executable)