using System;
 public class runme {
     static void Main() {
        Secret secret = secrethub.Read("secrethub-xgo/dotnet/test-secret");
        Console.WriteLine("Version: {0}", secret.Version);
        Console.WriteLine("Data: {0}", secret.Data);
         try {
             Console.WriteLine(secrethub.ReadString("secrethub-xgo/dotnet/test-secret"));
             Console.WriteLine(secrethub.ReadString("secrethub-xgo/dotnet/non-existent-secret"));
         } catch (Exception e) {
             Console.WriteLine("Exception caught: {0}", e);
        }
     }
 }
