#firebase 

Create a new auth.Client by simply calling one method and setting one env variable.

## Need 

* json file that contains your firebase token 
* set FIREBASE_KEY_PATH = "path to your json file"

# Code 

````go
package main 

import "github.com/alvarios/wireframe/firebase"
import "fmt"

func main(){
    client := firebase.SetupFirebaseFromEnv()
    fmt.Printf("This is my new client %v" , client)
}
```` 
The setup function panic if it's fail to init the nex firebase client.