using System;

namespace c_tutorial
{
    class Program
    {
        static int Add(int num1,int num2){
            return num1+num2;
        }
        static void Main(string[] args)
        {
            int sum = Add(40,60);
            Console.WriteLine(sum);
        }
    }
}
