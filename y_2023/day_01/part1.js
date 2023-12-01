const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './part1.txt'), 'utf-8');

const textByLine = text.split('\n');

const exampleInput = ['1abc2', 'pqr3stu8vwx', 'a1b2c3d4e5f', 'treb7uchet'];

const sumNums = (a, b) => a + b;

const result = textByLine
	.map((line) => {
		let firstDigit;
		let lastDigit;
		let lineNumber;
		for (let i = 0; i < line.length; i++) {
			if (/[0-9]/.test(line[i])) {
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
	.reduce(sumNums, 0);

console.log(result);
