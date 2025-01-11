//Começo
#include <stdio.h>

int main()
{
    
  //Declaração
    
    int num1;
    int num2;
    int resultado;
    
    
    //
    printf("Informe a primeira nota\n");//Entrada de dado
    scanf("%d",&num1);//Saída de dado
    
    printf("Informe a segunda nota\n");//Entrada de dado
    scanf("%d",&num2);//Saída de dado
    
    //Cálculo das notas
    resultado=num1+num2/2;
    
    
    //Imprimir o resultado das notas
    printf("O resultado %d\n",resultado);
    
    
    //Validação maior o menor(aprovado ou reprovado)
    if(resultado>=6){
    //Imprimir a mensagem    
    printf("\nO aluno foi aprovado");

        
    }else{

    //Imprimir a mensagem    
     printf("\nO aluno foi reprovado");

        
    }
    
    //Fim
    return 0;
}


