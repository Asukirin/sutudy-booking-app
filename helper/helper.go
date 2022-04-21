package helper

import "strings"

// 入力値のバリデーションチェック
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// 姓名がそれぞれ２文字以下か
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// メールアドレスに@が含まれているか
	isValidEmail := strings.Contains(email, "@")
	// チケット数の設定が適切か
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// TODO:予約済みユーザーのチェックを実装
	return isValidName, isValidEmail, isValidTicketNumber
}
