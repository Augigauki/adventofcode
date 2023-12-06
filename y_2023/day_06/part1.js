const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
//const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
//const edgecase = text3.split('\n');

const removeZeroes = (list) => {
	for (let i = 0; i < list.length; i++) {
		if (list[i] === 0) {
			list.splice(i, 1);
		}
	}
	if (list.includes(0)) {
		list = removeZeroes(list);
	}
	return list;
};

const calculateWins = (races) => {
	console.log(races);
    let moved = 0;
    let wins = 0;
    for(let i = 0; i < races.length; i++){
        let raceWins = 0;
        let record = races[i].distance;
        console.log(`Distance to beat: ${record}`)
        for(let j = 0; j < races[i].time; j++){
            let buttonPressed = j;
            let distance = buttonPressed * (races[i].time-j);
            console.log(`Button held down for: ${buttonPressed} | Distance travelled: ${distance}`)
            if(distance > record){
                raceWins += 1;
            }
        }
        if(wins){   
            wins *= raceWins;
        } else {
            wins += raceWins;
        }
        console.log(`Wins for race ${i}: ${raceWins}`)
    }

    return wins;
};

const day06 = (input) => {
	const races = [];
    let wins = 0;
	let times = input[0].split(':')[1].trim().split(' ');
	times = times.map((time) => Number(time));
	times = removeZeroes(times);
	let distances = input[1].split(':')[1].trim().split(' ');
	distances = distances.map((distance) => Number(distance));
	distances = removeZeroes(distances);
	console.log(times);
	console.log(distances);
	for (let i = 0; i < times.length; i++) {
		races.push({ time: times[i], distance: distances[i] });
	}
	wins = calculateWins(races);
    console.log(wins);
};

day06(input);
