import java.util.Scanner;


   //Double;

   //String.

 
 
public class Main {
    public static void main(String[] args) {


        //Realiza um código destinado para saber a quantidade exata que um cliente contém na sua conta do banco.Lembrando que possuí R$200,00,adicionou
        //R$4000 e retirou R$100,00.Qual é a quantidade final?


       //Destinado para entrada de dado e saída de dado

        Scanner ler = new Scanner(System.in);

        System.out.println("\n"+"\n");

        //Entrada de dado
        System.out.println("Informe quantos o cliente contém no banco\n");
        //Saída de dado
        double valor1=ler.nextDouble();

        //Entrada de dado
        System.out.println("Informe quantos o cliente  adicionou do banco\n");
        //Saída de dado
        double valor2=ler.nextDouble();

        //Entrada de dado
        System.out.println("Informe quantos o cliente retirou no banco\n");
        //Saída de dado
        double valor3=ler.nextDouble();

        //Realização do cálculo

       double calculo=valor1+valor2-valor3;

       //Imprimir o cálcaulo da mensagem
       System.out.println("O resultado do valor da conta R$"+calculo);


        //String

        //A seguir observa esse código de string:

      /* //Declaração de variável
        String frase = "Tecnologia é para todos!";

        System.out.println("\n"+"\n");

        //Imprimir a frase
        System.out.println(frase); */


        //Realiza de outra forma esse código com entrada de dado e saída de dado.


    }
}

