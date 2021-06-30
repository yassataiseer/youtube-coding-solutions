#include <iostream>
#include <vector>
#include <string>
using namespace std;
//Run cmnds below:
//g++ -std=c++14 main.cpp -o main
//./main


int main()
{
    vector <int>new_vector = {1,2,3,4,5};
    vector <int>vector_2(6,18);
    cout << new_vector.size()<<endl;
    new_vector.push_back(6);
    cout << new_vector.at(5)<<endl;
    cout << new_vector.size()<<endl;

    return 0;
}