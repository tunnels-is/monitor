# monitor

# tool
 - chackra ui
 - reactjs / htxm ? (dealers choice)

# ???
## basic connectivity
Server and monitor are connected via TCP socket
SERVER <- tcp -> MONITOR

## initial connection
SERVER -> diffie-hellman + apikey -> MONITOR

## log flow
SERVER -> tcp -> monitor

## websocket for feed
monitor > ws > ui

## api for searching
monitor < api > ui

# UI SOMETHING ??
## main page
 - flex row layout
 - each log TAG/TYPE has it's own section with log lines in it.
 - do not show sections that have no log lines.

## per server page
 - flex row layout
 - each server tag has it's own section

## filter page
 - flex column layout
 - single section with all logs
 - input box at top to apply filter

## search page
This sends a request to a search api instead of feeding from the websocket.
 - flex column layout
 - single section with all logs
 - input box at top + day selection (1days, 2days, 3days, etc..)

