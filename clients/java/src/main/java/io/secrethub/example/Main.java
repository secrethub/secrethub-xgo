package io.secrethub.example;

import io.secrethub.SecretHubClient;

public class Main {
    static public void main(String argv[]) {
        // Create a new client, sourcing configuration from the environment.
        SecretHubClient client = new SecretHubClient();
    
        // Source the path from the command-line arguments
        String path = argv[0];

        // Use the client
        System.out.println("The secret "+path);

        try {
            System.out.println("read: "+client.read(path));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}