# Effective STL

![](https://images-na.ssl-images-amazon.com/images/I/51oGU0vuNuL._SL200_.jpg)

### Metadata

- Author: Scott Meyers
- Full Title: Effective STL
- Category: #books

### Highlights

- To avoid potential parsing ambiguities (the details of which I’ll spare you), you are required to use typename to precede type names that are dependent on formal type parameters. ([Location 455](https://readwise.io/to_kindle?action=open&asin=B0019HW0K6&location=455))
- Contiguous-memory containers (also known as array-based containers) store their elements in one or more (dynamically allocated) chunks of memory, each chunk holding more than one container element. ([Location 554](https://readwise.io/to_kindle?action=open&asin=B0019HW0K6&location=554))
- You should prefer the construct using empty, and the reason is simple: empty is a constant-time operation for all standard containers, but for some list implementations, size may take linear time. ([Location 807](https://readwise.io/to_kindle?action=open&asin=B0019HW0K6&location=807))
- Whenever you have to completely replace the contents of a container, you should think of assignment. ([Location 862](https://readwise.io/to_kindle?action=open&asin=B0019HW0K6&location=862))
