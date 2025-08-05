//Faça um código destinado para uma passagem de são paulo para campinas utilizando string com vetor/array do cálculo
// de ida e volta para finalizar sem o vetor com outro valor de subtração do total 10.

package main // Ponto de entrada


import "fmt" // Importando o pacote fmt para impressão


func main() {  // Função principal do programa

    //Declaração de vetor/Array
    var nome_cidade =[2]string {"Campinas","São Paulo"};

    vetor1 := [5]int{20, 40, 50, 70, 90}; 
    vetor2 := [5]int{20, 40, 60, 80, 90};
    var resultado [1]int;
    //Declaração e cálculo
    var resultado_num int = 150-10;


    //Cálculo
    resultado[0]= vetor1[4] +vetor2[2];
    

    //Imprimir o nome a string
    fmt.Println(nome_cidade);

    //Imprimir o resultado do vetor/array
    fmt.Println(resultado);

    //Imprimir o resultado sem vetor/array
    fmt.Println(resultado_num);

}
