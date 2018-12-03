package main
import "fmt"

func main(){
  var n int
  fmt.Println("Enter the limit:")
  fmt.Scanln(&n)
  num1 := 0
  num2 := 1
  for i := 0; i<n; i++{
    if i == 0{
      fmt.Println(num1)
    }else if i == 1{
      fmt.Println(num2)
    }else {
      new := num1 + num2
      fmt.Println(new)
      num1 = num2
      num2 = new
    }
  }
}
