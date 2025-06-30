
import java.util.Scanner;




   //String.



public class Main {
    public static void main(String[] args) {



        //String

        //A seguir observa esse código de string:

      /* //Declaração de variável
        String frase = "Tecnologia é para todos!";

        System.out.println("\n"+"\n");

        //Imprimir a frase
        System.out.println(frase); */


        //Realiza de outra forma esse código com entrada de dado e saída de dado.

       //Destinado para entrada de dado e saída de dado

        Scanner ler = new Scanner(System.in);

        System.out.println("\n"+"\n");

        //Entrada de dado
        System.out.println("Informe  a frase\n");
        //Saída de dado
        String frase=ler.nextLine();

        System.out.println("\n"+"\n");

        //Imprimir o cálcaulo da mensagem
       System.out.println(frase);




    }
}
