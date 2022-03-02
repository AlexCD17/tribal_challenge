# Tribal Backend code challenge

This is a little REST miscroservice created with the least amount of packages outside of the standard library, I decided to go the dependecy injection way while maintaning a minimal aproach.


# Run with docker

```
# docker build -t app .
```
runing it, the app listens on the 8080 port

```
## docker run -it -p 8080:8080 app
```
# Sending request

The app has a single endpoint, '/eval', here's an example of a request using curl:

```
## curl -d '{                                
"foundingType": "STARTUP",
"cashBalance": 2000.30,
"monthlyRevenue": 4435.45,
"requestedCreditLine": 100,
"requestedDate": "2021-07-19T16:32:59.860Z"
}' -H 'Content-Type: application/json' http://localhost:8080/eval

```
