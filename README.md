# AmexChallenge
Here is my implementiation of the Orders challenge!

## Instructions

Run  `go run .` to start the server. The server will listen on port `8080`. I have branches for each step and the step 3 implementaion has been merged into the `dev` and `main` branches.
### Steps 1 and 2

`POST /order` - This end point expects a post request with two form values `apples` and `oranges` which both should have positive numberical arguments. Upon supplying a valid form the server will return a summary json that contains fields for the sent apple and orange values as well as the cost of the order.

### Step 3
Run `docker compose up -d` to start both the server and the postgres database. 

`POST /order` - This endpoint works the same as it does in step 1 and 2, this time though returning the id of the created order.

`GET /order?id=` - Here we expect a query paramater named `id` which should be a positive integer and will corrispond to the requested order. A valid `id` will cause the server to return the stored order from the database.

`GET /all` - This returns a json list of every order currently in the database table.

## Testing

I have implemented simple tests for step 1 and 2 that test for the correct cost. However due to me turing this in slightly late due to personal complications and other interview preperation I have ommited the use of a mocking library to correctly and fully test the database usage as well as incoming requests. This is due mostly to time constaints and my lacking familiarity with a apt mocking library for this project.

I would like to state that I do understand the importance of using mocking librarys to test very much value full testing in all my work, and hope that this will not impede my ability to move along further and meet the team!

### Steps 1 and 2

Running `go test ./test/...` will run the Cost test which tests for correct cost.

## Step 3

The same command will work however I have commented out the test code.