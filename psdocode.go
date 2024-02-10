package main

/* Import packages:
- "net/http" for http requets
- "html/template" for rendering serverside templates
- "log" for logging errors
-  Other packages like middleware added as needed

Global Varibles:
- "slice" to hold data
- "struct" for configs
- "constants" dafault port number to be used

Data Models/ Struct:
- "Data" we will be working with AE "json"
- ""

Main Func:
- "create" main function of program
- "create" routes for / and API ext
- "define" the port your created
- "start" the server

Routes:
 - "Forloop" through each routes
 - "Parse" each request
 - "Fetch" from DB or other sources as needed
 - "Setup" responces from APIs for json ext

   //HTML if creating webpages
     - define and create templates and data
   //Static Files
     - create file server to link to folders to create webspages

   //Middleware
     - setup of middleware (ea - auth ext...)
     - error handling

Tests:
 - "Setups tests"
*/
