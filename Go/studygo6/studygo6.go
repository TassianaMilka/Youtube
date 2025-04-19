package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)
 
func main() { //Função prinicipal do programa 

     //Declaração da variável 

       sigla := "NORC";

    //Norte,Nordeste,Centro-Oeste,Sudeste e Sul

   // NORC
   // SDES

     //Estrutura de condição(validação)

     if sigla=="NORC"{

     //Imprimir a mensagem no console(print)
     fmt.Println("A passagem comprada é para Região Norte,Nordeste ou Centro-Oeste do Brasil!");

         
     }else if sigla=="SDES"{

     //Imprimir a mensagem no console(print)
     fmt.Println("A passagem comprada é para Região Sudeste e Sul do Brasil!");
  
        
    }else{


     //Imprimir a mensagem no console(print)
     fmt.Println("Obteve um erro no sistema!");
  
       
   }

}
    
