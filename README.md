Sports News
=============================

### Build the application 

docker-compose up --build

### Structure

The application is composed of serveral microservices :

- **mongodb**
    The database where we store the articles from news providers

- **sportsnews**
    In charge of retreiving articles from the database and presenting them in json format

- **providers/**
    Several news providers can be implemented, we have one 'htafc' that retreive news from [Huddersfield Town](https://www.htafc.com/api/incrowd/getnewlistinformation)

### API Usage 

- http://localhost:8080/provider/realise/v1/teams/Huddersfield%20Town/news : To retreive all the news for **Huddersfield Town**

- http://localhost:8080/provider/realise/v1/teams/Huddersfield%20Town/news?page=0 : To retreive new with paging (incremente page number to retreive older news)

- http://localhost:8080/provider/realise/v1/teams/Huddersfield%20Town/news/608677 : To retreive a specific news

### Missing 

- Tests are only implemented to ensure that the application is working and to help development, they should be extended to cover the entire code but are not as this is an exercice.
- **cliparams** and **config** could be merged, config is aimed to parameters in all the code that would not be changed by user and decided at compile time.
- sportsnews/internal/api/responses.go : **TotalItems** is not implemented, I didn't as it's an exercice.
- I didn't used cron as suggested library as the for loop with sleep is simpler for that task for this implementation.
