const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './part1.txt'), 'utf-8');

const textByLine = text.split('\n');

const exampleInput = ['two1nine', 'eightwothree', 'abcone2threexyz', 'xtwone3four', '4nineeightseven2', 'zoneight234', '7pqrstsixteen'];

const sumNums = (a, b) => a + b;

const lineSums = exampleInput.map((line) => {
	let firstDigit;
	let lastDigit;
	let lineNumber;
	for (let i = 0; i < line.length; i++) {
		if (/[0-9]['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']/.test(line[i])) {
			if (!firstDigit) {
				firstDigit = line[i];
			} else {
				lastDigit = line[i];
			}
		}
	}
	if (firstDigit && lastDigit) {
		lineNumber = firstDigit + lastDigit;
	} else if (firstDigit && !lastDigit) {
		lineNumber = firstDigit + firstDigit;
	} else {
		lineNumber = '0';
	}
	return Number(lineNumber);
})

const inputTotal = lineSums.reduce(sumNums, 0);

console.log(lineSums);
console.log(inputTotal);
