package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strconv"
)

func max_file(files *[]os.FileInfo) (max int) {
    max = 0
    for _, file := range *files {
//            fmt.Println(file.Size())
            i64 := strconv.FormatInt(file.Size(), 10)
            if len(i64) >= max {
                max = len(i64)
            }
    }
    return
}

func main() {

    dirname := "." + string(filepath.Separator)

    d, err := os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()

    files, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    max := max_file(&files)
//    fmt.Println(max)
    for _, file := range files {
        fmt.Print(file.Mode().String(), " ")
        i64 := strconv.FormatInt(file.Size(), 10)
        for i := 0; i < max - len(i64); i++ {
            fmt.Printf(" ")
        }
        fmt.Print(file.Size())
        fmt.Println(" ", file.Name())
        fmt.Println(" ", file.Sys())
    }
}
