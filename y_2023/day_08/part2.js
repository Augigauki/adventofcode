const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example2.txt'), 'utf-8');
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
	
	return nextNode;
};

const doesEndInZ = (node) => {
    return getNodeCode(node).endsWith('Z');
}

const allZ = (nodes) => {
    for(let i = 0; i < nodes.length; i++){
        if(!getNodeCode(nodes[i]).endsWith('Z')){
            return false;
        }
    }
    return true;
}



const day08 = (list) => {
	const directions = list[0];
	let turnIndex = 0;
	let turn = directions[turnIndex];
	const nodesList = list.slice(2);
	let aNodes = nodesList.filter((node) => {
		return getNodeCode(node).endsWith('A');
	});

	let aNodesSteps = [];
    console.log(aNodes);

	//repeat until current node starts with 'ZZZ'
	while (!allZ(aNodes)) {
		//restart directions if we run out
		if (turnIndex === directions.length) {
			turnIndex = 0;
		}
        turn = directions[turnIndex];
		steps++;
        for(let i = 0; i < aNodes.length; i++){
            aNodes[i] = traverseNode(turn, aNodes[i], nodesList);
			if(doesEndInZ(aNodes[i])){
				aNodesSteps.push({stepCount: steps, divided: steps/directions.length});
			}
        }
        
        console.log(`${aNodes}`)
		if(aNodesSteps.length === 6){
			break;
		}
		//set node to next one based on which one we traverse to
		

		//increase turnIndex by one and set turn to the next letter in the instructions
		turnIndex++;
	}
	console.log('Reached end :)');
	console.log(aNodesSteps);
	
	for(let i = aNodesSteps.length-1; i > 0; i--){

		console.log()
		
	}
	//console.log(steps);
};

day08(input);
