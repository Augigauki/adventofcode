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
		if (list[i] === '') {
			list.splice(i, 1);
		}
	}
	if (list.includes('')) {
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

const calculateWin = (race) => {
    let wins = 0;
    let record = race.distance;
    for(let i = 0; i < race.time; i++){
        let buttonPressed = i;
        let distance = buttonPressed * (race.time-i);
        if(distance > record){
            wins += 1;
        }
    }
    return wins;
}

const day06 = (input) => {
    let wins = 0;
    let time = 0;
	let times = input[0].split(':')[1].trim().split(' ');
	//times = times.map((time) => Number(time));
	times = removeZeroes(times);
    let distance = 0;
	let distances = input[1].split(':')[1].trim().split(' ');
	//distances = distances.map((distance) => Number(distance));
	distances = removeZeroes(distances);
    times.map((curr) => time += curr);
    time = Number(time.slice(1));
    distances.map((curr) => distance += curr);
    distance = Number(distance.slice(1));
	console.log(times);
    console.log(time);
	console.log(distances);
    console.log(distance);
	race = {
        time: time,
        distance: distance,
    }
    console.log(race);
	wins = calculateWin(race);
    console.log(wins);
};

day06(input);
