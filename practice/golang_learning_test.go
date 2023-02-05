package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {

		str := "Ann,Jenny,Tom,Zico"
		// TODO str을 , 단위로 잘라주세요.
		actual := strings.Split(str, ",") //스플릿으로 ,를 기준으로 문자열을 구분해 슬라이스에 넣어준다.

		expected := []string{"Ann", "Jenny", "Tom", "Zico"}

		//TODO assert 문을 활용해 actual과 expected를 비교해주세요.
		assert.EqualValues(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {

		var numbers []int
		var wg sync.WaitGroup // 고루틴 실행을 제어하기 위한 WaitGroup 설정
		var m sync.Mutex      // 임계구역 제어를 위한 뮤텍스 설정

		size := 100
		wg.Add(size) // 고루틴 총 100개 대기 예정

		for i := 0; i < 100; i++ {
			go func(i int) {
				m.Lock() // 고루틴의 비동기 처리를 위해 수정이 발생하는 동안 다른 자원 접근 막아둠
				// TODO numbers에 i 값을 추가해보세요.
				numbers = append(numbers, i) // 이벤트 수행
				m.Unlock()                   // 이벤트 수행 후 임계구역 개방 -> 다른 고루틴이 임계구역으로 진입 가능
				wg.Done()                    // 고루틴 하나로 하나의 이벤트 수행 후 고루틴 소비
			}(i)
		}
		wg.Wait() // 모든 고루틴이 다 Done 될 때 까지 대기..

		var expected []int // actual : [0 1 2 ... 99]
		// TODO expected를 만들어주세요.
		for i := 0; i < 100; i++ {
			expected = append(expected, i)
		}
		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {
		/*
			TODO 주어진 코드를 수정해서 테스트가 통과하도록 해주세요!

			- inputCh에 1, 2, 3 값을 넣는다.
			- inputCh로 부터 값을 받아와, value * 10 을 한 후 outputCh에 값을 넣어준다.
			- outputCh에서 읽어온 값을 비교한다.
		*/

		inputCh := generate()
		outputCh := make(chan int)
		go func() {
			for {
				select {
				case value, ok := <-inputCh: // 인풋 채널에서 값을 꺼낸다.
					if !ok { // 값을 다 꺼냈으면
						close(outputCh) // 아웃풋 채널을 닫아준다.
						return          // 그리고 리턴
					}
					outputCh <- value * 10 // 꺼낸 인풋 값은 value에 담기고 여기에 10을 곱한다음 아웃풋 채널에 넣어준다.
				}
			}
		}()

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		// TODO 3초후에 종료하는 timeout context로 만들어주세요.
		ctx := context.TODO()                       //todo는 background와 같이 초기화 용도로만 사용
		timeout, _ := context.WithTimeout(ctx, add) // 특정 시간이 지나면 종료

		var endTime time.Time
		endTime.Add(add)
		select {
		case <-timeout.Done(): // 타임아웃이 발생하면 채널이 반환된다.
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.TODO()                                       // TODO 3초후에 종료하는 timeout context로 만들어주세요.
		timeout, _ := context.WithDeadline(ctx, startTime.Add(add)) // 마감 시간 설정 지금 시간에 3초 추가

		var endTime time.Time
		select {
		case <-timeout.Done(): // 타임아웃 발생 시 마찬가지로 채널 반환
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		// context에 key, value를 추가해보세요.
		ctx := context.TODO()
		ctx = context.WithValue(ctx, "key", "value")
		// 추가된 key, value를 호출하여 assert로 값을 검증해보세요.
		assert.Equal(t, "value", ctx.Value("key")) // 문자열 value가 ctx에 key라는 키의 벨류와 같은지 체크
		// 추가되지 않은 key에 대한 value를 assert로 검증해보세요.
		assert.Nil(t, ctx.Value("hahahoho")) // hahahoho의 키로 가져오는 벨류가 nil인지 체크
	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
