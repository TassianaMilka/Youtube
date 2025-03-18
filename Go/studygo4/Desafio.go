//Desafio

//Criar um código que a sua principal funcionalidade é imprimir
//um nome de um produto e que possa realizar um cálculo de 100 convertido para
// 0,1(juro)de uma subtração de R$900.00 para retirar R$700.00.
//Conversão(100=0,1).

package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)

func main() { //Função prinicipal do programa 

     //Declaração das variáveis

     produtoNome:="Celular";
    
     var cal float64=900.00-700.00*0.1;
    
     //Imprimir o nome do produto no console(print)
     fmt.Println(produtoNome);
    
     //Imprimir o nome do produto no console(print)
    
     fmt.Println(cal);
    
    
}
