package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

const timeOutDelay = time.Duration(10)

/*SetupFirebaseFromEnv create a new *auth.Client, using env variable You must init
FIREBASE_KEY_PATH environment variable to the "root directory path"
to the credentials file
*/
func SetupFirebaseFromEnv() *auth.Client {
	log.Printf("Setting up Firebase client.. \n")
	if keyPath := os.Getenv("FIREBASE_KEY_PATH"); keyPath != "" {
		opt := option.WithCredentialsFile(keyPath)
		//Firebase admin SDK initialization
		ctx, _ := context.WithTimeout(context.Background(), time.Second*timeOutDelay)
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(err)
		}
		//Firebase Auth
		authCtx, _ := context.WithTimeout(context.Background(), time.Second*timeOutDelay)
		auth, err := app.Auth(authCtx)
		if err != nil {
			panic(err)
		}
		log.Printf("Firebase client setup :) \n")
		return auth
	}
	log.Printf("Failed to setup wireframe. \n")
	return nil
}

/*SetupFirebase create a new *auth.Client. keyPath is the path to the
wireframe json file.
*/
func SetupFirebase(keyPath string) *auth.Client {
	log.Printf("Setting up Firebase client.. \n")
	if keyPath := os.Getenv(keyPath); keyPath != "" {
		opt := option.WithCredentialsFile(keyPath)
		//Firebase admin SDK initialization
		ctx, _ := context.WithTimeout(context.Background(), time.Second*timeOutDelay)
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(err)
		}
		//Firebase Auth
		authCtx, _ := context.WithTimeout(context.Background(), time.Second*timeOutDelay)
		auth, err := app.Auth(authCtx)
		if err != nil {
			panic(err)
		}
		log.Printf("Firebase client setup :) \n")
		return auth
	}
	log.Printf("Failed to setup wireframe. \n")
	return nil
}
