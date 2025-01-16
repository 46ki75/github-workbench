package org.example;

public class App {
    public String getGreeting() {
        return "Hello World! (from lib)";
    }

    public static void main(String[] args) {
        System.out.println(new App().getGreeting());
    }
}
