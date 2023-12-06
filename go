package main
import "fmt"

func main() {
  var a int
  var b int
  var operator string
  var result float64
  
  fmt.Print("Enter a operator: ")
  fmt.Scan(&operator)
  
  switch operator{
      case "+":
      fmt.Print("Enter a Number: ")
      fmt.Scan(&a)
      fmt.Print("Enter b Number: ")
      fmt.Scan(&b)
      result = float64(a + b)
      
      case "-":
      fmt.Print("Enter a Number: ")
      fmt.Scan(&a)
      fmt.Print("Enter b Number: ")
      fmt.Scan(&b)
      result = float64(a - b)
      
      case "*":
      fmt.Print("Enter a Number: ")
      fmt.Scan(&a)
      fmt.Print("Enter b Number: ")
      fmt.Scan(&b)
      result = float64(a * b)
      
      case "/":
      fmt.Print("Enter a Number: ")
      fmt.Scan(&a)
      fmt.Print("Enter b Number: ")
      fmt.Scan(&b)
      result = float64(a) / float64(b)
      
      if b != 0 {
			result = float64(a) / float64(b)
		} else {
			fmt.Println("Error: Division by zero.")
			return
		}
	default:
		fmt.Println("Enter a valid operator.")
		return
      
  }
  fmt.Printf("%.2f %s %.2f = %.2f", float64(a), operator, float64(b), result)
}
