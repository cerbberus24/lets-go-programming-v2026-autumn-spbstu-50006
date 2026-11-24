package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type rateFloat float64

func (r *rateFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	normalized := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")
	val, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", raw, err)
	}

	*r = rateFloat(val)
	return nil
}

func (r rateFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(r))
}

type Currency struct {
	NumCode  int       `xml:"NumCode" json:"num_code"`
	CharCode string    `xml:"CharCode" json:"char_code"`
	Value    rateFloat `xml:"Value" json:"value"`
}

type valCurs struct {
	Valutes []Currency `xml:"Valute"`
}

func Load(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var doc valCurs
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return doc.Valutes, nil
}

func SortByValueDesc(items []Currency) {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})
}
