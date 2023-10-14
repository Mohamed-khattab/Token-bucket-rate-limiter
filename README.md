# Token Bucket Rate Limiter in Go

This repository features a straightforward implementation of a token bucket rate limiter in Go. Designed for managing request rates across different entities, it handles global requests, services A and B, and individual users based on their IP addresses.

## TokenBucket Structure

The `TokenBucket` structure embodies a generic token bucket system, equipped with methods for creating a new token bucket, refilling tokens based on elapsed time, and checking token availability for a request.

## Rate Limiting for Global, Services, and Users

### - Global Rate Limiter

A global rate limiter orchestrates overall request control. The `globalTokenBucket` instance sets the ceiling for requests at 1000 per second.

### -  Service Rate Limiters

Two specialized rate limiters, namely `ServiceATokenBucket` and `ServiceBTokenBucket`, govern request rates for services A and B, respectively.

### - User Token Bucket Manager

The `UserTokenBucketManager` oversees individual token buckets for users, identified by their IP addresses. Each user faces a limit of 2 requests per second.

## Example Usage

The main file (`main.go`) provides illustrative examples showcasing the application of the rate limiter for global, service, and user requests. For seamless integration into your applications, feel free to tailor the provided examples.

In these examples, the global rate limiter permits consumption of up to 500 tokens with a refill rate of 5 tokens per second. Services A and B adhere to respective limits of 50 and 100 tokens, accompanied by refill rates of 10 and 5 tokens per second. Individual user requests are restricted to 20 tokens, replenished at a rate of 2 tokens per second based on their IP addresses.
