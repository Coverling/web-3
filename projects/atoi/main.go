package main

import (
    "fmt"
    "strconv"
)

func main() {
    var a string
    fmt.Scan(&a)

    // Преобразуем строку в целое число
    num, err := strconv.Atoi(a)
    if err != nil {
        fmt.Println("Ошибка преобразования:", err)
        return
    }

    var strNewNum string
    b := len(a)

    for i := 0; i < b; i++ {
        c := num % 10
        c *= c
        strNewNum = strconv.Itoa(c) + strNewNum // Преобразуем квадрат в строку и добавляем к результату
        num /= 10
    }
    newnum, err := strconv.Atoi(strNewNum)
    if err != nil {
        fmt.Println("Ошибка преобразования:", err)
        return
    }
    fmt.Println(newnum)
}

