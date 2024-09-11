package api

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetPriceData は指定されたティッカーシンボルの株価データを取得する
func GetPriceData(ticker string) (string, error) {
	// PythonでyFinanceを呼び出して価格データを取得
	cmd := exec.Command("python3", "-c", fmt.Sprintf(`
import yfinance as yf
ticker = yf.Ticker('%s')
print(ticker.history(period='1d')['Close'][-1])`, ticker))

	// コマンドの出力をキャプチャ
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// 結果を文字列として返す（改行を削除）
	return strings.TrimSpace(string(output)), nil
}
