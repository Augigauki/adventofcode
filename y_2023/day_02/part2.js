const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');

const input = text.split('\n');

const exampleInput = [
	'Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green',
	'Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue',
	'Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red',
	'Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red',
	'Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green',
];

const edgeCase = ['Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red'];

const maxRed = 12;
const maxGreen = 13;
const maxBlue = 14;

let sum = 0;

for (let i = 0; i < input.length; i++) {
	let line = input[i];
	let split = line.split(': ');
	let gameNum = split[0];
	let subsets = split[1];
	let id = Number(gameNum.split('Game ')[1]);
    let minRed, minGreen, minBlue;
    console.log(subsets);
	let rounds = subsets.split(';');
    for(let j = 0; j < rounds.length; j++){
        let hand = rounds[j];
        let red, green, blue;
        let cubes = hand.split(',');
        
        for(let k = 0; k < cubes.length; k++){
            let cubesOfColor = cubes[k];
            cubesOfColor = cubesOfColor.trim();
            //console.log(cubesOfColor)
            if(cubesOfColor.endsWith('blue')){
                blue = Number(cubesOfColor.split(' ')[0])
                if(minBlue){
                    blue > minBlue ? minBlue = blue : null;
                } else {
                    minBlue = blue;
                }

            } else if(cubesOfColor.endsWith('green')){
               
                green = Number(cubesOfColor.split(' ')[0])
                if(minGreen){
                    green > minGreen ? minGreen = green : null;
                } else {
                    minGreen = green;
                }
            } else if(cubesOfColor.endsWith('red')){
                red = Number(cubesOfColor.split(' ')[0])
                if(minRed){
                    red > minRed ? minRed = red : null;
                } else {
                    minRed = red;
                }
            }
            
        }
        if(blue > maxBlue || green > maxGreen || red > maxRed){
            possible = false;
        }
        //console.log(`Red: ${red} | Green: ${green} | Blue: ${blue}\n`)
        

    }
    sum = sum + minRed * minGreen * minBlue
	
}

console.log('Sum: ' + sum);
