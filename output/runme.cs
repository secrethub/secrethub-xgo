using System;
 public class runme {
     static void Main() {
         try {
             secrethub.Write("horiaculea/flags/infusedtea", "colazero");
             Console.WriteLine(secrethub.Exists("marton/asfdasjfk/asdfasdf"));
         } catch (Exception e) {
             Console.WriteLine("Exception caught: {0}", e);
        }
     }
 }
