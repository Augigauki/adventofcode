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

const edgeCase = [
	'.32...5123',
	'.+35*633..',
	'..22.5....'];

let numRegex = /[0-9]/;

let sum = 0;
let gearRatios = [];
const sumNums = (a, b) => a + b;

const day03 = (list) => {
	for (let i = 0; i < list.length; i++) {
		let line = list[i];
		for (let j = 0; j < line.length; j++) {
			let char = line[j];
			if (char === '*') {
				let gearIndex = j;
				let gearNum = '';
				let prevLine;
				let nextLine;
				i > 0 ? (prevLine = list[i - 1]) : null;
				i < list.length - 1 ? (nextLine = list[i + 1]) : null;
				let adjacentNums = [];

				//Sjekke char før og etter
				for (let k = gearIndex; k < gearIndex + 1; k++) {
					let before = false;
					let after = false;
					//Sjekke om char før * er tall
					if (k > 0 && !before) {
						if (line[k - 1].match(numRegex)) {
							let lastIndex = k - 1;
							for (let l = lastIndex; l >= 0; l--) {
								if (line[l].match(numRegex)) {
									gearNum += line[l];
								} else {
									gearNum = gearNum.split('').reverse().join('');
									if (gearNum.length > 0) {
										adjacentNums.push(Number(gearNum));
									}
									before = true;
									gearNum = '';
									break;
								}
								if (l === 0) {
									gearNum = gearNum.split('').reverse().join('');
									if (gearNum.length > 0) {
										adjacentNums.push(Number(gearNum));
										
									}
									before = true;
									gearNum = '';
								}
							}
						}
					}
					//Sjekke om char etter * er tall
					if (k < line.length - 1 && !after) {
						if (line[k + 1].match(numRegex)) {
							let firstIndex = k + 1;
							for (let l = firstIndex; l < line.length; l++) {
								if (line[l].match(numRegex)) {
									gearNum += line[l];
								} else {
									if (gearNum.length > 0) {
										adjacentNums.push(Number(gearNum));
									}
									after = true;
									gearNum = '';
									break;
								}
								if (l === line.length-1) {
									if (gearNum.length > 0) {
										adjacentNums.push(Number(gearNum));
									}
									after = true;
									gearNum = '';
								}
							}
						}
					}
				}
				//sjekke forrige linje
				if (prevLine) {
					for (let k = gearIndex; k < gearIndex + 1; k++) {
						let before = false;
						let over = false;
						let after = false;
						//sjekker char diagonalt opp til venstre
						if (k > 0) {
							if (prevLine[k - 1].match(numRegex)) {
								let lastIndex = k - 1;
								for (let l = lastIndex; l <= prevLine.length; l--) {
									if (prevLine[l].match(numRegex)) {
										lastIndex = l;
									} else {
										break;
									}
									if (l === 0) {
										lastIndex = l;
										break;
									}
								}
								for (let l = lastIndex; l < prevLine.length; l++) {
									if (prevLine[l].match(numRegex) && !before) {
										gearNum += prevLine[l];
										l >= k ? over = true : over = false;
										l >= k+1 ? after = true : after = false;
									} else {
										if (gearNum.length > 0) {
											adjacentNums.push(Number(gearNum));
										}
										before = true;
										gearNum = '';
										break;
									}
									if(l === prevLine.length-1){
										if(gearNum.length > 0){
											adjacentNums.push(Number(gearNum));
										}
										over = true;
										gearNum = '';
										
									}
								}
							}
						}
						//sjekker char over dersom det ikke er del av forrige tallet
						if (prevLine[k].match(numRegex) && !over) {
							for (let l = k; l < prevLine.length; l++) {
								if (prevLine[l].match(numRegex)) {
									gearNum += prevLine[l];
									l >= k+1 ? after = true : after = false;
								} else {
									if (gearNum.length > 0) {
										adjacentNums.push(Number(gearNum));
									}
									over = true;
									gearNum = '';
									break;
								}
								if(l === prevLine.length-1){
									if(gearNum.length > 0){
										adjacentNums.push(Number(gearNum));
									}
									over = true;
									gearNum = '';
									
								}
							}
						}
						//sjekker tall diagonalt opp til høyre dersom det ikke er del av et tidligere tall
						if (k < prevLine.length - 1 && !after) {
							if (prevLine[k + 1].match(numRegex)) {
								for(let l = k + 1; l < prevLine.length; l++){
									if(prevLine[l].match(numRegex)){
										gearNum += prevLine[l]; 
									} else {
										if(gearNum.length > 0){
											adjacentNums.push(Number(gearNum));
										}
										after = true;
										gearNum = '';
										break;
									}
									if(l === prevLine.length-1){
										if(gearNum.length > 0){
											adjacentNums.push(Number(gearNum));
										}
										over = true;
										gearNum = '';
									}
								}
							}
						}
					}
				}
				//sjekke neste linje
				if (nextLine) {
					for (let k = gearIndex; k < gearIndex + 1; k++) {
						let before = false;
						let over = false;
						let after = false;
						//sjekker char diagonalt opp til venstre
						if (k > 0) {
							if (nextLine[k - 1].match(numRegex)) {
								let lastIndex = k - 1;
								for (let l = lastIndex; l <= nextLine.length; l--) {
									if (nextLine[l].match(numRegex)) {
										lastIndex = l;
									} else {
										break;
									}
									if (l === 0) {
										lastIndex = l;
										break;
									}
								}
								for (let l = lastIndex; l < nextLine.length; l++) {
									if (nextLine[l].match(numRegex) && !before) {
										gearNum += nextLine[l];
										l >= k ? over = true : over = false;
										l >= k+1 ? after = true : after = false;
									} else {
										if (gearNum.length > 0) {
											adjacentNums.push(Number(gearNum));
										}
										before = true;
										gearNum = '';
										break;
									}
									if(l === nextLine.length-1){
										if(gearNum.length > 0){
											adjacentNums.push(Number(gearNum));
										}
										over = true;
										gearNum = '';
									}
								}
							}
						}
						//sjekker char over dersom det ikke er del av forrige tallet
						if (nextLine[k].match(numRegex) && !over) {
							for (let l = k; l < nextLine.length; l++) {
								if (nextLine[l].match(numRegex)) {
									gearNum += nextLine[l];
									l >= k+1 ? after = true : after = false;
								} else {
									if (gearNum.length > 0) {
										adjacentNums.push(Number(gearNum));
									}
									over = true;
									gearNum = '';
									break;
								}
								if(l === nextLine.length-1){
									if(gearNum.length > 0){
										adjacentNums.push(Number(gearNum));
									}
									over = true;
									gearNum = '';
								}
							}
						}
						//sjekker tall diagonalt opp til høyre dersom det ikke er del av et tidligere tall
						if (k < nextLine.length - 1 && !after) {
							if (nextLine[k + 1].match(numRegex)) {
								for(let l = k + 1; l < nextLine.length; l++){
									if(nextLine[l].match(numRegex)){
										gearNum += nextLine[l]; 
									} else {
										if(gearNum.length > 0){
											adjacentNums.push(Number(gearNum));
										}
										after = true;
										gearNum = '';
										break;
									}
									if(l === nextLine.length-1){
										if(gearNum.length > 0){
											adjacentNums.push(Number(gearNum));
										}
										over = true;
										gearNum = '';
									}
								}
							}
						}
					}
				}
				if (adjacentNums.length === 2) {
					gearRatios.push(adjacentNums[0] * adjacentNums[1]);
				}
			}
		}
	}
};

day03(input);
sum = gearRatios.reduce(sumNums, 0);
console.log(`Sum: ${sum}`);//too low
