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

# (For now) Know that http.NewRecorder is a fake response writer you can pass into writer

# Writing good error messages
- goal for error messages: tell developers what went wrong, help them debug the situation 
- Succint but useful template: Function name with parameters if they're short. It equaled got,  wanted want. 
-- Useful even if SomeFunc returns value & error. It's inferred we're not using
```go
t.Errorf("SomeFunc(%v) = %v; want %v", param, got, want)
```
- Similar, with err: 
```go
t.Errorf("SomeFunc(%v) err = %v; want %v")
```
- Don't worry about printing exact params into SomeFunc. For example, if p is a Person struct with many fields, and it was passed in to SomeFunc(p), we could say: `t.Errorf("SomeFunc(name=%s, age=%d)", p.Name, p.Age)`

# Examples As Test Cases (package example)
- examples can be run as test cases.
- a reason they're valuable: incentive to keep docs/examples up to date. otherwise they will fail
- Examples are the same as tests except: 
1. functions start with 'Example' like funct ExampleHello() instead of TestHello()
2. The example funcs don't take any args
3. Need comment: `// Output:`
- For dealing with maps/goroutines, where order will be different, use `//Unordered output:`
- For different examples on the same function, use `ExampleHello_spanish()` to produce `Example (Spanish)` in the docs

- Can run in docs like:
`godoc -http=:8080`

## Package level Examples: 
- In order to show imports in examples, create a new file for the example. Must also have a package level const or var

# Table Driven Tests (package underscore)
- with simple Table driven tests, must use Errorf in loop. If fatal is used, other rows of table won't be run

# Subtests 
- t.Run() takes a name and a subtestFunction. Signature of subTest takes t *testing.t, but doens't need the capital. Options for subtests:
1. Anonymous closure function
```go
for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Camel(tt.arg); got != tt.want {
                t.Errorf("Camel()=%v, want %v", got, tt.want)
        }
    })
}
```
2. Named function
```go
for _, tt := range tests {
        t.Run(tt.name, somesubTest)
}
func somesubTest(t *testing.T) {
    // ...
}
```
3. Function Returning closure
```go
for _, tt := range tests {
        t.Run(tt.name, appTest(app)) {
    }
}

func appTest(app *App) func(t *testing.T) {
    return func(t *testing.T) {
        // use app
    }
}
```
- Nesting in t.Run is possible
```go
for _, tt := range tests {
    r.Run(tt.name, func(t *testing.T) {
        t.Run("another subtest", func(t *testing.T) {
            if got := Came(tt.arg); got != tt.want {
                t.Errorf("Camel() = %v, want %v", got, tt.want)
            }
        })
    })
}
```
- Useful for top level with setup, then each test has different tests
- Alternatives of test name 
-- Instead of doing test struct with name, args, and want, can make a map[string]struct{} where string is name, and struct contains args and wants
-- Can sometimes pass arg as name

## Subtests can be failed with Fatal
- It's useful to use fatal if we have further subchecks that are useless if we fail the first one
- With subtests, we are able to use fatal in each subtest b/c each subtest is run in separate goroutines

# TestMain
- Steps that should be included:
	// 0. flag.Parse() if you need flags
	// 1. Setup anything you need
	// 2. Run tests
    // 3. Exit the code
- can't fail a test in Main
- os.Exit(exitCode) will stop the program immediately - no deferred funcs will be run