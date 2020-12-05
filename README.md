# Run app
go run <app.go>

# Build app
go build <app.go>

# go.mod file(could be custom name or github repo url)
go mod init <anyName>

# Check go version
go version

# API reference
https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them

# Get package(requires .mod file)
go get <package>

# Repo 
https://go.dev/

# allows you to use packages without using "go get <package>"
go mod tidy

# Find out what's behind a repo
go mod why -m <package>

# Clean up removed dependencies from codebase
go mod tidy

# Tutorial

https://golang.org/doc/tutorial/getting-started

# To use simple local packages, you must use "go mod init <bananas>" 
# so you can use something like this on main:

import "bananas/somegofile"

main() {
    somegofile.function()
}

# if you want a function to be public not private, you must use uppercase(making it lowercase will make the function NOT ACCESSIBLE)
* we are in "spotify" directory, we also used "go mod init spotify" to create base project name
utility/yoyo.yoyo.go
package yoyo

func Wing() string {
	return "yo bitch"
}

and use in main like:

import (
	"spotify/utility/yoyo"
)

...

func otherShit() {
	fmt.Println(yoyo.Wing())
}


# Package name importance.. folder name & package name must identical, if not you must alias(if your package name is different than the folder).. the package name can be anything..and file name does not matter; what's important is folder name and package name being used

(scenario one - you name "package" with a name not using FOLDER name)
../utility/strings.go

package blah

func getPackageName() string {
	return "yo"
}

then you must use it like this on main:

package main

import (
	blah "spotify/utility"
)

...
blah.someFunc()

(scenario two - standard - you use same package name as folder name)

package utility

func getPackageName() string {
	return "yo"
}

then you must use it like this on main:


import (
	"spotify/utility"
)

...

utility.blah()





