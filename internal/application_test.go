package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))
		app.Create(CreateRequest{"jenny", "naver"})

		// WHEN
		request, err := app.Create(CreateRequest{"jenny", "naver"})

		// THEN
		assert.NotNil(t, err)
		assert.Nil(t, request)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))

		// WHEN
		request, err := app.Create(CreateRequest{"", "naver"})

		// THEN
		assert.NotNil(t, err)
		assert.Nil(t, request)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))

		// WHEN
		request, err := app.Create(CreateRequest{"ohjuhyeon", ""})

		// THEN
		assert.NotNil(t, err)
		assert.Nil(t, request)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))
		//whiteSlice := []string{"naver", "toss", "payco"}

		// WHEN
		response, err := app.Create(CreateRequest{"ohjuhyeon", "skt"})

		// THEN
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("멤버십 정보를 갱신한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"jenny", "naver"})

		// WHEN
		requestReq := UpdateRequest{createRes.ID, "jenny", "toss"}
		updateRes, err := app.Update(requestReq)

		// THEN
		assert.Nil(t, err) // err 가 NIL 이면 이게 통과되고 에러가 없다.
		assert.Equal(t, requestReq.MembershipType, updateRes.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"jenny", "naver"})
		app.Create(CreateRequest{"ohjuhyeon", "naver"})

		// WHEN
		updateRes, err := app.Update(UpdateRequest{createRes.ID, "ohjuhyeon", createRes.MembershipType})

		// THEN
		assert.Nil(t, updateRes)
		assert.NotNil(t, err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		// GIVEN
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"ohjuhyeon", "naver"})

		// WHEN
		updateRes, err := app.Update(UpdateRequest{"", "ohjuhyeon", createRes.MembershipType})

		// THEN
		assert.Nil(t, updateRes)
		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		// GIVEN

		// WHEN

		// THEN
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		// GIVEN

		// WHEN

		// THEN
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		// GIVEN

		// WHEN

		// THEN
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		// GIVEN

		// WHEN

		// THEN
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		// GIVEN

		// WHEN

		// THEN
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		// GIVEN

		// WHEN

		// THEN
	})
}
