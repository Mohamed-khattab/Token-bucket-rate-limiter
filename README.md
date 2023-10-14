# Token Bucket Rate Limiter in Go

This repository contains a simple implementation of a token bucket rate limiter in the Go programming language. The rate limiter is designed to control the rate of requests for different entities, such as global requests, requests from services A and B, and requests from individual users based on their IP addresses.

## TokenBucket Structure

The `TokenBucket` structure represents a generic token bucket system. It includes methods for creating a new token bucket, refilling tokens based on elapsed time, and checking if there are enough tokens for a request.

## Rate Limiting for Global, Services, and Users

### - Global Rate Limiter

A global rate limiter is implemented to control the overall rate of requests. The `globalTokenBucket` instance limits the rate of requests to a maximum of 1000 requests per second.

### -  Service Rate Limiters

Two rate limiters, `ServiceATokenBucket` and `ServiceBTokenBucket`, are designed to control the rate of requests for services A and B, respectively.

### - User Token Bucket Manager

The `UserTokenBucketManager` manages individual token buckets for users based on their IP addresses. Each user is limited to a rate of 2 requests per second.

## Example Usage

The main file (`main.go`) includes examples demonstrating the usage of the rate limiter for global, service, and user requests. To integrate the rate limiter into your applications, customize the provided examples.

Feel free to explore and modify the code to suit your specific use case.
