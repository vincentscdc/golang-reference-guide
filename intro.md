# Introduction

## About this guide

This is a constantly evolving guide about go in crypto.com
You should and MUST add to it.

I use a little companion "😈" that gives bad opinions, to understand the choices made in this guide.

## About core principles

Idiomatic Go is:

1. Readable (magic is bad)
2. Simple
3. Lean/Minimal

## About frameworks

Unlike most other languages, gophers (that's what a golang practicioner is called) tend to avoid frameworks.

Why? Go favors composition over inheritence, and composing with libs instead of sticking to a framework stick to this.
Just throw your best libs for each part of your application, respecting the interfaces of the stdlib or most common go idioms.

On top of that, our go services are deployed on Kubernetes, which has a lot of functionalities that these frameworks can offer: discoverability, load balancing, tracability, resilience, rate limiting...

😈: but look, this fiber framework is top in techempower benchmarks!
    sure, but it's using non idiomatic go, even unsafe code. It won't probably benefit the next go version, it might even break.

😈: but look, there is this all-in-one go-kratos/go-zero/go-kit...!
    sure, but if you look carefully, these are maintained by a very small amount of people, do you really want your business at the mercy of 1 person?
    It does not mean you should not go there and look how they do things! They usually are very good references.
