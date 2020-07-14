# What happens when `go test` is run?
* Allows us to inspect what happens when `go test` is run
* go tool looks at source file, builds new binary out of all test source files
* stores binary in temporary folder
* see by $ ps -u rbennett | grep go
* binary sleep.test can be run
* running `go build` doesn't include any files ending in '_test.go'
```go
func TestTmpExecutable(t *testing.T) {
	time.Sleep(time.Minute)
}
```

# Naming convention caveats
- `export_test.go` to access unexported variables in external tests
- `xxx_internal_test.go` for internal tests, for example `math_internal_test.go` for math package
- `example_xxx_test.go` for examples in isolated files

# Naming conventions for test functions
- func TestDog (where Dog is a type), or
- func TestDog_puppy (where '_puppy' describes some specific use case)
- func TestSpeak(where Speak is a function)
- func TestDog_Bark(where Bark is a function of Dog)
- func TestDog_Bark_muzzeled(where Bark is a function of dog, and we're testing when the dog is muzzled)

# Naming conventions for variables
- 'got' for what we got
- 'want' for expected, what we want
- "if got != want" (got usually comes first)

# Signaling test failure
- Log, Logf: will only show up if a test fails, or upon `go test -v` 
- Don't use fmt.Println since that won't have info on which test printed it
- Fail, FailNow: mark that test as having failed. FailNow will stop that test immediately. Not used very frequently, since you can't pass log msgs in
- Error/Errorf: Log + Fail
- Fatal/Fatalf: Log + FailNow

# Which to use: Error(just mark a test as failed) or Fatal(mark a test as failed AND end execution)?
- If continuing the test gives more information, use Error. If no more useful info, use Fatal

# For now, note that http.NewRecorder is a fake response writer you can pass into writer