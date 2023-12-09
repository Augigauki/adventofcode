const path = require('path');
const fs = require('fs');

const text = fs.readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
const text2 = fs.readFileSync(path.join(__dirname, './example.txt'), 'utf-8');
//const text3 = fs.readFileSync(path.join(__dirname, './edgecase.txt'), 'utf-8');

const input = text.split('\n');
const example = text2.split('\n');
//const edgecase = text3.split('\n');

const day09 = (list) => {
    let total = 0;
    for(let i = 0; i < list.length; i++){
        let line = list[i].split(' ');
        console.log('\n***NEW LINE***')
        let lastVals = [];
        line = line.map((val) => Number(val));
        console.log(line);
        lastVals.push(line[line.length-1]);
        //console.log(lastVals);
        let diffs = [];
        for(let j = line.length-1; j > 0; j--){
            diffs.push(line[j]-line[j-1]);
        }
        console.log(`Diffs: ${diffs}`);
        while(diffs.some((diff) => diff !== 0)){
            
            console.log(`Adding ${diffs[0]} to lastVals`)
            lastVals.push(diffs[0]);
            
            diffs = diffs.map((value, index) => {
                //console.log(index);
                if(typeof diffs[index+1] !== 'undefined'){
                    //console.log(`Subtracting ${diffs[index+1]} from ${value}`)
                    return value-diffs[index+1]
                }
            })
            diffs.pop();
            console.log(`Diffs: ${diffs}`);
            
        }
        total += lastVals.reduceRight((acc, curr) => acc + curr);
        console.log(`Last vals: ${lastVals}`);
        
    }
    console.log(total);
}

day09(example);