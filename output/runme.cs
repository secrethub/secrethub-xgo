using System;
 public class runme {
     static void Main() {
        SecretVersion secret = secrethub.Read("teddy2008/demo/username");
        Console.WriteLine("Version ID: {0}", secret.SecretVersionID);
        Console.WriteLine("Secret ID: {0}", secret.Secret.SecretID);
        Console.WriteLine("Version: {0}", secret.Version);
        Console.WriteLine("Data: {0}", secret.Data);
        Console.WriteLine("Created At: {0}", secret.CreatedAt);
        Console.WriteLine("Secret Created At: {0}", secret.Secret.CreatedAt);
        Console.WriteLine("Status: {0}", secret.Status);
         try {
             Console.WriteLine(secrethub.ReadString("teddy2008/demo/username"));
             Console.WriteLine(secrethub.ReadString("teddy2008/demo/non-existent-secret"));
         } catch (Exception e) {
             Console.WriteLine("Exception caught: {0}", e);
        }
     }
 }
