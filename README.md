# download

Lazypic에서 사용하는 AWS S3에서 파일을 다운로드 하는 명령어입니다.
이미지파일, 블랜더파일을 다운로드해서 렌더링하기 위해 사용합니다.

#### 다운로드, 셋팅
- [Download](https://github.com/lazypic/download/releases)
- Lazypic은 기본적으로 `ap-northeast-2`(서울리전), `lazypic` S3 버킷을 사용합니다.
- 이 프로그램이 작동되기 위해서는 S3 읽기 권한이 있는 ~/.aws/credentials 파일이 셋팅되어있어야 합니다.

#### 사용법
- pilot 프로젝트의 모든 파일을 다운로드하기.
```bash
$ download -prefix pilot
```

- 파일이 존재하면 파일 리스트를 출력하고 다운로드 할지 물어보게 됩니다.
```bash
pilot/ep1/s1/c1.blend
pilot/ep1/s1/c2.blend
pilot/ep1/s1/c3.blend
Download? (y/n):
```

- `y`를 타이핑하면 명령어를 실행한 경로에 아래 파일들이 다운로드 됩니다.
```bash
./pilot_ep1_s1_c1.blend
./pilot_ep1_s1_c2.blend
./pilot_ep1_s1_c3.blend
```

- 경로를 유지하고 싶다면 아래처럼 -subdir 옵션을 붙혀주세요.
```bash
$ download -prefix pilot -subdir
```
- 아래처럼 각 파일이 각경로에 저장됩니다.
```bash
./pilot/ep1/s1/c1.blend
./pilot/ep1/s1/c2.blend
./pilot/ep1/s1/c3.blend
```

- 응용1 : pilot 프로젝트중 에피소드1 모든 파일 다운로드하기.
```
$ download -prefix pilot/ep1
```

- 응용2 : pilot 프로젝트중 에피소드1, scene1 파일을 다운로드하기.
```
$ download -prefix pilot/ep1/s1
```

- 응용3 : pilot 프로젝트중 에피소드1, scene1, cut1 파일을 다운로드하기.
```
$ download -prefix pilot/ep1/s1/c1
```

#### 직접 컴파일하기
- Go가 설치되어있다면, 터미널을 통해서 아래처럼 설치할 수 있습니다.
```bash
$ go get -u github.com/lazypic/download
```

#### 라이센스
- AWS API는 Apache License 2.0 라이센스 정책을 따르기 때문에 이 코드도 같은 라이센스를 사용합니다.
