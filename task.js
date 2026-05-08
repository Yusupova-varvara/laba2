const fs = require('fs');
const input = fs.readFileSync(0, 'utf8').trim().split(/\s+/);

let index = 0;
const n = parseInt(input[index++]);

for (let i = 0; i < n; i++) {
    const x = parseInt(input[index++]);
    
    if (x < 10) {
        continue;
    }
    
    const lastDigit = x % 10;
    let temp = x;
    let good = true;
    
    while (temp > 0) {
        if (temp % 10 !== lastDigit) {
            good = false;
            break;
        }
        temp = Math.floor(temp / 10);
    }
    
    if (good) {
        process.stdout.write(x + " ");
    }
}
console.log();


