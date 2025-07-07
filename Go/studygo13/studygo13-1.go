
//Realiza um código para demonstrar uma string de letras com números
//utilizando operador ou(||) e identifica com mensagem que foi
//encontrado no Banco de dado ou não foi encontrado.



package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)
 
func main() { //Função principal do programa 

     //Declaração da variável 

       cod_identificao := "IA589n"; 
 
     //Estrutura de condição(validação) com(ou ||)


         
     if cod_identificao=="IA589n" || cod_identificao=="tr1Ubq"{

     //Imprimir a mensagem no console(print)
     fmt.Println("Encontrado o código no Banco de dado!");
  

    }else{

     //Imprimir a mensagem no console(print)
     fmt.Println("Obteve um erro na execução no sistema ou não encontrado!");
  
       
       }
    }

