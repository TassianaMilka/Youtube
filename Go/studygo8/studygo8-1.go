package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)
 
func main() { //Função prinicipal do programa 

     //Declaração das variáveis 
      
       sigla := "B";
    
       var valor float64=300;
 
     //Estrutura de condição(validação) com(ou ||)

     if sigla=="A" || sigla=="B"{

     //Imprimir a mensagem no console(print)
     fmt.Println("A sala (a e b) obteve a maior média de pontuação!");

         
    } 
    if sigla=="C" || sigla=="D"{

     //Imprimir a mensagem no console(print)
     fmt.Println("A sala (c e d) obteve uma regular média de pontuação!");

}

     //Estrutura de condição(validação) com(e &&)

     if (valor>=200 && valor<=500){

     //Imprimir a mensagem no console(print)
     fmt.Println("O valor está entre R$200.00 até R$500.00!");

    }else{


     //Imprimir a mensagem no console(print)
     fmt.Println("Obteve um erro no sistema!");
     
        }   
   
    }
