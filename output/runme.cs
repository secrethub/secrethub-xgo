using System;
 public class runme {
     static void Main() {
         try {
             Console.WriteLine(secrethub.ReadString("marton/demo/username"));
             Console.WriteLine(secrethub.ReadString("marton/asfdasjfk/asdfasdf"));
         } catch (Exception e) {
             Console.WriteLine("Exception caught: {0}", e);
        }
     }
 }
