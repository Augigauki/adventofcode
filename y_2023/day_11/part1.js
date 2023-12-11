const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
const text3 = fs.readFileSync(path.join(__dirname, './example2.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
const example2 = text3.split('\n');

let galaxies = [];
let minPaths = [];
let allPaths = [];

const expandMap = (chart) => {
	let newChart = chart;
	let hAdded = 0;
	let vAdded = 0;
	for (let i = 0; i < chart.length; i++) {
		
		if (chart[i].includes('#')) {
			//console.log('Horizontal line contains a galaxy')
		} else {
			console.log('horizontal line empty at index ' + i);
			//newChart = chart.toSpliced(i, 0, chart[i]);
			newChart = newChart.toSpliced(i+hAdded, 0, newChart[i+hAdded])
			hAdded++;
			//console.log('\n');
			//console.log(chart);
		}
		let empty = true;
		for (let j = 0; j < chart[i].length; j++) {
			for(let k = 0; k < chart.length; k++){

				if (chart[i][j] !== '.') {
					empty = false;
					break;
				}
			}
		}
		if (empty) {
			console.log('vertical line empty at index ' + i);
			console.log(chart[i]);
			for (let j = 0; j < chart.length; j++) {
				console.log(`i: ${i} j: ${j}`)
				//console.log(chart[i].slice(0, j-1) + '.' + chart[i].slice(j));
				newChart[i] = chart[i].slice(0, j-1) + '.' + chart[i].slice(j);
			}
		}
		//console.log(newChart);
	}
	return newChart;
};

const day11 = (input) => {
	console.log(input);
	input = expandMap(input);
	console.log(input);
	return;
	for (let i = 0; i < input.length; i++) {
		for (let j = 0; j < input[i].length; j++) {
			let char = input[i][j];
			if (char === '#') {
				galaxies.push({ x: j + 1, y: i + 1, paths: [] });
			}
		}
	}
	//console.log(galaxies);
	for (let i = 0; i < galaxies.length; i++) {
		let index = 0;
		console.log(`\nGalaxy: ${i + 1}`);
		let curr = galaxies[i];
		while (galaxies[i].paths.length < galaxies.length - 1) {
			//if both x and y of the next galaxy are greater

			if (index === galaxies.length) {
				break;
			}
			let next = galaxies[index];
			if (curr !== next) {
				let highY,
					highX,
					lowY,
					lowX = 0;
				if (curr.x > next.x) {
					highX = curr.x;
					lowX = next.x;
				} else {
					highX = next.x;
					lowX = curr.x;
				}
				if (curr.y > next.y) {
					highY = curr.y;
					lowY = next.y;
				} else {
					highY = next.y;
					lowY = curr.y;
				}
				if (highY === lowY && highX === lowX) {
				} else {
					let steps = highY - lowY + (highX - lowX);
					console.log(
						`comparing galaxy ${i + 1} to galaxy ${
							index + 1
						} high y: ${highY} | high x: ${highX} | low y: ${lowY} | low x: ${lowX}. Steps: ${steps}`
					);
					galaxies[i].paths.push(steps);
				}
			}

			index++;

			if (index === 9) {
				//console.log('resetting index')
				index = 0;
			}
		}
	}
	//galaxies[0].paths.push((galaxies[1].x-galaxies[0].x)+(galaxies[1].y-galaxies[0].y))
	//console.log(galaxies);
	galaxies.forEach((galaxy, index) => {
		minPaths.push(Math.min(...galaxy.paths));
		allPaths.push(galaxy.paths.reduce((acc, curr) => acc + curr));
	});
	console.log(minPaths);
	console.log(allPaths.reduce((acc, curr) => acc + curr) / 2);
};

day11(example);
