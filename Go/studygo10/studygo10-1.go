
//Faça um código para saber sobre a idade de um aluno que contém 23 anos e quer saber
//quantos anos vai conter daqui 6 anos.Depois a professora quer ensinar os alunos o cálculo
// alunos para saber o ano de nascimento.

package main //Ponto de entrada 

import "fmt" //format(formatação para imprimir o println)

func main() { //Função principal do programa 

    //Declaração de variáveis
    
     var idadeS int=0;
     var idade int =23;
     var anoN int=0;
     var anoA int=2025;
    
    //Cálculos  

    idadeS=idade+6;
    anoN=anoA-idade;
    
    //Imprimir a mensagem no console(print) e o resultado
    //do cálculo da idade e ano de nascimento
    
    fmt.Println("A idade do aluno:",idadeS); 
    
    fmt.Println("O ano de nascimento do aluno:",anoN);

    
}

