// #1
//#include <iostream>
//#include <string>
//using namespace std;


//string compress(string s) {
    //string result;          
    //int count = 1;         
    
    
    //for (int i = 0; i < s.length(); i++) {
        
        //if (s[i] == s[i + 1]) {
            //count++;       
        //} else {
            
            //result += s[i] + to_string(count);
            //count = 1;      
        //}
    //}
    
    
    //if (result.length() < s.length()) {
       // return result;     
    //}
   // return s;              
//}

//int main() {
   // string input;
   // cout << "Введите строку: ";
   // cin >> input;                           
   // cout << compress(input) << endl;        
  //  return 0;
//}
// #2 

//#include <iostream>
//#include <string>
//#include <vector>
//using namespace std;

//string isPangram(string s) {
    //vector<bool> letters(26, false);
    //int count = 0;
    
    //for (char c : s) {
        //if (c >= 'a' && c <= 'z') {
            //if (!letters[c - 'a']) {
                //letters[c - 'a'] = true;
                //count++;
            //}
        //}
    //}
    
    //if (count == 26) {
        //return "Yes";
    //}
    //return "No";
//}

//int main() {
    //string input;
    //cout << "Введите строку: ";
    //getline(cin, input);
    //cout << isPangram(input) << endl;
    //return 0;
//}


//#3

#include <iostream>
using namespace std;

int main() {
    int n, x;
    cin >> n;
    
    for (int i = 0; i < n; i++) {
        cin >> x;
        
        
        if (x < 10) {
            continue;
        }
        
        int lastDigit = x % 10;
        int temp = x;
        bool good = true;
        
        while (temp > 0) {
            if (temp % 10 != lastDigit) {
                good = false;
                break;
            }
            temp = temp / 10;
        }
        
        if (good) {
            cout << x << " ";
        }
    }
    return 0;
}


