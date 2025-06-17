
//Faça um código que cálcula quantos de dinheiro vai conter depois do 
//pagamento da internet que é R$80,00 e energia de R$100,000 do valor exato 
//de R$980,00.

package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)


func main() { //Função principal do programa

    //Declaração de variáveis
    var calculoP float64=0;
    var valorP float64=80;
    var valorS float64=100;
    var calculoS float64=0;
    var valortotal float64=980;

    //Cálculos 

    calculoP=valorP+valorS;
    calculoS=valortotal-calculoP;

    //Imprimir a mensagem no console(print) e valor
    
    fmt.Println("o valor do resultado:",calculoS);

}
    
