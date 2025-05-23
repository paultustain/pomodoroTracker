// Buttons
const startButton = document.querySelector('#startButton')
const stopButton = document.querySelector('#stopButton')

// Outputs
let time = document.getElementById('time')

let timerId;
let workDuration = 60 * 25;
let shortBreak = 60 * 5;
let longBreak = 60 * 15;
let totalWorked = 0

startButton.addEventListener('click', function() {
	console.log(isPaused)
	if (isPaused) {

	} else {
		if (!timerId){ 
			duration = workDuration;
			timerId = setInterval(function() {
				time.textContent = formatTime(duration);
				duration--;
				totalWorked++;
			}, 1000);
		} else {
			console.log("time id exists")
		}
	}
});

stopButton.addEventListener('click', function() {
	clearInterval(timerId);
	isPaused=false;
	time.textContent = "Finished"
	console.log("You worked for " + formatTime(totalWorked) + " minutes")
	let timerId;
})


function formatSeconds(value) {
	return value < 10 ? "0" + value : value;
}

function formatTime(duration) {
		minutes = parseInt(duration / 60, 10);
		seconds = parseInt(duration % 60, 10);
		minutes = formatSeconds(minutes);
		seconds = formatSeconds(seconds);
		return minutes + ":" + seconds;
}
