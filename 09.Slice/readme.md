# Go Lang 슬라이스(Slice)

Go Lang의 슬라이스(Slice)는 값을 추가하여 확장가능한 자료구조로 크기가 고정되어 있지 않고 동적으로 배열의 크기를 증가시킬 수 있습니다.

슬라이스는 Go Lang 내장 함수인 make를 통해 공간을 확보할 수 있습니다.

make 함수는 (타입type, 길이length, 용량capacity)를 선언하여 사용합니다. 여기서 용량은 내부배열의 최대길이를 의미합니다.

슬라이스 선언은 배열(Array)과 거의 비슷합니다. 다만, []안에 길이 지정을 하지 않습니다.

