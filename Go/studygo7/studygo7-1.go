package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)
 
func main() { //Função prinicipal do programa 

     //Declaração da variável 

       sigla := "B"; 

     //Estrutura de condição(validação) com(ou ||)

     if sigla=="A" || sigla=="B"{

     //Imprimir a mensagem no console(print)
     fmt.Println("A sala (a e b) obteve a maior média de pontuação!");

         
    }else if sigla=="C" || sigla=="D"{

     //Imprimir a mensagem no console(print)
     fmt.Println("A sala (c e d) obteve uma regular média de pontuação!");
  

    }else{


     //Imprimir a mensagem no console(print)
     fmt.Println("Obteve um erro no sistema!");
  
       
   }
    }
