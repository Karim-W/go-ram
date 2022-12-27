# go-ram

small api server for relaying common gordon ramsey quotes/insults with added functionality of delivering gornald ramsay inspired insults to "end users"

## Usage
|Method|Endpoint|Description|type|
|---|---|---|---|
|GET|`/`|Returns an HTML page with random gordon ramsey quote|HTML|
|GET|`/api/v1/insults`|Returns a JSON array of insults|JSON|
|GET|`/api/v1/insults/random`|Returns a random insult|string|
|GET|`/api/v1/insults/endusers`|Returns a JSON array of inspired insults for end users|JSON|
|GET|`/api/v1/insults/endusers/random`|Returns a random inspired insult for end users|string|

## Contributing
Pull requests are welcome.
