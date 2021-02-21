package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	hiveError "github.com/Alvarios/wireframe/error"
	"github.com/kataras/iris/v12"
	"strings"
)

//ThemisJWTMiddlewareProtocol - jwt middleware protocol
type JWTMiddlewareProtocol interface {
	FirebaseAuthenticationMiddleware(ctx iris.Context)
}

/*JWTMiddleware Middleware struct that contains a *auth.Client that is used to check if the JWT token
// present in the header are valid Token. This middleware is use only if the server used firebase to
// authenticate user
*/
type JWTMiddleware struct {
	FirebaseAuthClient *auth.Client
}

//UIDParameterKey - use this value to get the user PublisherID store in the context parameters
const UIDParameterKey = "PublisherID"

//AuthorizationTokenKey - if the user use an authorization, this is the key of this element in the Header
const AuthorizationTokenKey = "Authorization"

//FirebaseAuthenticationMiddleware middleware used to check the firebase token
func (auth JWTMiddleware) FirebaseAuthenticationMiddleware(ctx iris.Context) {
	authorizationToken := ctx.GetHeader(AuthorizationTokenKey)
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		badRequest := hiveError.NewHiveBadRequestError(ctx, "No token present in the header")
		ctx.JSON(badRequest)
		ctx.EndRequest()
		return
	}
	//verify token
	token, err := auth.CheckToken(idToken)
	if err != nil {
		unauthorizedError := hiveError.NewHiveUnauthorizedError(ctx, "Invalid token")
		ctx.JSON(unauthorizedError)
		ctx.EndRequest()
		return
	}
	ctx.Params().Set(UIDParameterKey, token.UID)
	ctx.Next()
}

//CheckToken - verify ID Token. If the token is valid it return a auth.Token else it return an themisError
func (auth JWTMiddleware) CheckToken(idToken string) (*auth.Token, error) {
	//verify token
	token, err := auth.FirebaseAuthClient.VerifyIDToken(context.Background(), idToken)
	return token, err
}
