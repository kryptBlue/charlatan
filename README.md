我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Charlatan

[![Circle CI](https://circleci.com/gh/percolate/charlatan.svg?style=svg)](https://circleci.com/gh/percolate/charlatan)
[![codecov.io](https://codecov.io/github/percolate/charlatan/coverage.svg?branch=master)](https://codecov.io/github/percolate/charlatan?branch=master)
[![BSD](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/percolate/charlatan/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/percolate/charlatan)](https://goreportcard.com/report/github.com/percolate/charlatan)

Percolate's Go Interface Mocking Tool.  Please read our [introductory blog post](https://medium.com/percolate-engineering/introducing-charlatan-df9b5d3d3107).

## Installation

    go get github.com/percolate/charlatan

## Usage

```
  charlatan [options] <interface> ...
  charlatan -h | --help

Options:

  -dir string
        input package directory [default: current package directory]
  -file value
        name of input file, may be repeated, ignored if -dir is present
  -output string
        output file path [default: ./charlatan.go]
  -package string
        output package name [default: "<current package>"]
```

If you would like the mock implementations to live in the same package
as the interface definition then use the simplest invocation as a
directive:

    //go:generate charlatan Interface

or from the command line:

    charlatan -file=path/to/file.go Interface

You can chose the output path using `-output`, which must include the
name of the generated source file.  Any intermediate directories in the
path that don't exist will be created.  The package used in the
generated file's `package` directive can be set using `-package`.

## Example

Given the following interface:

```go
package example

//go:generate charlatan Service

type Service interface {
	Query(filter *QueryFilter) ([]*Thing, error)
	Fetch(id string) (*Thing, error)
}
```

Running `go generate ...` for the above package/file should produce
the file `charlatan.go`:

```go
package example

type QueryInvocation struct {
	Parameters struct {
		Filter *QueryFilter
	}
	Results struct {
		Ident1 []*Thing
		Ident2 error
	}
}

type FetchInvocation struct {
	Parameters struct {
		Id string
	}
	Results struct {
		Ident3 *Thing
		Ident4 error
	}
}

type FakeService struct {
	QueryHook func(*QueryFilter) ([]*Thing, error)
	FetchHook func(string) (*Thing, error)

	QueryCalls []*QueryInvocation
	FetchCalls []*FetchInvocation
}

func (f *FakeService) Query(filter *QueryFilter) (id1 []*Thing, id2 error) {
	invocation := new(QueryInvocation)
	invocation.Parameters.Filter = filter

	id1, id2 := f.QueryHook(filter)

	invocation.Results.Ident1 = id1
	invocation.Results.Ident2 = id2

	return
}

// other generated code elided ...
```

Now you can use this in your tests by injecting the `FakeService`
implementation instead of the actual one.  A `FakeService` can be used
anywhere a `Service` interface is expected.

```go
func TestUsingService(t *testing.T) {
	// expectedThings := ...
	// expectedCriteria := ...
	svc := &example.FakeService{
		QueryHook: func(filter *QueryFilter) ([]*Thing, error) {
			if filter.Criteria != expectedCriteria {
				t.Errorf("expected criteria value: %v, have: %v", filter.Criteria, expectedCriteria)
				return nil, errors.New("unexpected criteria")
			}
			return expectedThings, nil
		},
	}

	// use the `svc` instance in the code under test ...

	// assert state of FakeService ...
	svc.AssertQueryCalledOnce(t)
}
```

Create anonymous function implementations for only those interface
methods that should be called in the code under test.  This will force
a panic if any unexpected calls are made to the mock implementation.

The generated code has `godoc` formatted comments explaining the use
of the mock and its methods.
