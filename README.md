# Phoenix [![GoDoc](https://godoc.org/github.com/go-phoenix/phoenix?status.png)](http://godoc.org/github.com/go-phoenix/phoenix) [![Travis CI build status](https://travis-ci.org/go-phoenix/phoenix.svg)](https://travis-ci.org/go-phoenix/phoenix)

A code generator for Go.  See
[the Blog Post](https://blog.synapsegarden.net/update/phoenix/).

## What is Phoenix?

Phoenix is a Go code generator, meant to be used with packages from
https://github.com/go-phoenix to build interactive OpenGL applications.  Its
purpose is to cleanly support a simple subset of the sorts of things one might
use an in-language template engine for.  This allows static code to be generated
for each use-case of the types implementing a Go interface, for a particular set
of interfaces.

At the beginning, it will implement only a small set of types, and extensions
must be manually added by the user.  As we push development forward, we have
high hopes for what we might accomplish.

## How can I use it?

Phoenix can be used directly or through `go:generate` (which is how it's meant
to be used) by adding a comment to Go source code as explained in
[the go:generate Blog Post](https://blog.golang.org/generate).

For example:

```go
//go:generate phoenix ... [TODO: Finish me!]
```

## Why GNU Affero General Public License?

Phoenix is intended to run as a standalone utility which can be driven by other
programs, including ones offering network interfaces.  Since Phoenix is free
software, its licensing supports the freedom for users to obtain and use its
source code.

If the freedom to obtain and use its source code could be taken away by hosting
the program behind a network interface, then the software wouldn't really be
free.

The GNU Affero GPL specifies that anyone who takes and redistributes the
software, with or without modifications, must pass along the freedoms it
specifies to their users.

We think this is a good thing.  So do lots of other people, including e.g.
MongoDB.

#### What does this mean for me?

You don't need to use a free license for your software unless you create a
linked version importing free-licensed code.  You do, however, need to make sure
your users can obtain a copy of the Phoenix source code, including any
modifications you've made.

---

For more info, see the article
[What is Copyleft?](http://www.gnu.org/licenses/copyleft.en.html)
and
[the difference between an “aggregate” and other kinds of “modified versions”](http://www.gnu.org/licenses/gpl-faq.html#MereAggregation)
on [the GNU organization website](http://www.gnu.org).
