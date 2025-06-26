// Buttons
const startButton = document.querySelector('#startButton')
const stopButton = document.querySelector('#stopButton')
const skipButton = document.querySelector('#skipButton')
// Values 
const stage = document.querySelector('#stage')
let selectedProject = "";
// Outputs
let time = document.getElementById('time')

// Defaults 
let timerId;
let workDuration = 60 * 25;
let shortBreak = 60 * 5;
let longBreak = 60 * 15;
let totalWorked = 0;
let isRunning = false;


window.onload = function() {
	getProjects();
	document.getElementById('timer-section').style.display ='none';
}

// Can we store project ID when selected?

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


stopButton.addEventListener('click', async function() {
	clearInterval(timerId);
	isRunning = false;
	time.textContent = "25:00"
	stage.textContent = "Stopped"


	await updateProjectTime(totalWorked);
	 
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

document.getElementById('create-task-form').addEventListener('submit', async (event) => {
	event.preventDefault();
	await createTask();
})

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
		if (!res.ok) {
			const data = res.json();
			throw new Error(`Failed to create project:  ${data.error}`);
		}

		await getProjects();
	} catch (error) {
		alert(`Error: ${error.message}`);
	}
}

async function createTask() {
	const name = document.getElementById('task-name').value;
	console.log(`Task: ${name}`)
	console.log(`Within Project: ${selectedProject}`)
	try {
		const res = await fetch('api/createTask', {
			method: 'POST', 
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				project_id: selectedProject, 
				task_description: name,
				
			})
		})
		await getTasks();
	} catch (error) {
		alert(`Error: ${error.message}`)
	}
}

async function getProjects() {
	try {
		const res = await fetch('api/getProjects', {
			method: 'GET', 
		});

		if (!res.ok) { 
			const data = await res.json(); 
			throw new Error(`Failed to get projects. Error: ${data.error}`)
		}

		const projects = await res.json();
		const projectList = document.getElementById('project-list')
		projectList.innerHTML = '';
		
		if (projects !== null) {
			for (const project of projects) {
				const listItem = document.createElement('li');
				listItem.onclick = () => openProject(project);
				listItem.textContent = project.Name;
				projectList.appendChild(listItem);
			}
		}
	} catch (error) {
		alert(`Error: ${error.message}`);
	}
}

async function getTasks() {
	try {
		const res = await fetch(`api/getTasks/${selectedProject}`, {
			method: 'GET', 
		});

		if (!res.ok) { 
			const data = await res.json(); 
			throw new Error(`Failed to get projects. Error: ${data.error}`)
		}

		const tasks = await res.json();
		const taskList = document.getElementById('open-tasks')
		var apiLink = ``;
		taskList.innerHTML = '';

		if (tasks !== null) {
			for (const task of tasks) {
				const listItem = document.createElement('li');
				// On click make it change to light grey
				listItem.onclick = () => {
					if (listItem.style.color !== "gray") {
						apiLink = `api/completeTask/${task.ID}`
						listItem.style.color = "gray";
					} else {
						apiLink = `api/completeTask/${task.ID}`
						listItem.style.color = "black";
					}
					completeTask(apiLink)
				}
				listItem.textContent = task.Task;
				taskList.appendChild(listItem);
			}
		}
	} catch (error) {
		alert(`Error: ${error.message}`);
	}
}

async function completeTask(taskID) {
	try {

		const res = fetch(apiLink, {
			method: 'POST'
		});
		if (!res.ok) { 
			const data = await res.json(); 
			throw new Error(`Failed to update task. Error: ${data.error}`)
		}
	} catch (error) {
		throw new Error (`Failed to complete task`);
	}
}

async function updateProjectTime(timeAdded) {
	try {
		const res = await fetch('api/updateTime', {
			method: 'POST', 
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				id: selectedProject,
				time: timeAdded, 
			}),
		});
		if (!res.ok) {
			const data = res.json();
			throw new Error(`Failed to update project time: ${data.error}`)
		}
		const project = await res.json();
		const projectTime = document.getElementById('time-spent');
		projectTime.textContent = formatTime(project.TimeSpent);
		
		await getProjects();
	} catch (error) {
		alert(`Error: ${error.message}`);
	}

}

async function openProject(project) {
	document.getElementById('timer-section').style.display ='block';
	selectedProject = project.ID;
	const projectName = document.getElementById('project-name-selected');
	projectName.textContent = project.Name;
	
	const projectTime = document.getElementById('time-spent');
	projectTime.textContent = formatTime(project.TimeSpent);
	await getTasks();

}
