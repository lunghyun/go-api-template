# Restapi 구현
## 목표
RESTful API를 알아보자. 그리고 api 서버 만들기
## 내용
- Restful api 정의
- restful api 서버 만들기
- 웹 서버 발전 과정
- gin으로 서버 만들기

## Restful api란?
REST는 자원을 이름으로 구분, 자원 상태를 주고 받는 소프트웨어 아키텍처
Restful api는 rest 규약을 따르는 api를 말함
웹서버에서는 url과 http 메서드로 데이터와 동작을 정의하는 방식을 의미

## 효과
- url 과 메서드를 사용해 데이터, 동작 정의 -> 동일한 방식으로 처리
- fe, be가 분리 -> restapi가 보편화, 데이터 프로바이더로써 웹 서버 역할이 중요
- restful api를 사용 -> 여러 서버가 동일한 인터페이스로 통신

# 과정
1. gorilla/mux 같은 restful api 웹서버 지원 패키지 설치
2. restful api에 맞추서 웹 핸들러 함수 정의
3. restful api 테스트 코드
4. 웹 브라우저로 데이터 조회

## 대충
GET /student
GET /student/id
POST /student/id
DELETE /student/id

# 시작
- 이전보다 간단하게 패키지를 써보자잉
- gorilla/mux

