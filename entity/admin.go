package entity

type Admin struct {
	// 別のパッケージで呼び出す場合は変数名を大文字で始める必要がある
	Id                  int
	LineMessagingId     int
	LineMessagingSecret string
	LineMessagingToken  string
	LineLoginId         int
	LineLoginSecret     string
}
