The following command can be used to initialize a new go module.
<br>
`$ go mod init [module-name]`


The following command can be used to build the go program into binaries.
<br>
`$ go build [go-file-name]`


To execute the binary, use `$ ./[binary-name]`.


The following command can be used to run the go program.
<br>
`$ go run [go-file-name]`


To add a new module requirements & sums, use the following command. Go will add the external module as a requirement, as well as a go.sum file for use in authenticating the module.
<br>
`$ go mod tidy`


The following command can be used to run the dependencies along with main program.
<br>
`$ go run .`