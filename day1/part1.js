const fs = require('fs');

// open input.txt into an array of strings
const input = fs.readFileSync('input.txt', 'utf8').split('\n');

let sum = 0;

// loop through each line of input
for (let i = 0; i < input.length; i++) {
    let left, right
    for (let c = 0; c < input[i].length; c++) {
        // if the current character is a number
        if (input[i][c] >= '0' && input[i][c] <= '9') {
            // if the left number has not been set
            if (left === undefined) {
                // set the left number
                left = input[i][c];
            } else {
                // set the right number
                right = input[i][c];
            }
        }

    }

    // if right is undefined, set it to left
    if (right === undefined) {
        right = left;
    }

    let val = parseInt(left) * 10 + parseInt(right);
    sum += val;
}

console.log(sum);