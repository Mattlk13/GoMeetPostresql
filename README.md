# Go Meet DB - A wrapper for interacting with databases in pure Go

![screenshot](http://i.imgur.com/VxDboZW.png)

This is a small wrapper for interacting with a database in Go, it is a great way to quickly get your project connected to a database and output queries in a human-readable format. 

A simple command line interface has been implemented, however this is not essential for the wrappers use. Simply strip away the interface and integrate the required functions into your project.

As an example, the `lib/pq` driver has been used to connect to a PostgreSQL database. However, you can swap out the driver and connect this wrapper to another database, such as MSSQL.

`go run gomeetdb.go`

Type `-h` for help


## MIT License

Copyright (c) 2017 Rohan Likhite

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


