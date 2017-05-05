Title: How to Contribute to Code?
Desc: Describing steps and guidelines of how to contribute to aah framework code?
Keywords: contribute to code, guidelines
---
## Contributor Guidelines
* Get the codebase using `aah contribute` command
* Implement your changes and Format code using `goimports` command
* Write unit test cases for your implementation, to know it works
* Provide documentation for your implementation
* Push your fork
* Submit your Pull Request against `integration` branch

### Getting aah framework codebase
It's easy to get started with aah framework codebase. aah CLI tool comes handy here.

**Execute aah contribute command**
```bash
aah contribute -gu=<github-username> -m=<module-name> -g=<goal> -b=<branch-name>

# Execute help command to know about usage
aah help contribute

# For example:
aah contribute -gu=jeevatkm -m=log -g=feature -b=logging-http-receiver
aah contribute -gu=jeevatkm -m=log -g=bug -b=logging-fix
```

**Above command prepares the aah codebase for your contribution. What does `contribute` command do?**

* It checks `git`, `go` and `GOPATH`
* It does `git clone` of all required aah codebase (basically github.com/g-aah/<required-one>) against `integration` branch
* It does `git remote add fork` to add your fork as remote repository
* It creates a branch based on given `module-name`, `goal` and `branch-name`
* Now you're ready to make changes to your local copy of aah framework codebase

### Implement your changes
Now go ahead and implement your changes using your favorite editor. I use Atom editor with `go-plus` add-on installed. Once you're done format Go code using `goimports` command.

### Write unit test cases for your implementation, to know it works
Whichever the module you have implemented your changes, execute below command.
```bash
# go to aah module directory
cd $GOPATH/src/aahframework.org/<module-name>

# execute the all the test cases
go test -race -cover ./...
```

### Provide documentation for your implementation
Documentation is vital information for every developer like you. Help your fellow developer with details.

### Push your fork
Execute below command.
```bash
git push fork integration
```

### Submit your Pull Request against `integration` branch
Follow the below steps-

* Goto your repository `https://github.com/<your-username>/<module-name>`
* If you see a button `Compare & pull request`, click on it OR Click on button `New pull request`, next to branch dropdown
* Choose `base` as `integration` branch on left hand side and review your changes displayed on the page
* Provide your pull request title and description
* Click on button `Create pull request`

Congrats, you have successfully made your code contribution :)
