## Installation
```bash
go get weavelab.xyz/wiggum
```

For more information on `weavelab.xyz`, see the projects [readme](https://github.com/weave-lab/ops-xyz/blob/master/README.md).

Overview
======

![wiggum](assets/wiggum.png)    
Wiggum is an OAuth2 middleware that uses [JWT](http://jwt.io) to support [SSO](http://en.wikipedia.org/wiki/Single_sign-on).
It follows the pattern outlined by Google for [OAuth2 for Client-side Applications](https://developers.google.com/accounts/docs/OAuth2UserAgent).  

Wiggum checks for a cookie, or header and extracts the JWT token and validates its claims.  If the JWT is not present or invalid, it will reject the request.  It also provides convenience methods for extracting token data (claims).

The cookie name, validation secret, validation path(s), and rejection protocol are currently set in config.go.

Dependencies
-------
* [negroni](https://github.com/codegangsta/negroni)
* [jwt-go](https://github.com/dgrijalva/jwt-go)
* SSH keys

    openssl genrsa -out [app].rsa 1024 // 2048 or 4096
    openssl rsa -in [app].rsa -pubout > [app].rsa.pub


Example Consumer/Client
------
your app main.go 

    // insert wiggum when setting up negroni
    n := negroni.Classic()
    n.Use(wiggum.New(tokenDecryptKey, except, redirect))
    n.UseHandler(...)
    n.Run(":8080")

config.go

    //Set the values for
    cookieName      = "nameOfAuthCookie"
    tokenDecryptKey = shared.rsa.pub 
    except          = []string{"/public", "/non", "/restricted", "/paths", "/prefixed/path/has/*"}
    redirect        = "/" // or "" for {"error":"invalid_token"}

* cookieName* => the name of the cookie that contains the JWT token 
* pubKeyPath => path to the public rsa key that matches the private rsa key used to create the token
* except     => a slice of strings that contains the paths that should not be validated, i.e. skipped by the middleware - use * at end for prefix only
* redirect   => where to redirect the user if an invalid token is passed.  If blank (""), a JSON error response will be provided
 
 *not required if passing JWT token in header: Authorization: Bearer [JWT String]


handlers.go

    //access JWT data (claims)
    userID := wiggum.TokenRequest(r *http.Request).UserID()  // string
    claims := wiggum.TokenRequest(r *http.Request).Claims()  // map[string]interface{}


Example Provider (see [Fallout-Boy](https://github.com/weave-lab/fallout-boy))
------
Creating a JWT token

    // to create a new JWT token
    tokenString, err := wiggum.Make(acls, userID, username, currentPractice, exp, buff, signKey)

* acls             => string of comma separated numbers, i.e "{1,4,3,6}"
* userID           => user id
* username         => user name
* currentParactice => current practice slug
* exp              => int64 representing token expiration time, i.e. time.Now().Add(time.Minute * 15).Unix()
* buff             => int64 representing token refresh buffer, i.e. time.Now().Add(time.Minute * 120).Unix()
* signKey          => the private rsa key used to generate client rsa public keys

Managing cookies

    // logging in
    wiggum.SetCookie(tokenString, w, r)
    // logging out
    wiggum.DeleteCookie(w, r)

More
------
For examples on how a client app should obtain/refresh a JWT token please see [Fallout-Boy](https://github.com/weave-lab/fallout-boy)
