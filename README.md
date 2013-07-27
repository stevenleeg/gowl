# Gowl
Gowl is a simple implementation of the [prowl](https://www.prowlapp.com) API in Go because implementing the Prowl API is a pretty good way to get started with learning a lanuage.

## Example code

Gowl's goal is to make sending prowl notifications as easy as possible. This is all the code you need to start sending notifications to your iOS devices:

```
package main

import (
    "gowl"
    "fmt"
)

func main() {
    fmt.Println("Sending test notification")

    notification := &gowl.Notification{
        Url: "http://example.com",
        Application: "Gowl",
        Event: "Test",
        Description: "Hello world",
    }

    g := gowl.New("[Your API key goes here]")
    g.Add(notification)
}
```

# Documentation
Since this is my first venture into writing a Go library, I wanted to make a simple API that accomplishes its task elegantly. If you think anything here breaks this philosophy, go ahead and open an issue and tell me what sucks.

## Func: `New(api_key string) *gowl.Gowl`
Creates a new instance of a `gowl.Gowl` type and returns a pointer to the struct

## Type: `Gowl`
Manages interactions with an API based on an API key.

### Variable: `ApiKey`
The API key for this Gowl instance.

### Func: `Add(notification *Notification) error`
Sends the provided `notification` struct to the API. Returns an error object if anything has gone wrong, nil if the notification was sent successfully.

## Type: `Notification`
Represents a notification that will be sent to the API using `Gowl.Add(...)`

### Variable: `Url`
See [prowl docs](http://www.prowlapp.com/api.php#add).

### Variable: `Application`
See [prowl docs](http://www.prowlapp.com/api.php#add).

### Variable: `Event`
See [prowl docs](http://www.prowlapp.com/api.php#add).

### Variable: `Description`
See [prowl docs](http://www.prowlapp.com/api.php#add).

### Variable: `Priority`
See [prowl docs](http://www.prowlapp.com/api.php#add).

# License
Gowl is released under the [MIT License](http://opensource.org/licenses/MIT). Do whatever you want with the code as long as you have fun with it!
