# RESTful API Design

![](https://images-na.ssl-images-amazon.com/images/I/51dPKUlNnJL._SL200_.jpg)

### Metadata

- Author: Matthias Biehl
- Full Title: RESTful API Design
- Category: #books

### Highlights

- It is triggered by a POST to the controller resource, which returns a status code 202 Accepted. ([Location 1240](https://readwise.io/to_kindle?action=open&asin=B01L6STMVW&location=1240))
- If a resource is updated with the PATCH method (see section 7.5.2.5), the input data is sent as a difference representation via the HTTP body. ([Location 1638](https://readwise.io/to_kindle?action=open&asin=B01L6STMVW&location=1638))
- Both PATCH and PUT are used for updating a resource; what is the difference between the two? PUT is used for a complete update, and the complete resource needs to be sent to the API. Thus PUT can also be used for creating new resources. PATCH is used for a partial update, only changed fields of the resource need to be sent to the API. Thus PATCH can only be used for updating and not for creating a resource. ([Location 1887](https://readwise.io/to_kindle?action=open&asin=B01L6STMVW&location=1887))
- HTTP methods (GET, POST, PUT, DELETE, ([Location 2206](https://readwise.io/to_kindle?action=open&asin=B01L6STMVW&location=2206))
