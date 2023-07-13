# Surly
Url Shortener tiers

|free|plus|
|----|----|
|10 links/mo|200 links/mo|
|link-bio   |link-bio|
|           -|edit links|
|           -|password protection|
|           -|dataClick (device origin, clicks per timespan) |

## Histories-ToDo
- Change Password 🚧
- Forgot password 🚧
- Reset forgot password 🚧

- Create new short URL (10 per month free)
- Multi refresh token
- iupi editor       // modify shorts
- iupi bux          // create shorts in bulk, export shorts
- iupi 4you         // custom shorts
- DataClick
- iupi collections  // equal to link-bio
- Patreon
### Histories-Done
- Create new user
- Confirm registration
- Login user
- Refresh token

<p>A percentage of the revenue will go towards helping stray cats and dogs. Expect good news soon. 😽 🐶</p>

## Tooling
- Swagger 2.0
    swag init -g ./security/cmd/main.go
- MySQL
- github.com/devfeel/mapper
- github.com/go-playground/validator/v10

- Open Telemetry

## Folder structure
- domain [entities, ports]
- application (CQRS)
- presentation [controller, models]
- infrastructure [adapters, data]

__"Marker interface"__
```

    u := entities.User{}
	var a interface{} = u
	auditable, ok := a.(intf.IAuditable)
	fmt.Println(auditable, ok)
```
## Color Palettes
https://www.color-hex.com/color-palette/1023031

## tech
- docker
- go
- jwt
- mysql
- serverless
- https
- aws

## domains

`surly-security.com`: tokens issuer  
`surly-eco.com`: Surly Ecosystem, audience

## Docker
1. docker build -t surly-security .
2. docker container run -i --name surly-alpha -e Jwt:Key=$ecre7 -e APP_PORT=6001 -e DB_SERVER=localhost -e DB_NAME=surly_security -e DB_PORT=3306 -e DB_USR=root -e DB_PWD=eder -dp 6001:6001 surly-security 
## AWS Elastic Container Registry
1. download and install AWS CLI
2. In Console, get access keys for CLI
3. configure AWS credencials access keys - aws credentials
--------------------------
1. aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 948022990339.dkr.ecr.us-east-1.amazonaws.com
2. docker build -t surly-security .
3. docker tag surly-security:latest 948022990339.dkr.ecr.us-east-1.amazonaws.com/surly-security:latest
4. docker push 948022990339.dkr.ecr.us-east-1.amazonaws.com/surly-security:latest


## About iupi.io
Hi my name is Eder a software developer and enthusiastic on Golang language.

I made yupi.dev because I wanted to create a new player in this world of url shorteners.
This project is my baby (passion project/brainchild) and I will continue working hard on it, building more features for you.

If you would like to say thanks, I would really appreciate you buying me a coffee.
Include your email address in the description if you want to get supporter benefits.


It's my will maintening (I intend to keep) this site free but if I have to monetise, supporters will get something for free as a thank you.


