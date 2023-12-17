# http-cache-server-with-go

> http 캐시 서버를 구현해보자

[Learning HTTP caching in Go](https://www.sanarias.com/blog/115LearningHTTPcachinginGo) 를 참고하여 진행되었습니다.

`HTTP` 의 `etag` 매커니즘과 동시에 캐시서버의 동작이 궁금해서 직접 구현하게 되었습니다.

## Start

`go run main.go` 으로 서버를 실행한 후 [http://localhost:8080](http://localhost:8080) 로 들어가서 요청을 날려보시면 됩니다.

개발자 도구에서 네트워크 탭을 열고 한번 확인해보세요!
