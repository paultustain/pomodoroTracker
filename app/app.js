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

let currentTime = 0;

window.onload = function() {
	getProjects();
	document.getElementById('timer-section').style.display ='none';
}


// FIX DELAY TO START BY 1 SECOND.
// Initial code is delayed. 
startButton.addEventListener('click', function() {
	if (!isRunning) {
		stage.textContent="Work"
		isRunning = true;
		duration = workDuration;

		timerId = setInterval(function() {
			time.textContent = formatTime(duration);
			document.getElementById('time-spent').textContent = formatTime(currentTime);
			currentTime++;

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
		 $("body").css("background-color", yellow);
		return shortBreak
	} else {
		stage.textContent = "Work"
		 $("body").css("background-color", green);
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
	document.getElementById('create-project-form').reset();
});

document.getElementById('create-task-form').addEventListener('submit', async (event) => {
	event.preventDefault();
	await createTask();
	document.getElementById('create-task-form').reset();
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
				
				const projectLabel = document.createElement('b');
				const deleteButton = document.createElement("b");
				const pipeMarker = document.createElement("b");

				projectLabel.onclick = () => openProject(project);
				projectLabel.textContent = project.Name; 

				deleteButton.textContent = "(Delete - should be picture of a bin)";
				deleteButton.style.color = "red";
				deleteButton.onclick = async () => await deleteProject(project);
				
				pipeMarker.textContent = " | ";
				
				projectList.appendChild(projectLabel);
				projectList.appendChild(pipeMarker);
				
				projectList.appendChild(deleteButton);
				projectList.appendChild(document.createElement("br"));

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
				const label= document.createElement('label');
				const description = document.createTextNode(task.Task);
				const checkbox = document.createElement("input");
				const deleteButton = document.createElement("b");
				const pipeMarker = document.createElement("b");

				if (task.Completed) {
					label.style.color = "gray";
					label.style.setProperty("text-decoration", "line-through");
					checkbox.checked = true;
				}
				checkbox.type = "checkbox";
				checkbox.name = "slct[]"
				checkbox.value = task.Task;
				
				deleteButton.textContent = "(Delete - should be picture of a bin)";
				deleteButton.style.color = "red";
				deleteButton.onclick = async () => await deleteTask(task);
				
				pipeMarker.textContent = " | ";

				checkbox.addEventListener('change', async (event) => { completeTask(checkbox, label, task.ID) })
				
				label.appendChild(checkbox);
				label.appendChild(description);
								
				taskList.appendChild(label);
				taskList.appendChild(pipeMarker);
				taskList.appendChild(deleteButton);
				taskList.appendChild(document.createElement("br"));
			}
		}
	} catch (error) {
		alert(`Error: ${error.message}`);
	}
}

async function completeTask(checkbox, label, taskID) {
	// can change this function based on how to display
	// tasks. Hide all tasks on completion, set to grey and leave in session
	// show all tasks for a project. 
	
	if (checkbox.checked) {
		apiLink = `api/completeTask/${taskID}`
		label.style.color = "gray";
		label.style.setProperty("text-decoration", "line-through");
	} else {
		apiLink = `api/completeTask/${taskID}`
		label.style.color = "black";
		label.style.setProperty("text-decoration", "none");
	}
	try {
		const res = await fetch(apiLink, {
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
			const data = await res.json();
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
	currentTime = project.TimeSpent;
	projectTime.textContent = formatTime(project.TimeSpent);
	await getTasks();

}

async function deleteProject(project) {
	try {
		const res = await fetch( `api/deleteProject/${project.ID}`, {
			method: "DELETE"
		});

		if (!res.ok) {
			const data = await res.json();
			throw new Error(`Failed to delete project ${data.error}`)
		}
	} catch (error) {
		throw new Error(`Failed to delete project.`)
	}

	await getProjects();
}


async function deleteTask(task) {
	try {
		const res = await fetch( `api/deleteTask/${task.ID}`, {
			method: "DELETE"
		});

		if (!res.ok) {
			const data = await res.json();
			throw new Error(`Failed to delete task: ${data.error}`)
		}
	} catch (error) {
		throw new Error(`Failed to delete task.`)
	}

	await getTasks();
}
