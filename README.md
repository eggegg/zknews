1.UsersService
2.SportNewsService 
3.FashionNewsService
4.PoliticsNewsService
5.NewsOrchestrator
6.SearchService
7.RecommendationService
8.EventStore




start nats-streaming
docker run -d -p 4222:4222 -p 8222:8222 nats-streaming

start elasticsearch
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.2.4
curl http://localhost:9200 for test


protoc ./eventstore.proto --go_out=plugins=grpc:./pb
protoc --go_out=plugins=grpc:. *.proto

NewsOrchestrator:
    http: 
        GetAllNews: call grpc of *NewsService get all news
        GetAllNewsByType: call grpc of oneNewsService
        CreateNews: call grpc of NewsService (newsType, newsId)
        SearchNews: call grpc of SearchService
        RecommendNews: call recommend news

PoliticsNewsService
    grpc: 
        CreateNews: save to db,publish event
        GetNewsById: select from db 
        GetAllNews: select from db
    pub:
        watch-news


SearchService
    grpc:
        GetAllNews:(skip, from, keyword)
    sub:
        create-news: insert into 

Recommendation
    sub:
        watch-news:
