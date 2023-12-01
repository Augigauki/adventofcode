const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './part1.txt'), 'utf-8');

const input = text.split('\n');

const exampleInput = [
	'two1nine',
	'eightwothree',
	'abcone2threexyz',
	'xtwone3four',
	'4nineeightseven2',
	'zoneight234',
	'7pqrstsixteen',
];
const edgeCase = ['two1sixnine2one5six'];

const sumNums = (a, b) => a + b;

const regex = /one|two|three|four|five|six|seven|eight|nine/;

const lineSums = input.map((line) => {
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
		} else {
			const typedNums = line.slice(i, line.length).match(regex);
			if (typedNums) {
				let index = typedNums.index;
				let value = typedNums[0];
				/* console.log('\n');
				console.log(typedNums);
				console.log(typedNums[0]); */
				if (index === 0) {
					switch (value) {
						case 'one':
							firstDigit ? (lastDigit = '1') : (firstDigit = '1');
							break;
						case 'two':
							firstDigit ? (lastDigit = '2') : (firstDigit = '2');
							break;
						case 'three':
							firstDigit ? (lastDigit = '3') : (firstDigit = '3');
							break;
						case 'four':
							firstDigit ? (lastDigit = '4') : (firstDigit = '4');
							break;
						case 'five':
							firstDigit ? (lastDigit = '5') : (firstDigit = '5');
							break;
						case 'six':
							firstDigit ? (lastDigit = '6') : (firstDigit = '6');
							break;
						case 'seven':
							firstDigit ? (lastDigit = '7') : (firstDigit = '7');
							break;
						case 'eight':
							firstDigit ? (lastDigit = '8') : (firstDigit = '8');
							break;
						case 'nine':
							firstDigit ? (lastDigit = '9') : (firstDigit = '9');
							break;
						default:
							break;
					}
				}
			}
		}
		/* console.log('Char: ' + line[i]);
		console.log(`First Digit: ${firstDigit} | Last digit: ${lastDigit}`); */
	}
	if (firstDigit && lastDigit) {
		lineNumber = firstDigit + lastDigit;
	} else if (firstDigit && !lastDigit) {
		lineNumber = firstDigit + firstDigit;
	} else {
		lineNumber = '0';
	}
	return Number(lineNumber);
});

const inputTotal = lineSums.reduce(sumNums, 0);

console.log(lineSums);
console.log(inputTotal);
