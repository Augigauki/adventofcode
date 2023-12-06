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
    //console.log(input);
    if(!input){
        return;
    }
	for (let i = 0; i < input.length; i++) {
        let added = false;
		for (let j = 0; j < map.length; j++) {
			if (input[i] >= map[j].start && input[i] <= map[j].stop) {
				values.push(input[i] + map[j].diff);
                added = true;
			}
		}
        if(!added){
            values.push(input[i]);

        }
	}
	/* for (let i = 0; i < input.length; i++) {
		values.push(input[i]);
	} */

	return values;
};

const day05 = (input) => {
	let seeds = input[0].split(':')[1].trim().split(' ');
	seeds = seeds.map((number) => Number(number));
    let lowestLocation = 0;
	
	//console.log(seeds);

	let maps = [];

	for (let i = 0; i < input.length; i++) {
		if (input[i] !== '') {
			if (input[i].endsWith('map:')) {
				maps.push(generateMap(input, i + 1));
			}
		}
	}
	console.log(maps[0]);
	let values = [];
    values.push(seeds);
	for (let i = 0; i < maps.length; i++) {
        values.push(getValues(values[i], maps[i]));
		if(i === 0){
            
        }
	}
	console.log(values);
    let locations = values[values.length-1];
    console.log(locations);
    for(let i = 0; i < locations.length; i++){
        if(lowestLocation){
            if(locations[i] < lowestLocation){
                lowestLocation = locations[i];
            }
        } else {
            lowestLocation = locations[i];
        }
    }
    return lowestLocation;

};

result = day05(input);

console.log(result);

/* 
Seed number 79 corresponds to soil number 81.
Seed number 14 corresponds to soil number 14.
Seed number 55 corresponds to soil number 57.
Seed number 13 corresponds to soil number 13.
 */
