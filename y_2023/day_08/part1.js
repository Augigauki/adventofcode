const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
const edgecase = text3.split('\n');

let steps = 0;

const getNodeCode = (nodeLine) => {
	return nodeLine.split(' = ')[0];
};

const getLeft = (nodeLine) => {
	return nodeLine.split(' = ')[1].split(', ')[0].split('(')[1];
};

const getRight = (nodeLine) => {
	return nodeLine.split(' = ')[1].split(', ')[1].split(')')[0];
};

const traverseNode = (turn, node, nodeList) => {
	let nextNode;
	if (turn === 'L') {
		nextNode = getLeft(node);
	} else {
		nextNode = getRight(node);
	}
	for (let i = 0; i < nodeList.length; i++) {
		if (nodeList[i].startsWith(nextNode)) {
			nextNode = nodeList[i];
		}
	}
	steps++;
	return nextNode;
};

let traversed = '';
let traversed2 = '';
const day08 = (list) => {
	const directions = list[0];
	let turnIndex = 0;
	let turn = directions[turnIndex];
	const nodesList = list.slice(2);
	let currNode = nodesList.filter((node) => {
		return getNodeCode(node) === 'AAA';
	});
    currNode = currNode[0];

	//repeat until current node starts with 'ZZZ'
	while (!currNode.startsWith('ZZZ')) {
		//restart directions if we run out
		if (turnIndex === directions.length) {
			turnIndex = 0;
		}
        turn = directions[turnIndex];
        
        if(traversed.length !== directions.length && traversed.length < directions.length){
            traversed += directions[turnIndex];
        } else {
            traversed2 += directions[turnIndex];
        }
		//console.log('*****BEFORE CHANGE*****')
		console.log(`${currNode} | Turn: ${turn}`);

		//set node to next one based on which one we traverse to
		currNode = traverseNode(turn, currNode, nodesList);

		//increase turnIndex by one and set turn to the next letter in the instructions
		turnIndex++;
	}
	console.log('Reached end :)');
	console.log(steps);
};

day08(input);
