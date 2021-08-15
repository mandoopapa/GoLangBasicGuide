# Go Lang 고루틴(Go Routine)

Go Lang에서는 동시에 실행되는 작업을 Go Routine이라고 부릅니다.

다른 프로그래밍언어에서 쓰레드(Thread)라고 부르는 것과 유사한 개념이지만, Go Routine은 OS 쓰레드와 1 대 1로 대응되지 않고 훨씬 적은 OS 쓰레드를 사용합니다. 따라서 한 번에 더 많은 Go Routine을 실행할 수 있습니다.

동시성(Concurrent)은 싱글 코어의 경우에서 스레드를 여러 개 만들어서 마치 동시에 실행되는 것'처럼' 보이도록 하는 것입니다.

병렬성(parallel)은 멀티 코어에서 실제로 나눠서 작업을 동시에 처리하는 것을 말합니다.

동시성을 구현하기 위해 사용되는 개념이 멀티프로세싱과 멀티쓰레딩입니다.

프로세스는 운영체제로부터 자원을 할당받는 작업의 단위이며, 쓰레드는 프로세스가 할당받는 자원을 이용하는 실행의 단위를 말합니다.

동시성을 지원하도록 작성된 프로그램은 여러 작업을 동시에 실행하는 병렬성도 지원할 수 있습니다. 이를 통해 프로그램의 속도를 크게 향상시킬 수 있습니다.

Go Routine은 사용하기도 매우 간편합니다.

```go
go 함수명()
```

