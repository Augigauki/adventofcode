const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
const example2 = text3.split('\n');

let sum = 0;

const calcPoints = (winNums, elfNums) => {
	let points = 0;
	for (let i = 0; i < elfNums.length; i++) {
		for (let j = 0; j < winNums.length; j++) {
			if (elfNums[i] === winNums[j]) {
				if (points === 0) {
					points = 1;
				} else {
					points = points * 2;
				}
			}
		}
	}
	return points;
};

const removeZeroes = (list) => {
	for (let i = 0; i < list.length; i++) {
		if (list[i] === 0) {
			list.splice(i, 1);
		}
	}
	return list;
};

const day04 = (list) => {
	let sum = 0;
	for (let i = 0; i < list.length; i++) {
		let line = list[i];
		let splitNums = line.split(':')[1].split('|');
		let id = line.split(':')[0].split('Card')[1].trim(' ');
		let winningNums = splitNums[0].trim();
		winningNums = winningNums.split(' ');

		let elfNums = splitNums[1].trim().split(' ');
		winningNums = winningNums.map((digit) => Number(digit));
		elfNums = elfNums.map((digit) => Number(digit));
		winningNums = removeZeroes(winningNums);
		elfNums = removeZeroes(elfNums);
		sum += calcPoints(winningNums, elfNums);
	}
	return sum;
};

sum = day04(example);

console.log(sum);
