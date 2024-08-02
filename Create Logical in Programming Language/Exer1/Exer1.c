//Biblioteca
#include <stdio.h>

int main()
{ 
    
    //Declaração de variáveis
    
    int num=10;
    
    
    printf("Digite o número");//Entrada de dado
    scanf("%d",&num);//Saída de dado
    
   //Validação
   if(num==10){
       
       //Imprimir uma mensagem
       
        printf("O número  10 não é raiz quadrada");
        printf("\n\n");
       
   }
   /*else if(num<=9){
       
        //Imprimir uma mensagem
       
        printf("O número  10  é raiz quadrada");
       
       
   }*/
    
    //Declaração de variáveis
    int soma=0;
    int subtracao=100;
    int multiplicacao;
    int divisao;
    
    //Realizar os  cálculos das operações
    soma=5+5;
    subtracao=10-5;
    multiplicacao=2*5;
    divisao=10/2;
    
    //Imprimir o resultado
    printf("%d=",soma);
    printf("%d= ",subtracao);
    printf("%d=", multiplicacao);
    printf("%d=",divisao);
    

    return 0;
}