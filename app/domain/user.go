package domain

// Userエンティティ
type User struct {
	ID       int
	Name     string
	Email    string
	Password string  // パスワードはハッシュ化された状態で保存される
}

// UserRepositoryインターフェース
// ユーザーに関連するデータ操作を定義する
type UserRepository interface {
	// ユーザーを保存する
	Save(user User) error
	// ユーザーIDでユーザーを取得する
	GetByID(id int) (*User, error)
	// メールアドレスでユーザーを取得する（認証でよく使う）
	GetByEmail(email string) (*User, error)
	// ユーザーを削除する
	Delete(id int) error
}
