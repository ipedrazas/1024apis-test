# 1024apis - Test framework


When you create a distributed system using [1024apis](https://github.com/ipedrazas/1024apis) you need an easy way to create traffic to analyse what's going on.

This project aims to provide the testing artifacts necessary for that.

You have to define the following environment variables:

* WAIT_TIME: this is the time the test waits between series of requests.
* NUM_APIS: number of services deployed, or that wanrt to be tested.

Remember that this test framework only works with 1024apis for now.
