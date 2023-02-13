package e2e_test

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal/membership"
	"net/http"
	"testing"
)

func TestTossRecreate(t *testing.T) {
	echoServer := echo.New()
	membership.NewDefaultServer().Routes(echoServer)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(echoServer),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	// todo 과제 3 - 1
	t.Run("멤버십의 주인만 멤버십을 조회할 수 있다", func(t *testing.T) {
		// GIVEN : 멤버십을 생성한다.
		createMembership := e.POST("/v1/memberships").WithJSON(membership.CreateRequest{
			UserName:       "ohjuhyeon",
			MembershipType: "toss",
		}).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		// WHEN :멤버십을 생성한 사용자가 로그인한다.
		e.POST("/v1/memberships/logins").WithJSON(membership.LoginRequest{
			ID:       createMembership.Value("id").String().Raw(),
			UserName: "ohjuhyeon",
		})

		// THEN :사용자의 멤버십 단건 조회가 가능하다.
		e.GET(fmt.Sprintf("/v1/memberships/%s", createMembership.Value("id").Raw())).
			Expect().
			Status(http.StatusOK)

	})

	// todo 과제 3 - 2
	t.Run("Admin 사용자는 멤버십 전체 조회를 할 수 있다.", func(t *testing.T) {

	})

	t.Run("멤버십을 발급 받고 조회하면 발급한 정보가 나온다.", func(t *testing.T) {
		// GIVEN

		// WHEN : 멤버십을 발급 받는다.
		createMembership := e.POST("/v1/memberships").WithJSON(membership.CreateRequest{
			UserName:       "ohjuhyeon",
			MembershipType: "toss",
		}).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		// THEN
		e.GET(fmt.Sprintf("/v1/memberships/%s", createMembership.Value("id").Raw())).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("토스 멤버십을 신청한 후 삭제했다면, 다시 신청할 수 없다.", func(t *testing.T) {
		// given: 토스 멤버십을 신청한다.
		membershipCreateRequest := e.POST("/v1/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// when: 토스 멤버십을 삭제한다.
		e.DELETE(fmt.Sprintf("/v1/memberships/%s", membershipCreateRequest.Value("ID").Raw())).
			Expect().
			Status(http.StatusOK)

		// then: 토스 멤버십을 다시 신청할 수 없다. 멤버십의 상태가 "탈퇴한 회원"이다.
		e.POST("/v1/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusBadRequest).
			JSON().Object().
			Value("message").Equal("재가입할 수 없습니다.")
	})
}
