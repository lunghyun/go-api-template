# TODO

1. [x] dto, model 분리
2. [x] binding 태그로 validation 변경
3. [ ] ctx 매개 변수로 추가
4. [ ] 공통 errors 패키지 도입
5. [ ] 테스트 코드 작성
6. [ ] postgreSQL, MySql, mariaDB 구현체 구현 및 라우팅 시 생성자 변경
7. [ ] Student와 연결성 있는 도메인 구상 및 구현
8. [ ] 검증 후 다른 브랜치(아키텍처 패턴) 진행
9. [ ] 현재 Students를 타입으로 사용하는 경우가 많은데 ...Student를 쓰는게 나을지 고려

# 지식
## 레포지토리 넣을만한 메소드
```java
// Create
save(entity) // 단일 엔티티 저장(id있으면 저장, 없으면 생성)
saveAll(entities) // 여러 엔티티 한번에 저장

// Read
findById(id) // id로 단일 조회(없으면 optional.empty())
findAll() // 전체 조회(슬라이스로 반환)
findAllById(ids []int) // 여러 ID 한번에 조회 
existsById(id) // 존재 여부만 id로 확(bool)
count() // repo에서 전체 갯수 반환

// Update
save(entity)  // 생성/수정 겸용

// Delete
deleteById(id) // id로 삭제
delete(entity) // 엔티티로 삭제(객체를 이미 가지고 있을떄)
deleteAll() // 전체 삭제 즉 테이블 비우기 (위험함)
deleteAllById(ids []int) // 여러 id 한번에 삭제
```
