package main

type createRepoCommand struct{}

var createRepoCmd createRepoCommand

// Execute
func (x *createRepoCommand) Execute(args []string) error {
	name, err := askQuestion("Please, enter new repository name:")
	if err != nil {
		return err
	}

	return ctx.ES.CreateAWSRepository(name, creds.BucketName, creds.Region)
}
