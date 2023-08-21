## Fibonacci service
This service store in Redis the Fibonacci numbers and can get any segment 

See **api/api.yaml** for information about allow endpoints and methods

### Build & Run(locally)

##### Prerequisites

* docker
* docker-compose
* make

#### Build
To build a service for a native platform just call

``` sudo docker-compose up -d ```

To build only fibonaccisrv without docker-compose call

``` sudo make build ```


