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
			digits = digits.map((digit) => Number(digit));

			map.push({
				range: digits[2],
				start: digits[1],
				diff: digits[0] - digits[1],
				stop: digits[1] + digits[2] - 1,
			});
		} else {
			break;
		}
	}
	return map;
};

const getValues = (input, map) => {
	let values = [];
	if(typeof input[0] === 'object'){
        for(let i = 0; i < input.length; i++){
            for(let j = 0; j < map.length; j++){
                console.log(`Seed start: ${input[i].seed} | Seed stop: ${input[i].end} | Map start: ${map[j].start} | Map end: ${map[j].stop}`)
                if(input[i].seed >= map[j].start && input[i].end <= map[j].stop){
                    console.log('included');
                    input[i].seed += map[j].diff;
                    input[i].end += map[j].diff;
                } else {

                }
            }
        }
    }
	/* for (let i = 0; i < input.length; i++) {
		let added = false;
		for (let j = 0; j < map.length; j++) {
			if (input[i] >= map[j].start && input[i] <= map[j].stop) {
				values.push(input[i] + map[j].diff);
				added = true;
			}
		}
		if (!added) {
			values.push(input[i]);
		}
	} */
	/* for (let i = 0; i < input.length; i++) {
		values.push(input[i]);
	} */

	return values;
};

const day05 = (input) => {
	let seeds = input[0].split(':')[1].trim().split(' ');
	let seedPairs = [];
	let lowestLocation = 0;
	for (let i = 0; i < seeds.length; i++) {
		if (i % 2 === 0) {
			seedPairs.push({ seed: Number(seeds[i]), range: Number(seeds[i + 1]), end: Number(seeds[i])+Number(seeds[i+1]-1) });
		}
	}

	seeds = [];
	for (let i = 0; i < seedPairs.length; i++) {
		for (let j = 0; j < seedPairs[i].range; j++) {
			seeds.push(seedPairs[i].seed + j);
		}
	}

	//console.log(seeds);
    //console.log('Seeds.length: ' + seeds.length);

	let maps = [];

	for (let i = 0; i < input.length; i++) {
		if (input[i] !== '') {
			if (input[i].endsWith('map:')) {
				maps.push(generateMap(input, i + 1));
			}
		}
	}
	let values = [];
	values.push(seedPairs);

	for (let i = 0; i < maps.length; i++) {
		values.push(getValues(values[i], maps[i]));
		if (i === 0) {
            //break;
		}
	}
	let locations = values[values.length - 1];
	for (let i = 0; i < locations.length; i++) {
		if (lowestLocation) {
			if (locations[i] < lowestLocation) {
				lowestLocation = locations[i];
			}
		} else {
			lowestLocation = locations[i];
		}
	}
	//console.log('Lowest location: ' + lowestLocation);
};

day05(example);
