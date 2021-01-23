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


/*SetupFirebaseFromEnv create a new *auth.Client, using env variable You must init
FIREBASE_KEY_PATH environment variable to the "root directory path"
to the credentials file
*/
func SetupFirebaseFromEnv() *auth.Client {
	log.Printf("Setting up Firebase client.. \n")
	if keyPath := os.Getenv("FIREBASE_KEY_PATH"); keyPath != "" {
		opt := option.WithCredentialsFile(keyPath)
		//Firebase admin SDK initialization
		ctx, ctxErr := context.WithTimeout(context.Background(), time.Second * 10)
		if ctxErr != nil {
			panic(ctxErr)
		}
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(err)
		}
		//Firebase Auth
		authCtx, authCtxErr := context.WithTimeout(context.Background(), time.Second * 10)
		if authCtxErr != nil {
			panic(authCtxErr)
		}
		auth, err := app.Auth(authCtx)
		if err != nil {
			panic(err)
		}
		log.Printf("Firebase client setup :) \n")
		return auth
	}
	log.Printf("Failed to setup firebase. \n")
	return nil
}

/*SetupFirebase create a new *auth.Client. keyPath is the path to the
firebase json file.
*/
func SetupFirebase(keyPath string) *auth.Client {
	log.Printf("Setting up Firebase client.. \n")
	if keyPath := os.Getenv(keyPath); keyPath != "" {
		opt := option.WithCredentialsFile(keyPath)
		//Firebase admin SDK initialization
		ctx, ctxErr := context.WithTimeout(context.Background(), time.Second * 10)
		if ctxErr != nil {
			panic(ctxErr)
		}
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(err)
		}
		//Firebase Auth
		authCtx, authCtxErr := context.WithTimeout(context.Background(), time.Second * 10)
		if authCtxErr != nil {
			panic(authCtxErr)
		}
		auth, err := app.Auth(authCtx)
		if err != nil {
			panic(err)
		}
		log.Printf("Firebase client setup :) \n")
		return auth
	}
	log.Printf("Failed to setup firebase. \n")
	return nil
}