//command line app 3 digit number try to guess
//perfect good bad
//-d number of digits prompt user for answer

const program = require('commander');
const promptly = require('promptly');

program
    .version('0.0.1')
    .option('-d, --digits <number>', 'Number of digits')
    .parse(process.argv);


function generateNumber(digits) {
    //	 Math.floor(Math.random() * Math.floor(max));
    let arr = [];
    for (let i = 0; i < digits; i++) {
        arr.push(Math.ceil(Math.random() * Math.floor(9)));
    }
    return arr;
}

async function main() {
    if (program.digits) {

        const newAnswer = generateNumber(program.digits);
        //console.log('numbers', newAnswer);
        let loop = true;
        while (loop) {
            let  guess = await promptly.prompt('Enter your guess');
            let  answers = checkAnswer(guess,newAnswer);
	   console.log("answers", answers);
	    let perfect = 0;
	  for (let i=0; i < newAnswer.length; i++) {
	  	if ("perfect" == answers[i]) {
			perfect++;
		}
	  }
		if (perfect == answers.length) {
			console.log("You got it")
			break;
		}


        }
    } else {
        console.log('Check input to program: needs -d with number')
    }
}

function checkAnswer(guess, answer) {
	let arr = [];
	for (let i=0; i < guess.length; i++) {
		if (guess[i] == answer[i]) {
			arr.push("perfect");
		} else if (checkIncludes(answer, guess[i])) {
			arr.push("good");
		} else {
			arr.push("bad");
		}
	}
	return arr;
}

function checkIncludes(answer, guess) {
	for (let i=0; i < answer.length; i++) {
		if (answer[i] == guess) {
			return true;
		}
	}
	return false;
}

main();
