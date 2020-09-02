using System;
 public class runme {
     static void Main() {
         try {
             Console.WriteLine(secrethub.Read("marton/demo/username"));
             Console.WriteLine(secrethub.Read("marton/asfdasjfk/asdfasdf"));
         } catch (Exception e) {
             Console.WriteLine("Exception caught: {0}", e);
        }
     }
 }
