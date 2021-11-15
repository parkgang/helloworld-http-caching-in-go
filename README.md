# Overview

> etag와 관련된 http 캐시 서버를 구현해보자

1. [Learning HTTP caching in Go](https://www.sanarias.com/blog/115LearningHTTPcachinginGo) 를 참고하여 진행되었습니다.
1. HTTP의 etag 매커니즘이 궁금해서 직접 구현하게 되었습니다.

# Feature

1. 앱의 리소스를 캐싱하고 있는지 확인하려고 합니다.
1. `/black/` URL이 호출 될 때 이미지를 생성하며 그 이미지를 캐시하려고 합니다.
