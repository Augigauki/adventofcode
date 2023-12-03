const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');

const input = text.split('\n');

const example = [
	'467..114..',
	'...*......',
	'..35..633.',
	'......#...',
	'617*......',
	'.....+.58.',
	'..592.....',
	'......755.',
	'...$.*....',
	'.664.598..',
];

const edgeCase = ['..+35*633..'];

let sum = 0;
let partNums = [];
const sumNums = (a, b) => a + b;

const arrayRange = (start, stop, step) => {
	return Array.from(
		{ length: (stop - start) / step + 1 },
		(value, index) => start + index * step
	);
};

const day03 = (list) => {
	for (let i = 0; i < list.length; i++) {
		let currNum = '';
		let maybePartNum = '';

		let line = list[i];
		for (let j = 0; j < line.length; j++) {
			let char = line[j];
			if (char.match(/[0-9]/)) {
				currNum += char;
				checkNext = () => {
					if (line[j + 1].match(/[0-9]/)) {
						currNum += line[j + 1];
						j += 1;
					} else {
						maybePartNum = currNum;
					}
					if (j + 1 === line.length) {
						console.log('yo');
						maybePartNum = currNum;
					}
				};
				while (j + 1 < line.length && !maybePartNum) {
					checkNext();
				}
			} else {
				currNum = '';
				maybePartNum = '';
			}
			if (maybePartNum !== '') {
				console.log(maybePartNum);
				let added = false;
				let index = j + 1 - maybePartNum.length;
				let numIndexes = arrayRange(
					index,
					index + maybePartNum.length - 1,
					1
				);
				if (index > 0) {
					if (line[index - 1] !== '.') {
						partNums.push(Number(maybePartNum));
						added = true;
					}
				}
				if (index + maybePartNum.length < line.length && !added) {
					if (line[index + maybePartNum.length] !== '.') {
						partNums.push(Number(maybePartNum));
						added = true;
					}
				}
				if (i > 0 && !added) {
					let prevLine = list[i - 1];

					for (let k = 0; k < prevLine.length; k++) {
						if (
							prevLine[k] !== '.' &&
							!prevLine[k].match(/[0-9]/)
						) {
							let symbolIndex = k;

							if (
								numIndexes.includes(symbolIndex) ||
								numIndexes.includes(symbolIndex + 1) ||
								numIndexes.includes(symbolIndex - 1)
							) {
								partNums.push(Number(maybePartNum));
								added = true;
							}
						}
					}
				}
				if (i < list.length - 1 && !added) {
					let nextLine = list[i + 1];
					for (let k = 0; k < nextLine.length; k++) {
						if (
							nextLine[k] !== '.' &&
							!nextLine[k].match(/[0-9]/)
						) {
							let symbolIndex = k;
							if (
								numIndexes.includes(symbolIndex) ||
								numIndexes.includes(symbolIndex + 1) ||
								numIndexes.includes(symbolIndex - 1)
							) {
								partNums.push(Number(maybePartNum));
								added = true;
							}
						}
					}
				}
			}
		}
	}
};

day03(input);
sum = partNums.reduce(sumNums, 0);
console.log(`Sum: ${sum}`);
