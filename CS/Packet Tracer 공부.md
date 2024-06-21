라우터: 2620사용
End Device: PC
콘솔 연결

Fast ethernet연결법
커넥터: Copper Cross-Over(점선)으로 라우터-PC 연결
## 단축키

ctrl + c 뒤로가기
ctrl + a
Tab: 명령어 자동 완
enable: 특권모드
disable: 특권모드 해제
config t: config 모드

## 11주차
인터페이스 모드: int fa 0/0 (fa=fastethernet)
(config-if) = 인터페이스 모드라는 것을 알려줌

ip주소 설정 : ip address _ _ip주소 입력_ _ _서브넷마스크 입력_ 
No shut = 셧다운 해제

%LINK(2계층)-5-CHANGED: Interface FastEthernet0/0, changed state to **==up==**
%LINEPROTO(3계층_ip주소)-5-UPDOWN: Line protocol on Interface FastEthernet0/0, changed state to **==up==**

ip주소 확인 법: 제어판\\네트워크 및 인터넷\\네트워크 및 공유 센터\\연결\\속성\\TCP/IP확인

화면에 Fa 0/0 보이게 하는 법
옵션 - 프레퍼런스 - Always Show Port Labels

DCE는 서버라서 속도 설정이 필요함.
속도 설정법 (interface serial mode 안에서 가능): <font color="#ffc000">clock rate 0000</font>(속도 단)

DCE(서버) - DTE(터미널=클라이언트)

## 12주차
라우터 시리얼 넣는법: 라우터 클릭 -> 전원 OFF -> WIC-2T 넣기 -> 전원 ON

시리얼 연결은 라우터와 라우터를 연결하기 위함

시리얼 넘버 확인법: (특권모드) sh controllers s 0/0

sh ip int brief
int s0/0




## 13주차

스태틱 라우팅

- 관리자가 직접입력하는것, 외부에서 우리 네트워크에 들어오지못하게 하고 우리는 외부와 통신이 가능하게 하는것

- 스태틱 라우팅은 특정 네트워크로 가기 위해 경로를 자신이 직접 써줘야함.

- 특정 네트워크로 가기 위해 exit i/f, next hop address를 지정해야줘야함.
	- exit i/f: 내 라우터의 나가려는 방향의 인터페이스
	- next hop address: 나간 방향과 연결되어있는 다른 라우터의 ip 주소

Default 라우팅 구성

- 모든 네트워크와 연결
- ip address 0.0.0.0 0.0.0.0 exit i/f or next hop adress

copy run strart: run에 있는 정보를 start에 복사
show ip route: 라우팅 테이블 보는 명령어
- c = 연결되어있는 ip 주

## 14주차
rip=동적 라우팅

rip보다 static이 신뢰성이 더 높기 때문에 static만 동작함

rip 설정
```
RouterA(config)#router rip

RouterA(config-router)#network 168.90.10.0

RouterA(config-router)#network 168.90.40.0

```
스태틱 라우팅 취소방법
```
RouterA(config)#no ip route 168.90.30.0 255.255.255.0 Serial0/0

RouterA(config)#no ip route 168.90.50.0 255.255.255.0 Serial0/0

RouterA(config)#no ip route 168.90.20.0 255.255.255.0 Serial0/0
```

```
R 168.90.10.0 [120/2(홉 수)] via 168.90.50.1(Next hop address), 00:00:15, Serial0/0(라우터의 시리얼 No.)

R 168.90.20.0 [120/1] via 168.90.50.1, 00:00:15, Serial0/0

C 168.90.30.0 is directly connected, FastEthernet0/0

R 168.90.40.0 [120/1] via 168.90.50.1, 00:00:15, Serial0/0

C 168.90.50.0 is directly connected, Serial0/0
```