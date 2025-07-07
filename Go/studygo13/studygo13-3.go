//Imprima três  vezes sobre duas mensagens conforme a descrição:
//Todas linguagens de programação contém as suas finalidades no mercado de trabalho!
//Para aprimorar o conhecimento é necessário querer aprender!


package main // Ponto de entrada

import "fmt" // Importando o pacote fmt para impressão

func main() { // Função principal do programa

    // Declaração de variável 
    var i int 

    // Loop(Estrutura de repetição)
    
    for i = 0; i < 3; i++ {
        // Imprimir as mensagens no console(print)
        fmt.Println("Todas linguagens de programação contém as suas finalidades no mercado de trabalho!",i);
        fmt.Println("Para aprimorar o conhecimento é necessário querer aprender!",i);

       
    }


    
}

