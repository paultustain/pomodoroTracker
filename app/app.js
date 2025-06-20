// Buttons
const startButton = document.querySelector('#startButton')
const stopButton = document.querySelector('#stopButton')
const skipButton = document.querySelector('#skipButton')
// Values 
const stage = document.querySelector('#stage')

// Outputs
let time = document.getElementById('time')

// Defaults 
let timerId;
let workDuration = 60 * 25;
let shortBreak = 60 * 5;
let longBreak = 60 * 15;
let totalWorked = 0;
let isRunning = false;


// FIX DELAY TO START BY 1 SECOND.
// Initial code is delayed. 
startButton.addEventListener('click', function() {
	if (!isRunning) {
		stage.textContent="Work"
		isRunning = true;
		duration = workDuration;

		timerId = setInterval(function() {
			time.textContent = formatTime(duration);
			duration--;
			if (stage.textContent == "Work") {
				totalWorked++;
			}
			if (duration == 0){
				duration = changeStage();
			}
		}, 1000);
	} 
});

skipButton.addEventListener('click', function() {
	duration = changeStage();
	clearInterval(timerId);
	isRunning=true;

	time.textContent = formatTime(duration);
	timerId = setInterval(function() {
		time.textContent = formatTime(duration);
		duration--;
		if (stage.textContent == "Work") {
			totalWorked++;
		}
		if (duration == 0){
			duration = changeStage();
		}
	}, 1000);
})


stopButton.addEventListener('click', function() {
	clearInterval(timerId);
	isRunning = false;
	time.textContent = "25:00"
	stage.textContent = "Stopped"
	console.log("You worked for " + formatTime(totalWorked) + " minutes")
	totalWorked=0;
})

function changeStage() {
	if (stage.textContent == "Work") {
		stage.textContent = "Break"
		return shortBreak
	} else {
		stage.textContent = "Work"
		return workDuration
	}
}
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

document.getElementById('create-project-form').addEventListener('submit', async (event) => {
	event.preventDefault();
	await createProject();
});

async function createProject() {
	const name = document.getElementById('project-name').value;
	try {
		const res = await fetch('api/createProject', {
			method: 'POST', 
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({name}),
		});
		const data = res.json();
		if (!res.ok) {
			throw new Error(`Failed to create project:  ${data.error}`);
		}
	} catch (error) {
		alert(`Error: ${error.message}`);
	}
}
