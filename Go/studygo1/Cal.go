package main //Ponto de entrada 

import "fmt" //format(formatação para imprimir o println)

func main() { //Função prinicipal do programa 

    //Declaração de variáveis
    
    var num1 int =5;
    var num2 int =5;
    var num3 int=2;
    var cal int;

    //Cálculo 

    cal= num1+num2-num3;
    
    //Imprimir a mensagem no console(print) e o resultado
    //do cálculo
    
    fmt.Println("O resultado é=",cal) 
    
}
