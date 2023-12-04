const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
const edgecase = text3.split('\n');

let sum = 0;

const calcMatches = (winNums, elfNums) => {
	let matches = 0;
	for (let i = 0; i < elfNums.length; i++) {
		for (let j = 0; j < winNums.length; j++) {
			if (elfNums[i] === winNums[j]) {
				matches += 1;
			}
		}
	}
	return matches;
};

const removeZeroes = (list) => {
	for (let i = 0; i < list.length; i++) {
		if (list[i] === 0) {
			list.splice(i, 1);
		}
	}
	return list;
};

const getLineID = (line) => {
	return Number(line.split(':')[0].split('Card')[1].trim(' '));
};

const day04 = (list) => {
	let matches = 0;

	for (let i = 0; i < list.length; i++) {
		let line = list[i];
		let splitNums = line.split(':')[1].split('|');
		let id = getLineID(line);
		let winningNums = splitNums[0].trim().split(' ');
		let elfNums = splitNums[1].trim().split(' ');

		winningNums = winningNums.map((digit) => Number(digit));
		elfNums = elfNums.map((digit) => Number(digit));
		winningNums = removeZeroes(winningNums);
		elfNums = removeZeroes(elfNums);
		matches = calcMatches(winningNums, elfNums);
		let toBeAdded = [];
		for (let j = id-1; j < list.length; j++) {
			if (getLineID(list[j]) > id) {
				if (!toBeAdded.includes(getLineID(list[j]))) {
					if (getLineID(list[j]) > id) {
						//console.log(`Adding ${getLineID(list[j])} to toBeAdded`)
						toBeAdded.push(getLineID(list[j]));
					}
				}
			}

			if (toBeAdded.length === matches) {
				break;
			}
		}
		//console.log(`Card: ${id} | i: ${i} | Matches: ${matches}`);
		//console.log(toBeAdded);

		for (let j = list.length - 1; j > 0; j--) {
			let otherId = getLineID(list[j]);
			if (toBeAdded.includes(otherId)) {
				if (getLineID(list[j - 1]) !== otherId) {
					list.splice(j, 0, list[j]);
					toBeAdded.pop();
				}
			}
            if(toBeAdded.length === 0){
                break;
            }
		}

		list.sort();
		//console.log(list);
		/* if(id === 2){
            break;
        } */
	}
	return list.length;
};

sum = day04(input);

console.log(sum);
