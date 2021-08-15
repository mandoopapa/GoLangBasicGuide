# Go Lang 구조체 (Structure)

Go Lang의 구조체(Structure)는 필드를 하나의 개념으로 묶는 것을 의미합니다.

Go Lang에는 타 언어에서 사용하는 Class와 상속을 지원하지 않습니다. 

대신 메소드를 구조체에 연결하는 방식을 통해 Class의 형태를 구현합니다. 

이 경우 메소드는 구조체안에서 정의되는 것이 아니라 구조체 밖에서 정의되어 사용합니다.

구조체의 형태

```go
struct { field1 string field2 int }
```

구조체 내부는 각각의 필드명과 필드 타입으로 구성됩니다.