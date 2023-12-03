startTimer = new Date();
const input = require('fs').readFileSync('input.txt', 'utf8').split('\n');

let sum = 0;
for (let i = 0; i < input.length; i++) {
    let left, right
    for (let c = 0; c < input[i].length; c++) {
        if (input[i][c] >= '0' && input[i][c] <= '9') {
            if (left === undefined) {
                left = input[i][c];
                right = input[i][c];
            } else {
                right = input[i][c];
            }
        }

    }
    sum += parseInt(left) * 10 + parseInt(right);
}
console.log(sum);
endTimer = new Date();
console.log(endTimer - startTimer + " ms");