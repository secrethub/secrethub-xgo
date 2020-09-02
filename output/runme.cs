using System;
 public class runme {
     static void Main() {
        SecretVersion secret = secrethub.Read("secrethub-xgo/dotnet/test-secret");
        Console.WriteLine("Version ID: {0}", secret.SecretVersionID);
        Console.WriteLine("Version: {0}", secret.Version);
        Console.WriteLine("Data: {0}", secret.Data);
        Console.WriteLine("Created At: {0}", secret.CreatedAt);
        Console.WriteLine("Status: {0}", secret.Status);
         try {
             Console.WriteLine(secrethub.ReadString("secrethub-xgo/dotnet/test-secret"));
             Console.WriteLine(secrethub.ReadString("secrethub-xgo/dotnet/non-existent-secret"));
         } catch (Exception e) {
             Console.WriteLine("Exception caught: {0}", e);
        }
     }
 }
