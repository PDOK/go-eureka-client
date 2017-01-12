Go Eureka Client
================

Based on code from https://github.com/ArthurHlt/go-eureka-client and https://github.com/bryanstephens/go-eureka-client .

## Getting started

```go
import (
	"github.com/pdok/go-eureka-client/eureka"
)
func main() {

	client := eureka.NewClient([]string{
		"http://127.0.0.1:8761/eureka", //From a spring boot based eureka server
		// add others servers here
	})
	instance := eureka.NewInstanceInfo("test.com", "test", "69.172.200.235", 80, false) //Create a new instance to register
	client.RegisterInstance(instance) // Register new instance in your eureka(s)
	applications, _ := client.GetApplications() // Retrieves all applications from eureka server(s)
	client.GetApplication(instance.App) // retrieve the application "test"
	client.GetInstance(instance.App, instance.HostName) // retrieve the instance from "test.com" inside "test"" app
	client.SendHeartbeat(instance.HostName) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
}
```


All these strange behaviour come from Eureka.

