package main  
import "fmt" 

//Placeholders for user's selection, and their chosen numbers
var s string  
var num1 int  
var num2 int  

func promptNums(){
	//Prompt user to enter two chosen numbers for the arithmetic operation
	fmt.Println("Select the first number:") 
	fmt.Scanf("%d", &num1) 
	fmt.Println("Select the second number:") 
	fmt.Scanf("%d", &num2) 
}

func main(){ 
	loop: 
		for { 
			//Print out instructions for the user
			fmt.Println("Insert your selection:") 
			fmt.Println("1. Addition") 
			fmt.Println("2. Subtraction") 
			fmt.Println("3. Division") 
			fmt.Println("4. Multiplication") 

			//Capture user's arithmetic operation choice
			fmt.Scanf("%s", &s) 

			switch s { 
			case "1": 
				//Addition
				promptNums()
				res := num1 + num2
				fmt.Printf("%d + %d = %d" ,num1,num2,res) 
				break loop 
			case "2":
				//Subtraction
				promptNums()
				res := num1 - num2
				fmt.Printf("%d - %d = %d" ,num1,num2,res) 
				break loop 
			case "3":
				//Division
				promptNums()
				res := num1 / num2
				fmt.Printf("%d / %d = %d" ,num1,num2,res) 
				break loop 
			case "4":
				//Multiplication
				promptNums()
				res := num1 * num2
				fmt.Printf("%d * %d = %d" ,num1,num2,res) 
				break loop 
			default:
				//Any other unknown operations  
				fmt.Println("Unknown number, try again.") 
				continue loop  
		} 
		} 
} 
