package errors

import (
	"errors"
)

// New initializeというか単なるwrap
func New(e string) error {
	return errors.New(e)
}

// ErrNotFoundRecord レコードが見つからないエラー
var ErrNotFoundRecord = errors.New("not found record")

// ErrEmptyAffectedRows affected_rowsが0
var ErrEmptyAffectedRows = errors.New("empty affected rows")

// ErrDuplicatedRecord レコードが存在した場合のエラー
var ErrDuplicatedRecord = errors.New("isset record")

// ErrOutOfRange 範囲外更新した場合のエラー
var ErrOutOfRange = errors.New("out of error")

// ErrArgument 引数エラー
var ErrArgument = errors.New("argument error")

// ErrInvalidToken トークンエラー
var ErrInvalidToken = errors.New("invalid token")

// ErrExpired 有効期限切れ
var ErrExpired = errors.New("expired")

// ErrUserIsDead ユーザが無効
var ErrUserIsDead = errors.New("user is dead")

// ErrUserIsBanned withdraw_typeが0(強制退会)
var ErrUserIsBanned = errors.New("user is banned")

// ErrUserAlreadyWithdraw withdraw_typeが1(通常退会)
var ErrUserAlreadyWithdraw = errors.New("user already withdraw")

// ErrInvalidPassword パスワード間違い
var ErrInvalidPassword = errors.New("invalid password")

// ErrUserAlreadySignup ユーザが登録済み
var ErrUserAlreadySignup = errors.New("user is already signup")

// ErrUserNotSignup ユーザが未登録
var ErrUserNotSignup = errors.New("user is not signup")

// ErrInvitedAlready すでに招待を受けている
var ErrInvitedAlready = errors.New("alread invited")

// ErrCouponEndedAtOlderThanStartedAt 終了日の方が開始日より古い
var ErrCouponEndedAtOlderThanStartedAt = errors.New("startedAt must be older than endedAt")

// ErrOptimisticLock 楽観ロックエラー
var ErrOptimisticLock = errors.New("optimistic lock error, version mismatch")

// ErrModelNotEnoughAppend モデルがすべての要素を保持した状態ではない
var ErrModelNotEnoughAppend = errors.New("model has not append models")
