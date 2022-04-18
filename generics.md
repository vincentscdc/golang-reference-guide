# Generics

don't! Or think about it twice.

refer to [this article](https://go.dev/blog/when-generics)

quoting Ian Lance Taylor, in this article...

```text
If you find yourself writing the exact same code multiple times, where the only difference between the copies is that the code uses different types, consider whether you can use a type parameter.

Another way to say this is that you should avoid type parameters until you notice that you are about the write the exact same code multiple times.


```
