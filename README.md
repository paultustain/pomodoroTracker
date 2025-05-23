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

Project will be started in Go. Using Postgres. 

### v0 
This is going to be a pomodoro timer, ran in local host. 
- ~Create project and run local host.~ 
- ~Add a start, stop button.~ 
- ~Add a countdown.~
- ~Add Work/Break notification~
- ~Time resets if stopped.~ 
- Add skip to next phase button

### v1
Add todo list element and project name to timer 
- Add to do list 
- Add tick boxes to mark as complete 
- Add colours to indicate work, break, long break and pause 

### v2
Store project details in local database. 
- Create DB using postgres 
- Create DB functionality to get projects/details from DB 
- Load details and show on screen
- Overall project aims, session aims. 

### v3 
Make it look better.
- Add pause button
- Make sure it looks clean and polished. 
- Add user login ability 

### v4 
- Earn certain number of coins for time focussed. 
- Visuals for collected coins

### v4 
- Ability to show hide all completed tasks for a project
- Create a tauri desktop app to go alongside. 


### Potential future enhancements ideas - may or may not be added after usage. 
- Ability to fuzzy match / search projects to prevent multiple different names. 
- Show all projects on the front page to open up - like the tubely list all videos/ then go into videos
