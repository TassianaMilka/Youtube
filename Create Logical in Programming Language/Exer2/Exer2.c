

//Fruit Ninja== Halfbrick Studios
//Lançado em 2010
//Principal objetivo é cortar frutas
 



//Biblioteca
#include <stdio.h>

int main()
{ 
    
    //Declaração de variáveis
    
    int num;
    int ano;
    int num2;

    
    
    printf("Digite quem é a empresa criadora do jogo Fruit Ninja para[1]-Gamenow ou [2]-Halfbrick Studios= ");//Entrada de dado
    scanf("%d",&num);//Saída de dado
    printf("\n\n");
    
   //Validação
   if(num==1){
       
       //Imprimir uma mensagem
       
        printf("Resposta errada");
        printf("\n\n");
       
   }
    else if(num==2){
       
        //Imprimir uma mensagem
       
        printf("Resposta verdadeira");
        printf("\n\n");
       
   }
    
    printf("Digite o ano que foi criadora o jogo Fruit Ninja em 2010 ou 2006=");//Entrada de dado
    scanf("%d",&ano);//Saída de dado
    printf("\n\n");
    
   //Validação
   if(ano==2010){
       
       //Imprimir uma mensagem
       
        printf("Resposta verdadeira");
        printf("\n\n");
       
   }
    else if(ano==2006){
       
        //Imprimir uma mensagem
       
        printf("Resposta errada");
        printf("\n\n");
       
       
   }
   
   
   printf("Digite a principal funcionalidade do jogo Fruit Ninja para [3]-Jogo de carro ou [4]-Jogo de cortar frutas =");//Entrada de dado
   scanf("%d",&num2);//Saída de dado
    printf("\n\n"); 
    
   //Validação
   if(num2==3){
       
       //Imprimir uma mensagem
       
        printf("Resposta errada");
        printf("\n\n");
       
   }
    else if(num2==4){
       
        //Imprimir uma mensagem
       
        printf("Resposta verdadeira");
        printf("\n\n");
   }
   
     //Iprimir a mensagem final
    
   printf("Finalização  do Jogo de lógica sobre o Fruit Ninja =)");
   
    

    return 0;
}
