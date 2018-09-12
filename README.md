# golangJWT
JWT Authorization token with golang


##imports
	"github.com/codegangsta/negroni"
  	"github.com/dgrijalva/jwt-go"
  	"github.com/dgrijalva/jwt-go/request"

##example

##login
localhost:8080/login
{
	"username":"mauran",
	"password": "p@ssword"
}

response

{"token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21Vc2VySW5mbyI6eyJOYW1lIjoibWF1cmFuIiwiUm9sZSI6Ik1lbWJlciJ9LCJleHAiOjE1MzY3NTc0NjIsImlhdCI6MTUzNjc1Mzg2Mn0.DPglxfGpoRirtTHt9wtm5EBSY0bzEBtOoATA6xlZk1M6E8rVyRj456BNAeCquGXR9EWjw6ib9L0KjfQKDtARp8HGRScm9r1F4sF5jCdXbB5zFDu3m_PBkrTEVwTdLNOVoXWhkxSaO3m3Ym-ildnchH9U8HAXLsqj0JOzBcb_6Vg"}

##resource

localhost:8080/resource

Authorization  bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21Vc2VySW5mbyI6eyJOYW1lIjoic29tZW9uZSIsIlJvbGUiOiJNZW1iZXIifSwiZXhwIjoxNTM2NzU1MTUxLCJpYXQiOjE1MzY3NTE1NTF9.Uxoc5OCe9pIg3-SrGhEFhfa4UL9x0Aw0lppeJwMjy8e7YhqeTjmjFRpbEl3IliWktjTMQPTPTIrH68F-fDJ0Nyk-71hgskoDY2pQjPC693nNZaKbJAgLJZH6Yjy94MnKbuSqEf6kmNqetvnRk6xsaVaPvGtBo9mqO36yHall_mA