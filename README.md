# Overview

> http 캐시 서버를 구현해보자

1. [Learning HTTP caching in Go](https://www.sanarias.com/blog/115LearningHTTPcachinginGo) 를 참고하여 진행되었습니다.
1. HTTP의 etag 매커니즘과 동시에 캐시서버의 동작이 궁금해서 직접 구현하게 되었습니다.
1. `go run main.go` 으로 서버를 실행한 후 [http://localhost:8080](http://localhost:8080) 로 들어가서 요청을 날려보시면 됩니다.
   > 개발자 도구에서 네트워크 탭을 열고 한번 확인해보세요!
1. 해당 제품의 사용법과 관련된 블로그 글은 [간단한 HTTP 캐시 서버를 만들어보자 (with golang)](https://parkgang.github.io/golang/lets-create-an-http-cache-server-with-golang) 에 작성되어있습니다.
