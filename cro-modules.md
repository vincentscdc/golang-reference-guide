# Internal golang modules

Links to the different common golang modules used in crypto.com

* Configuration (with AWS secrets)
* HTTPTypedWrapper
* Tracing HTTP middleware

## Choosing the appropriate external module

A good module:

* Follows idiomatic go
* Has unit tests
* Has examples
* Is used by a lot of gophers and golang OSS (follow github stars)
* Is documented
* Is sharply focused
* Is readable
* Has not so many dependencies

😈: look at this lib, it's performing better the one you have selected!
    is it worth the risk of using a non widely used dependency? Is this really the bottleneck in your application?
