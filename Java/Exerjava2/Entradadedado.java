import java.util.Scanner;

// Entrada de dado e saída;

//Diferente forma de declaração de variável


public class Main {
    public static void main(String[] args) {

       //Destinado para entrada de dado e saída de dado

        Scanner ler = new Scanner(System.in);

       //Declaração de variável

        int num;

        System.out.println("\n"+"\n");

        //Entrada de dado
        System.out.println("Informe o número");
        //Saída de dado
        num=ler.nextInt();

        //Imprimir a mensagem e o número informado pelo usuário

        System.out.println("O número informado pelo usuário é="+num);
    }
}



       /* Double valor =200.00;

        System.out.println("\n"+"\n");

        System.out.println("O valor informado pelo usuário é="+valor);*/


   //Realiza um código destinado para saber a quantidade exata que um cliente contém na sua conta do banco.Lembrando que possuí R$200,00,adicionou
   //R$4000 e retirou R$100,00.Qual é a quantidade final?
