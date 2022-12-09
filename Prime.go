package main

import "fmt"
func main(){
    fmt.Print("Enter the number:")
    var n int
    fmt.Scanln(&n)
    m:=n/2
    flag:=0
    if n==0||n==1{
        fmt.Print(n," is not prime number")
    }else{
        for i:=2;i<=m;i++{
        if n%i==0{
            fmt.Print(n," is not prime number")
            flag=1
            break
        }
    }
    if flag==0{
        fmt.Print(n," is prime number")
    
        
    }
    
    
    }
}