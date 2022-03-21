## Client

### what was done
- Added typescript.
- Converted main component from class to function base.
- Moved fetching data from the api to it's own file.
- Created generic function to fetch data from the api. Ideally we could be created
a more robust abstraction to allow differents http methods, but for now to keep it
simple I only added one function to get requests.
- Segregated product listing into it's own component.
- Created a repository for the products listing to hide all the tecnical problem,
for now we only have the call to the api endpoint. Also created an interface to
define the contract between server-side and client-side.
- Added unit tests.

### what we could be done
- Add linting.
- Add e2e tests.
- Get server url from env instead of hardcoded.
- More robust http comunication.
- Better error handling.
- Separate loading into its own component.
- Maybe use websockes instead of polling.
