#include <iostream>
#include <string>
using namespace std;


string compress(string s) {
    string result;          
    int count = 1;         
    
    
    for (int i = 0; i < s.length(); i++) {
        
        if (s[i] == s[i + 1]) {
            count++;       
        } else {
            
            result += s[i] + to_string(count);
            count = 1;      
        }
    }
    
    
    if (result.length() < s.length()) {
        return result;     
    }
    return s;              
}

int main() {
    string input;
    cout << "Введите строку: ";
    cin >> input;                           
    cout << compress(input) << endl;        
    return 0;
}

  

