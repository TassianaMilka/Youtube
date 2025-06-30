 //Uma escola realiza um evento dos alunos para praticar esporte
 // e cada um contém a sua númeração de identificação na
//sua camiseta,mas uma aluna precisa de uma camiseta do número
//89 com isso faça um código de validação descrevendo a frase
// que possui uma camiseta com essa numeração.

package main //Ponto de entrada  

import "fmt" //format(formatação para imprimir o println)
 
func main() { //Função principal do programa 

     //Declaração da variável

       var numeracao_Camiseta int=89;
    

     //Estrutura de condição(validação)

     if numeracao_Camiseta==89{

     //Imprimir a mensagem no console(print)
     fmt.Println("Contém uma camiseta do número 89!");
         
           
    }else{


     //Imprimir a mensagem no console(print)
     fmt.Println("Não contém a numeração da camiseta!");
  
       
   }
