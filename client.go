package main
import (
	"net/rpc"
	"fmt"
	"os"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Q, R int
}

func main () {
	
	service := "localhost:13232"
	client, err := rpc.Dial("tcp", service)
	
	defer client.Close()
	
	checkError("Dial: ", err)
	
	fmt.Println("* - multiplicação")
	fmt.Println("/ - divisão")
	fmt.Println("+ - soma")
	fmt.Println("- - subtracao")
	
	var op byte
	fmt.Scanf("%c\n", &op)
	
	switch op {
	/* código para implementar os
	procedimentos disponíveis */
		case '*':
			args := readArgs()
			var reply int
			mulCall := client.Go("Arith.Multiply", args, &reply, nil)
			fmt.Println("Esperando servidor..")
			replyMul := <- mulCall.Done
			checkError("Multiply: ", replyMul.Error)
			fmt.Printf("%d * %d = %d\n",
			args.A, args.B, reply)
			os.Exit(0)
		case '/':
			args := readArgs()
			var reply Quotient
			divCall := client.Go("Arith.Divide", args, &reply, nil)
			replyDiv := <- divCall.Done
			checkError("Divide: ", replyDiv.Error)
			fmt.Printf("%d / %d = (%d,%d)\n",
			args.A, args.B, reply.Q, reply.R)
			os.Exit(0)
		case '+':
			args := readArgs()
			var reply int
			err = client.Call("Arith.Sum", args, &reply)
			checkError("Sum: ", err)
			fmt.Printf("%d + %d = %d\n",
			args.A, args.B, reply)
			os.Exit(0)
		case '-':
			args := readArgs()
			var reply int
			err = client.Call("Arith.Minus", args, &reply)
			checkError("Minus: ", err)
			fmt.Printf("%d - %d = %d\n",
			args.A, args.B, reply)
			os.Exit(0)
		default:
			fmt.Println("Opção inválida: ", op)
			os.Exit(1)
	}
}

func readArgs () Args {
	var a, b int
	fmt.Println("A: ")
	fmt.Scanln(&a)
	fmt.Println("B: ")
	fmt.Scanln(&b)
	return Args{a, b}
}

func checkError (str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}