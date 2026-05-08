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

