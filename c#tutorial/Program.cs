using System;

namespace c_tutorial
{
    class Program
    {

        static void Main(string[] args)
        {
            int num = add(1,3);
            Console.WriteLine(num);
            //Output: 4
        }
        static int add(int num1,int num2){//add function
            //return an int
            //take arguments num1, num2
            return num1+num2;//Value returned to main function
        }
    }
}
