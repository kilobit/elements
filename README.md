Elements
========

Build markup documents out of simple objects in Golang.

Inspired by the now retired [Apache ECS](http://jakarta.apache.org/ecs/).

Status: In development.

```{.golang}
	el := NewElement("tag").
		AddAttr("tagged").
		SetAttr("foo", "bar").
		AddChild(NewElement("nav")).
		AddChild(NewElement("main"))
		
	fmt.Println(el)	
```

Features
--------

- Build markup without resorting to string manipulation.
- Add any Stringer type.

Roadmap:
- Use Formatters to change between output types (html, xhtml, html5).
- Built-in library of HTML5 elements.
- Use Nodes directly as http.Handlers.

Installation
------------

```{.bash}
go get kilobit.ca/go/elements
```

Building
--------

```{.bash}
cd kilobit.ca/go/elements
go test -v
go build
```

Contribute
----------

Please submit a pull request with any bug fixes or feature requests
that you have.  All submissions imply consent to use / distribute
under the terms of the LICENSE.

Support
-------

Submit tickets through [github](https://github.com/kilobit/elements).

License
-------

See LICENSE.

--  
Created: Oct 1, 2019  
By: Christian Saunders <cps@kilobit.ca>  
Copyright 2019 Kilobit Labs Inc.
