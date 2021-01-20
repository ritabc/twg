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

# Running Tests In Parallel (package parallel)
## Why to use?
- Simiulating a real-world scenario, for eg when testing a web app that will have multiple users at once
- Verify that a type, like a cache, is truly threadsafe
## Why to NOT use?
- Paralellism is not free, and could mean more work
-- Tests can't use as many hard-coded values; eg unique email constrainst
-- It might be complicated for tests to use shared resources, as in the same output file, or same shared DB connection
- Don't use parallel tests just for the speed test - not worth it b/c could introduce bugs by running in parallel
## How to use
- Add t.Parallel() to any func you want run in parallel
- Any tests not having this line will not run in parallel ; they will run first
## Running subtests in parallel
- Subtests can be run in parallel, but will only be run with their sister subtests
- We need some way to tell when all parallel subtests are finished, in order to run teardown after that. To do so, wrap all parallel subtests in another "group" t.Run()
## Closure Gotcha for Parallel tests
- (Especially with loops or table driven tests) Shadow testCase variable or add another closure. If shadowing, recommended to add comment (`// Do not delete`), as the line will look like: `tt := tt`

# Race Conditions : Testing them
## What is a race condition? 
```go
// Given a value that multiple goroutines will write to / read from, without disregard for what other goroutines are doing to the value
var balance = 100

func main() {
    go spend(30)
    go spend(40)
}

func(spend(amount int) {
    b := balance
    time.Sleep(time.Second)
    b -= amount
    balance = b
}
```
What are the different goroutines doing?
```
time  |  goroutine1  |  goroutine2   | balance
=====================================================
t0    |              |               | `balance := 100`
t1    |              |`b := balance` | 100
t1    |`b := balance`|               | 100
t1    |`b - 30 = 70` |               | 70
t1    |              |`b - 30 = 60`  | 60
t1    |              |               | 60
```
Basically, the balance was copied in each goroutine before decremenint happened

## Using race detection flag
```bash
$ go test -race
$ go run -race thing.go
$ go build -race thing.go
$ go install -race pkg
$ go get -race golang.org....
```
Will notify you if it detects a race condition

Sometimes can be used: if deploying to multiple locations, one of them will be deployed with `go build -race`

It doesn't always catch them., EG:
- if the race involves a DB read/write scenario, the race isn't in memory so it won't be caught

## Testing explicitly for race conditions (dir race_fail)
Looking at code, we don't expect to have race condition

But we show in the tests by explicitly testing for it that it exists

# Comparing Objects for Equality (package compare)
- Two objects with same field values but different adddresses can still be considered equal with == operator
- Structs with func fields cannot be compared
- Unless using reflect.DeepEqual 
- Can use golden files to easily create the 'want' variable, for instance comparing large files/amounts of data. Human-eye check the golden file, then compare to the output

# Helper comparison functions (package alert)
- For example, when testing http responses, we don't want to recreate entire page/response
-- if the footer changes and the footer has nothing to do with our test, we don't want to rewrite the entire test
-- so, write helper function to check for different things within our html functions 
- Other situations which could use helper comparison funcs: 
-- verify specific field in struct is set
-- verify json response has specific subset of data

# Building things with helper functions
- Candidate for helper function: db setup (package helper)
- Candidate for helper function: generating test data (package gen)
-- quick package generates a bunch of fake data and (with the Check function), it checks your function many times with fake input
- Candidate: Interface testing utilities

# Running Specific Test (package parallel)
- Test all tests matching 'TestB'
`go test -v -run TestB` , where TestB is a regexp
- Run all tests for a subpackage
`go test ./...`
`for pkg in *; do go test "./$pkg"; done`
- Skip tests,  few methods:
1. set env variable or flag, etc: `shouldBeSkipped`
```go
var shouldBeSkipped = true

func TestThing(t *testing.T) {
	if shouldBeSkipped {
		t.Skip()
	}
	t.Log("this test ran!")
}
```
2. set testing.Short, run test with `go test -short`:
```go
func TestThing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Log("this test ran!")
}
```
3. Custom Flags
- init with flag.BoolVar: integration
- within TestMain, setup if integration == true
- within test, check first: `if !integration {t.Skip}`, then proceed to test code
4. Build tags
- Tell compiler, only build this source file if we include the build tag psql. At top of the file: `// +build psql`
- To run, `go test -v -tags="psql mysql"` OR `go test -v -tags=psql`

# Benchmarks
- Function names start with 'Benchmark'
- Get passed *testing.B
- loop that runs b.N times (how many times you run your benchmark)
- To run, `go test -bench .` (run all benchmarks)
- Output will be:
 BenchmarkHello-12       22497444 (b.N number)          49.6 ns/op (ns per operation)
- To only time functionality that you want, surround code you actually want to benchmark with b.ResetTimer() and b.StopTimer()
- People sometimes use benchmarks to determine if their application is getting slower over time

# Testing Verbose
- Can do: `if testing.Verbose() { t.Logf("") }`

# Coverage (package cover)
- short version: go tool takes all lines and adds counter before each of them, If the line is run, that line's counter incrememts to 1
- Counter will never be > 1
## Create profile
`$ go test -coverprofile=cover.out`
View profile in function mode 
`$ go tool cover -func=cover.out`
View profile in html mode (there are also others)
`$ go tool cover -html=cover.out`

# Timeouts are useful if things are blocking, to determine if something is broken on your time
`go test -timeout 5s` Default is 10minutes

# External vs Internal tests
## - When to use external tests? 
Anytime when you can, you should
* Forces you to write tests that are agnostic to internal/exported functionality
* Tests will be good indicator about versions - if you break your tests, you're breaking the package and you might want to bump the major version
* Tests will be good examples for other examples
* To get around cyclical dependencies
## Can use unexported vars, funcs, and types in an external test file by...:
Create a new `export_test.go` file
```go
package draw

var Fib = fib
type Dog = dog // Can still not access unexported struct fields. Actually should use internal tests for this reason 
```
This will only be exported in the test env b/c of the `_test.go` in the file name. No docs will be created for it, and the variables will not be exported to other packages 

# Global State
* What is it?
    * variables or other info that's not isolated to whats running
    * eg: database values, global variables
* Avoid when can
## What if we can't avoid global state? 3 Tips
1. avoid paralell tests.
2. and/or specify test order like:
```go
func TestApp(t *testing.T) {
    // Inside here we can run specific tests in a specific order
    testThingA(t)
    testThingB(t)
    testThingC(t)
}
```
3. Use separate setup/teardown btwn each test that alters state
```go
func TestWidget(t *testing.T) {
    // setup -create stuff we need in the DB
    db := // open DB
    user := createTestUser(db)
    widget := createTestWidget(db, user)

    // run tests that use the data in the database
    ...
    
    // teardown - delete stuff
    resetDB(db)
}
```

## Alternatives to global state
- dependency injection

# Dependency Injection (package logger)
## Definition
DI is a design pattern, where when we have a function, we provide the dependencies the function relies on. 

Instead of having the function create the dependency itself

DI enables a lot of other things:
- make implementation agnostic code
- write tests more easily
- simulate specific behavoris in our tests
- avoid global state
- pacakage level functions (package git)

DI can add lines of code, a bit of complexity. But for large projects, is worth it.

# Mocks
## Terminology
- Dummy: simple, doesn't exist in go, a blank implementation, not used at all
- Stub: returns the bare minimum, Not close to real behavior, but takes and returns the right types? Upon calling User.Create(), it'll return the right stuff but won't remember the user
- Fake, aka Double: not quite as intricate as, but sort of comparable to, real implementation. Upon calling User.Create(), it'll return the right stuff AND will remember the user

- Spy: Will keep track of what methers were called
- Mock: Will also keep track of what methods were called
- Spies & mocks are the same in Go, but in other langs a Spy would NOT fail if a method wasn't called, but would give you info

## Remember mocks do not replace integration tests entirely

## Why do we mock?
To simplfy testing , or to make testing possible
- Setup/teardown may be simpler
- Simulating a specific situation (often an error) can be easier (ie; error from full memory/disk, without actually filling memory/disk)
- External APIs may be slow or unreliable
- API may not have a test env, or it may have limits
- We can verify that specific behavior occurs - eg that we call the EmailClient.Welcome after a user is sccessfully created

## Mocking Examples
**Simulate specific situation** 
`twg/race_pass/users_test.go`: Us a fake implementation to simulate a very specific race condition

**Setup/teardown**
Create a real DB and even seed it with data. If there is a good bit of setup/teardown, it can be excessive. Sometimes a test doesn't really need all this to actually tests something. Testing this function might not need a REAL userStore:
```go
func Signup(name, email string, ec EmailClient, us *UserStore) (*User, error) {
    email = strings.ToLower(email)
    user, err := us.Create(name, email)
    if err != nil {
        return nil, err
    }
    err = ec.Welcome(name, email)
    if err != nil {
        return nil, err
    }
    return user, nil
}
```

We can mock out the UserStore entirely and avoid any SQL setup - we just return a user when we want to tests a successful situation and return an error when we want to test an error case

**External APIs**
Email clients are similar, but image you are using an API to order postgage labels for your packages

What happens if the shipping company doesn't offer a test API, so you can't actually run tests with that API integration? Use mocking! 

The API has a test env but it doesn't match production

Or test env has limitattions and can't be hit as often as devs are hitting tests

Or using a real API is just too slow and you don't want that many network calls, or you want to test without an internet connection

**We can verify that specific behavior occurs**
Signup example again:

```go
func Signup(name, email string, ec EmailClient, us *UserStore) error {
    email = strings.ToLower(email)
    user, err := us.Create(name, email)
    if err != nil {
        return nil, err
    }
    err = ec.Welcome(name, email)
    if err != nil {
        return nil, err
    }
    return user, nil
}
```

We might use mocking and an actual mock - one that tracks which methods are called - to verify that we actually call the Welcome method in somesituations but DO NOT in others

## Example of Faking APIs (package stripe)
- demov0: Shows real example usage of stripe API
- demov1: Shows test examples of faking API

# Interface test suites (package suite)
- test multiple interface implementations using the same suite of tests
- IRL eg, function accepts io.Reader (a file, network connection, etc) and whatever the Reader is, the function does the correct thing
- Test for propert implementation: 
```go
// Create variable of type interface, be implemented by type stub.UserStore : compile test
var _ suite.UserStore = &stub.UserStore{}
```

- sometimes each implementation needs its own setup and teardown
  * handle setup and teardown in userstore.UserStore() test function
  * have suitetest.UserStore accept beforeEach & afterEach

# Testing Subprocesses (package sub)
- Consider: subprocesses (execs) might not exist. Create a flag. If the subprocess doesn't exist, skip the test
- Consider: subprocesses might depend on state. Can copy dir/file to temp dir/file