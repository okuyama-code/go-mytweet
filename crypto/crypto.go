package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncrypt はパスワードをハッシュ化します
func PasswordEncrypt(password string) (string, error) {
	// bcryptパッケージを使用してパスワードをハッシュ化します
	// bcrypt.DefaultCostはハッシュの強度を指定します
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword はハッシュと非ハッシュパスワードを比較します
func CompareHashAndPassword(hash, password string) error {
	// bcryptパッケージを使用してハッシュと非ハッシュパスワードを比較します
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// PasswordEncrypt関数は、与えられたパスワードをハッシュ化します。bcrypt.GenerateFromPassword関数は、[]byte(password)をハッシュ化して強固なハッシュを生成します。bcrypt.DefaultCostはハッシュの計算コストを指定し、この値が高いほどハッシュの計算に時間がかかりますが、より強力なハッシュが生成されます。関数はハッシュ化されたパスワードとエラーを返します。

// CompareHashAndPassword関数は、ハッシュと非ハッシュパスワードを比較します。bcrypt.CompareHashAndPassword関数は、ハッシュと非ハッシュパスワードをバイト配列として受け取り、比較を行います。エラーがなければnilを返し、エラーがあればエラーを返します。

// これらの関数を使用することで、パスワードのハッシュ化とハッシュと非ハッシュパスワードの比較が容易に行えます。




