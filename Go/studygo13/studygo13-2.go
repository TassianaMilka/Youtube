//Descreve para encontrar na utilização de operadores e(&&) que possa conter conforme
//esses números 2,4,1,3,5,7,6,8 e 4,7,1,5,2,8,3,6 na identificação da mensagem
//Os números estão aleatórios! e para declarar utilizando 1,2,3,4,5,6,7,8 de uma
//definição Os números estão em ordem!.


//Explicação no vídeo

//1
//9
//3

//1,2,3,4,5,6,7,8
//9,7,6,5,4,3,2,1
//2,4,1,3,5,7,6,8
//4,7,1,5,2,8,3,6


package main //Ponto de entrada 

import "fmt" //format(formatação para imprimir o println)


func main() { //Função principal do programa 

     //Declaração da variável 

      var numeros int=12345678;
  

     //Estrutura de condição(validação) com(e &&)

     if (numeros>=471528360 && numeros<=24135768){

     //Imprimir a mensagem no console(print)
     fmt.Println("Os números estão aleatórios!");

    }else{


     //Imprimir a mensagem no console(print)
     fmt.Println("Os números estão em ordem!");

          

         }
     }

