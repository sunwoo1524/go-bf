# GO-BF
Go로 만들어진 Brainfuck 인터프리터

## 사용법
1. 다운로드
   
   다음 명령어로 소스를 다운로드 합니다.
   ```
   git clone https://github.com/sunwoo1524/go-bf.git
   ```
   또는 ZIP 파일로 다운로드하고 압축을 풉니다.

2. 실행
   
   Windows를 사용하고 있다면, 다음 명령어로 실행할 수 있습니다.
   ```
   ./go-bf <filename>
   ```
   \<filename>은 실행할 Brainfuck 파일의 이름입니다.

   만약 다른 운영체제를 사용하고 있다면, 직접 컴파일할 수도 있습니다.
   ```
   go build
   ```
   또는 컴파일 없이 바로 실행할 수도 있습니다.
   ```
   go run .
   ```
   Go가 없다면 설치할 필요가 있습니다.

## 참고
- 최적화를 진행하지 않기 때문에 최적화가 되어있는 타 구현체에 비해 조금 느릴 수 있습니다.
- example 폴더에서 다양한 Brainfuck 예제들을 만날 수 있습니다.
- 인터프리터를 작동시키는 주요 코드들은 brainfuck/interpreter.go에 있습니다.
- 버그가 있다면 구체적인 내용과 함께 꼭 제보 부탁드립니다!