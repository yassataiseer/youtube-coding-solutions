using System;

namespace MyApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Cars NewCar1 = new Cars("Monda Civic",2020,"Black");
            Console.WriteLine(NewCar1.Name);
        }
    }
    class Cars{
        public string Name{get;set;}
        public int Year{get;set;}
        public string Color{get;set;}
        public Cars(string CarName,int YearMade,string CarColor){
            Name = CarName;
            Year = YearMade;
            Color = CarColor ;
        }
    }


}
