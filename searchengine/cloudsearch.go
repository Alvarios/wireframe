package searchengine

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"os"
)

/*NewClient create a new cloudsearchdomain.CloudSearchDomain instance. The instance is configured using
aws environment variables REGION, PROFILE, AWS_SECRET_ACCESS_KEY, AWS_ACCESS_KEY_ID.
*/
func NewClient() *cloudsearchdomain.CloudSearchDomain {
	region := os.Getenv("REGION")
	profile := os.Getenv("PROFILE")
	cred := credentials.NewEnvCredentials()
	config := aws.Config{
		Credentials: cred,
		Region:      aws.String(region),
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config:  config,
	}))

	// Create a CloudSearchDomain client from just a session.
	svc := cloudsearchdomain.New(sess)

	return svc
}

/*SearchEngine used aws CloudSearch to search element into the database or can be used to get search suggestions.
 */
type SearchEngine interface {
	// Suggest take a string in parameter, this string contains the query of the user. The method will return
	// suggestions based on this phrase.
	Suggest(Suggest string) (string, error)

	//SimpleSearch take a query search all text and text-array fields for the specified string.
	//Search for phrases, individual terms, and prefixes.
	SimpleSearch(string) (*cloudsearchdomain.SearchOutput, error)

	//StructuredSearch search specific fields, construct compound queries using
	//Boolean operators, and use advanced features such as term boosting and
	//proximity searching.
	StructuredSearch(string) (*cloudsearchdomain.SearchOutput, error)
}
