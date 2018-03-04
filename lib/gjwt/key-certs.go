/*
* Go Library (C) 2017 Inc.
*
* @project     Test
* @package     main
* @author      @jeffotoni
* @size        01/06/2017

* @description Our main auth will be responsible for validating our
* handlers, validating users and will also be in charge of creating users,
* removing them and doing their validation of access.
* We are using jwt to generate the tokens and validate our handlers.
* The logins and passwords will be in the AWS Dynamond database.
*
* $ openssl genrsa -out private.rsa 1024
* $ openssl rsa -in private.rsa -pubout > public.rsa.pub
*
 */

package gjwt

const (
	RSA_PRIVATE = `-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDROesLPLdNRbGVdUo8FXFD3a4Jmlg5zlt2RUDq7x3xz+0HXMXX
z4a+7PHa8Vt7FuvaQ6xNbgOJZb4kgfT0cCBkrYPLuNmDdOqOSsiQqnp3gckYInku
oTQSq1o6TfH6dO6UvVBxARPY7oOob7aXKbvq1DpZ9FcLaPL53C+2f4yh5QIDAQAB
AoGBAImVZJLHImKV6ek2b9KC5zCRndiCvnGE2XA0qjPTegWBjYTB5Pe9aAY1GfW3
sUiIEiA9UBi6t2iqlxa6vrHe6e0aKpZsJY3Yl3DzJ1OpJ+2jxZsTDa0u1/w0PkYs
540Ujb2glPjv2RMUk0+/77a9V6a6Y5WspumxMfzLf1KpkvMhAkEA6M0adHodq9fK
qftbBRbI0IvWeEm384wlXotFr4A793biWHvbwwnsJUNa52F1rtK3QtLVs7cf1nmI
ViJhghj7ZwJBAOYTZ2z67LX8UyDrM+6xuGmOnuYO9Q88LUlrbjVW7hUj81GRkXWg
mWY5q/BokzehLoS/9E62H2r2P2cvmyyntNMCQQDb3hpGN68eRUgbElH8lHBExk1g
ff9F/e6tREwkXLBGH/nWJ+R+aDinWN3Z+anz2v9KFWXPvfxLy6x/7Si6fm2xAkEA
wHS4zlOz0KNqUh4NSfzF2x6fUqhuW7kl1MWV4e5+p5Z0AeZ3u+KEjxts9WQ0yZL9
m7QKRBFahJnJ/aG39wmuWwJADyaoxNp4fmJtIirTDkbVHdk02Z2eb3PB5yJFRnzo
OR5nqVfpWOxKbvx1lMMHcvAZP9jngHcza1b6OCgWmZPgPw==
-----END RSA PRIVATE KEY-----
`

	RSA_PUBLIC = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDROesLPLdNRbGVdUo8FXFD3a4J
mlg5zlt2RUDq7x3xz+0HXMXXz4a+7PHa8Vt7FuvaQ6xNbgOJZb4kgfT0cCBkrYPL
uNmDdOqOSsiQqnp3gckYInkuoTQSq1o6TfH6dO6UvVBxARPY7oOob7aXKbvq1DpZ
9FcLaPL53C+2f4yh5QIDAQAB
-----END PUBLIC KEY-----`
)
