const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
const text3 = fs.readFileSync(path.join(__dirname, './example2.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
const example2 = text3.split('\n');

const nRegex = /7|F|S|\|/;
const sRegex = /L|J|S|\|/;
const eRegex = /7|J|S|\-/;
const wRegex = /L|F|S|\-/;

const pipePath = [];

const followPath = (pos) => {
	let x = pos.x;
	let y = pos.y;
    let prevX = pos.prevX;
    let prevY = pos.prevY;
    console.log('\n')
    console.log(pos);
	let pipe = pipePath[pipePath.length-1];
	console.log(pipe);
	if (pipePath.length > 1 && pipePath[pipePath.length - 1] === 'S') {
		console.log(path.length);
	} else {
        /* if(pipePath.length > 10){
            return;
        } */
		if (y > 0) {
			if (map[y - 1][x].match(nRegex) && prevY+1 !== y) {
				if (pipe.match(sRegex)) {
					console.log('Found matching pipe for ' + pipe + ' north: ' + map[y - 1][x]);
					pipePath.push(map[y - 1][x]);
					//followPath({ x: x, y: y - 1, prevX: x, prevY: y });
					return { x: x, y: y - 1, prevX: x, prevY: y };
				}
			}
		}
		if (x < map[y].length-1) {
			if (map[y][x + 1].match(eRegex) && prevX-1 !== x) {
				if (pipe.match(wRegex)) {
					console.log('Found matching pipe for ' + pipe + ' east: ' + map[y][x + 1]);
					pipePath.push(map[y][x + 1]);
					//followPath({ x: x + 1, y: y, prevX: x, prevY: y });
					return { x: x + 1, y: y, prevX: x, prevY: y };
				}
			}
		}
		if (y < map.length-1) {
			if (map[y + 1][x].match(sRegex) && prevY-1 !== y) {
				if (pipe.match(nRegex)) {
					console.log('Found matching pipe for ' + pipe + ' south: ' + map[y + 1][x]);
					pipePath.push(map[y + 1][x]);
					//followPath({ x: x, y: y + 1, prevX: x, prevY: y });
					return { x: x, y: y + 1, prevX: x, prevY: y };
				}
			}
		}
		if (x > 0) {
			if (map[y][x - 1].match(wRegex) && prevX+1 !== x) {
				if (pipe.match(eRegex)) {
					console.log('Found matching pipe for ' + pipe + ' west: ' + map[y][x - 1]);
					pipePath.push(map[y][x - 1]);
					//followPath({ x: x - 1, y: y, prevX: x, prevY: y });
					return {x: x-1, y: y, prevX: x, prevY: y};
				}
			}
		}
        
	}
};

const day10 = (input) => {
	for (let i = 0; i < input.length; i++) {
		for (let j = 0; j < input[i].length; j++) {
			if (input[i][j] === 'S') {
				pipePath.push(input[i][j]);
                let nextPipePos = followPath({x: j, y:i, prevX: j-1, prevY: i-1})
                while(input[nextPipePos.y][nextPipePos.x] !== 'S'){
                    if(pipePath.length > 2 && pipePath[pipePath.length-1] === 'S'){
                        //console.log('breaking')
                        break;
                    }
                    nextPipePos = followPath({x: nextPipePos.x, y: nextPipePos.y, prevX: nextPipePos.prevX, prevY: nextPipePos.prevY})
                }
			}
		}
	}
    console.log(pipePath);
	console.log((pipePath.length-1)/2);
};
const map = input;
//console.log(map);
day10(input);
