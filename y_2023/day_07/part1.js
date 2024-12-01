const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
//const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
//const edgecase = text3.split('\n');

let handTypes = ['highCard', 'onePair', 'twoPair', 'threeOfAKind', 'fullHouse', 'fourOfAKind', 'fiveOfAKind'];

let cardStrength = {
	2: 2,
	3: 3,
	4: 4,
	5: 5,
	6: 6,
	7: 7,
	8: 8,
	9: 9,
	T: 10,
	J: 11,
	Q: 12,
	K: 13,
	A: 14,
};

const cards = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'];

const checkHandType = (hand) => {
	let handType;
	const counts = {};
	hand.forEach((card) => {
		counts[card] = (counts[card] || 0) + 1;
	});
	console.log(counts);
	let val1, val2;
	for (let i = 0; i < cards.length; i++) {
		
		let cardCount = counts[cards[i]];
		//console.log(`i: ${i} | Checking: ${cardCount}`);
		//console.log(`Hand: ${hand} | Cardcount: ${cardCount}`)
		if (cardCount === 5) {
			val1 = `five of ${cards[i]}s`;
			//i += cardCount - 1;
		} else if (cardCount === 4) {
			val1 = `four of ${cards[i]}s`;
			//i += cardCount - 1;
		} else if (cardCount === 3) {
			if (!val1) {
				val1 = `three of ${cards[i]}s`;
				//i += cardCount - 1;
			} else if (val1.startsWith('pair of')) {
				val2 = `full house: 3 ${cards[i]}s, ${val1.split('of ')[1][0]}s`;
				//i += cardCount - 1;
			}
		} else if (cardCount === 2) {
			if (!val1) {
				val1 = `pair of ${cards[i]}s`;
				//i += cardCount - 1;
			} else if (val1.startsWith('pair of')) {
				val2 = 'two pairs: ' + val1.split('of ')[1][0] + 's and ' + cards[i] + 's';
				//i += cardCount - 1;
			} else if (val1.startsWith('three of')) {
				val2 = `full house: 3 ${val1.split('of ')[1][0]}s, 2 ${cards[i]}s`;
				//i += cardCount - 1;
			}
		} else {

		}
		if (val1) {
			if (val2) {
				handType = val2;
			} else {
				handType = val1;
			}
		} else {
			handType = 'high card'
		}
	}

	return handType;
};

const day07 = (hands) => {
	for (let i = 0; i < hands.length; i++) {
		let line = hands[i];
		let hand = line.split(' ')[0];
		let score = line.split(' ')[1];
		hand = hand.split('');
		//hand.sort();
		//console.log(hand);

		let type = checkHandType(hand);
		console.log('\n' + hand + ' ' + type + '\n');
	}
};

day07(example);
