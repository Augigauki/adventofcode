const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
const edgecase = text3.split('\n');

//console.log(example);



const generateMap = (list, index) => {
	let map = [];
	for (let i = index; i < list.length; i++) {
		if (list[i] !== '') {
            let digits = list[i].split(' ');
            digits = digits.map((digit) => Number(digit))
			map.push(digits);
		} else {
			break;
		}
	}
	return map;
};

const day05 = (input) => {
	let seeds = input[0].split(':')[1].trim().split(' ');
    seeds = seeds.map((number) => Number(number));
    let seedNums = [];
    for(let i = 0; i <= Math.max(...seeds); i++){
        seedNums.push(i);
    }
	console.log(seeds);
    //console.log(seedNums);

	let maps = [];

	for (let i = 0; i < input.length; i++) {
		if (input[i] !== '') {
			if (input[i].endsWith('map:')) {
				maps.push(generateMap(input, i + 1));
			}
		}
	}
	console.log(maps);
};

day05(example);
