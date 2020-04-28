package csv

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"strings"

	iconv "github.com/djimenez/iconv-go"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// ConvertCSV csvのバイト列を返却
// つまり、二次元配列をcsvにして返したいときに使う
// ２つ目の引数はShiftJISに変換して返すかどうかのbool型
func ConvertCSV(rows [][]string, convertSJIS bool) ([]byte, error) {
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	for _, row := range rows {
		w.Write(row)
	}
	w.Flush()

	if !convertSJIS {
		return b.Bytes(), nil
	}
	return toSJIS(b.Bytes())
}

func toSJIS(in []byte) ([]byte, error) {
	out := make([]byte, len(in))
	converter, err := iconv.NewConverter("utf-8", "sjis")
	if err != nil {
		return nil, err
	}
	if _, _, err := converter.Convert(in, out); err != nil {
		return nil, err
	}

	// 最後の 空白の行を削除する
	out = removeZeroByte(out, 0)

	return out, nil
}

func removeZeroByte(numbers []byte, search byte) []byte {
	result := []byte{}
	for _, num := range numbers {
		if num != search {
			result = append(result, num)
		}
	}
	return result
}

// ConvertCsvAsShiftJis csvのバイト列を返却
func ConvertCsvAsShiftJis(rows [][]string, convertSJIS bool) ([]byte, error) {
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	for _, row := range rows {
		for k, v := range row {
			row[k], _ = utf8ToSjis(v)
		}
		w.Write(row)
	}
	w.Flush()

	return b.Bytes(), nil
}

// UTF-8 から ShiftJIS
func utf8ToSjis(str string) (string, error) {
	iostr := strings.NewReader(str)
	rio := transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
	ret, err := ioutil.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), err
}

