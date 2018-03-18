# GoToGym
<img src="https://ci.jpoles1.com/api/badges/jpoles1/GoToGym/status.svg" height="25"/>     <a href="https://ci.jpoles1.com/jpoles1/GoToGym" target="_blank"><img src="https://raw.githubusercontent.com/jpoles1/GoToGym/master/coverage_badge.png" height="25"/></a>
A go app to tell you to go to the gym.
## Intended Behaviour
- Connect to a google calendar url and fetch events from a "GoToGym" calendar
 - We are currently receiving a webhook from IFTTT after each event in the calendar ends.
- After each gym event passes, send an email asking if the session was completed
- Track attendance on dashboard.
