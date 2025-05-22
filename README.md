# pomodoroTracker

## Overview
Not been releasing anything, just sticking it in tutorial hell. Spinning my wheels and going nowhere. 
The plan for this project is to build, release and use something useful. A small fast feedback loop, just with myself. 

Current issues when I try projects.
- Starting with huge ideas. Want the next big thing that solves all problems. Work for a week no stopping, then give up and move on. 
- No scope. It keeps moving along and changing on what is shiny on that day. 
- Nothing is released or finished to a point I actually use it. 
- Only build from tutorials, and when I reach the end and it hits the 'keep going with this as a base' stage, guess what I move on. 
- Always finding the perfect solution and agonising over it. Rust vs Go, Svelt, Tauri or Dioxus. etc.

Aims in this build to get over that. 
- Make it super easy. Maximum of 4 points in a 'release'
- Plan the scope (done below) into blocks. 
- First few releases, actually use the tool as feedback. 
- Once I am done, share it on boot.dev discord for feedback. 
- State what the plan will be including languages, tools and frameworks before starting.

## Plans 

The project will be written using a rust backend utilising Tauri and Sveltkit for frontend and Postgres databasing with refinery for management.
These may or may not be best for this, but for this project I won't be questioning it. 
All sections should have some test cases run. 

### v0 
- ~Get rust working in nvim.~ 
- ~Rust linting / code completion etc set up.~
- ~Create simple app in Tauri~
- Build a webapp using Svelte.
- Get a timer running.
- Print timer to a GUI screen. 

### v1
- Add buttons to pause/stop. 
- Add functionality for task list 
- Background colours Green for work, yellow short break, red for long break. 

### v2
- Store time worked on a project if a project is reopened. 
- Tick boxes by the tasks to mark as completed. 
- Earn coins for time focused. 
- Store task list for a project - link in DB.

### v3 
- Visual element to the timer. 
- Earn certain number of coins for time focussed. 
- Visuals for collected coins

### v4 
- Ability to show hide all completed tasks for a project
- Create a tauri desktop app to go alongside. 


### Potential future enhancements ideas - may or may not be added after usage. 
- Ability to fuzzy match / search projects to prevent multiple different names. 
- Store
