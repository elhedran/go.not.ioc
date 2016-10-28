# Golang doesn't need IoC

This is a golang implementation and tests for the project used to demonstrate IoC in:
https://www.youtube.com/watch?v=5lIeky2V4dc

Golang structs do not name the interfaces they inherit from, and nor do
structures have constructors. Both of these make implementing IoC difficult.

However IoC is also unnecessary in golang. This example uses the same pattern
as used in golang's net/http module.

If something needs to be replaceable at runtime, then make it a public
variable.  If that same thing needs a default, then define one, or a function
that provides one.

Its as testable as the IoC example but isn't using IoC.
