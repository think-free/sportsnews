Notes
----------------

1. logging and restilog are libs that I wrote for another project and are imported to be used in this challenge
- logging is a wrapped around "go.uber.org/zap"
- restilog is a wrapper around "gopkg.in/resty.v1" that add tracing of the request and response for debugging purpose

2. datamodel and icmongodb are used in all the microservices implemented (providers/htafc and sportsnews)
- datamodel is the model saved in the database to store the articles
- icmongodb is used to connect to mongodb and create the database/collection
