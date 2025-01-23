package principal;

import modelo.Modelo;

public class Principal {


    public static void main(String[] args) {



        Modelo modelo=new Modelo();

        modelo.setNome("Tassiana");
        modelo.setPais("Brasil");
        modelo.setArea("Tecnologia");
        modelo.setTelefone("20102-1456");


        System.out.println(modelo.getNome()+"\n"+modelo.getPais()+"\n"+modelo.getArea()+"\n"+modelo.getTelefone());


    }


}