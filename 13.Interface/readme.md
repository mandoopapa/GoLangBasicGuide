# Go Lang 인터페이스(Interface)

Go Lang의 인터페이스(interface)는 특정 값이 가지고 있기를 기대하는 메소드의 집합, 동작을 수행할 수 있는 타입이 지녀야하는 동작들의 집합입니다.

이 Go Lang에서는 인터페이스를 통한 객체의 상태와 기능을 디커플링(Decupling)시켜 object에 종속되어있던 기능을 종속해제합니다. 

디커플링을 통해 기존에 OOP의 단점인, 클래스마다 차지하고 있는 기능들이 메모리를 많이 차지하던 문제에서 효율성을 추구할 수 있게 합니다.


인터페이스의 형태

 
```go
type 인터페이스명 interface {
	메소드명()
	메소드명(매개변수타입)
	메소드명() 반환값타입
}
```

인터페이스 정의에 나열된 모든 메소드를 가진 타입은 해당 인터페이스를 만족한다고 합니다. 

인터페이스를 만족하는 타입은 해당 인터페이스가 필요한 모든 곳에서 사용할 수 있습니다.

인터페이스를 만족하려면 인터페이스에 정의된 메소드명, 매개변수 타입, 반환값 타입이 모두 일치해야합니다. 

인터페이스 정의에 나열된 메소드는 반드시 모두 구현해야하며 하나라도 구현하지 않으면 인터페이스를 만족할 수 없습니다
