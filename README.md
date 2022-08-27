## Dynamic Yaml Parser Helper

## Get Common Values

syntax is very simple, path is separated by ".",when there is an array,using "[]" to index.

```go

func main() {

    // result:=ReadString("metadata.name","test.yaml")
    result:=ReadString("spec.template.spec.containers[0].name","test.yaml")
    fmt.Println(result)
    
    result2:=ReadInt("spec.replica","test.yaml")
    fmt.Println(result2)

}
```


## Build & Develop

```bash
go mod tidy
go run main.go
```

