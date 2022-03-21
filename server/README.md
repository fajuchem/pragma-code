## Server

I didn't had time to do much here, most of the time was used adding tests, also
most of the changes was done to make easier to test.

### what was done
- Removed unused code.
- Created package for configs (easier to test the code)
- Isolated sensor calls to its own package.
- Added unit tests.

### what we could be done
- Get server url from env instead of hardcoded.
- More robust http comunication.
- Few edge cases errors with async code are not handle correctly.
- Better overall error handling.
- Maybe use websocket or webhook instead calling the sensor api if the sensor
api supports that, and also use websockets to push changes to client.
