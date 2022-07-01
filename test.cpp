// program below illustrates the 
// vector::insert() function 
#include <iostream> 
#include <vector>
#include <iterator>
using namespace std; 
  
int main() 
{ 
    // initialising the vector 
    vector<int> vec1; 
    vector<int>::iterator i_vec1;
    vector<int> vec2; 
    vec1.push_back(10);
    vec1.push_back(20);
    vec1.push_back(30);
    vec1.push_back(40);
  
    // inserts at the beginning of vec2 
    vec2.insert(vec2.begin(), vec1.begin(), vec1.end()); 

    vec1[0] = 2;
  
    cout << "The vector2 elements are: "; 
    for (i_vec1 = vec2.begin(); i_vec1 != vec2.end(); ++i_vec1) 
        cout << *i_vec1 << " "; 
    
    cout << endl << vec1[0];
  
    return 0; 
}