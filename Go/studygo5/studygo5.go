
package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)
 
func main() { //Função prinicipal do programa 

     //Declaração da variável

       var dia_consulta int=18;
    

     //Estrutura de condição(validação)

     if dia_consulta==18{

     //Imprimir a mensagem no console(print)
     fmt.Println("Hoje é o dia da consulta por volta das 16:00!");

         
     }else if dia_consulta==19{

     //Imprimir a mensagem no console(print)
     fmt.Println("Hoje não é dia da consulta!");
  
        
    }else{


     //Imprimir a mensagem no console(print)
     fmt.Println("Não encontrado o dia da consulta no sistema!");
  
       
   }
    
   
    
    
    
    
}
