package mojikiban

type Kanjishisaku struct {
	Jyouyoukanji   bool `json:"常用漢字"`
	Jinmeiyoukanji bool `json:"人名用漢字"`
}

type JISX0213 struct {
	Housetsukubun  string `json:"包摂区分"`
	Gamenkutenichi string `json:"面区点位置"`
	Suijun         string `json:"水準"`
}

type UCS struct {
	UCS      string `json:"対応するUCS"`
	Category string `json:"対応カテゴリー"`
}

type Font struct {
	Version string `json:"フォントバージョン"`
	UCS     string `json:"実装したUCS"`
}

type MJMojizukei struct {
	Uri     string `json:"uri"`
	Version string `json:"MJ文字図形バージョン"`
}

type Busyunaikakusuu struct {
	Naikakusuu int `json:"内画数"`
	Bushu      int `json:"部首"`
}

type Yomi struct {
	Onyomi  []string `json:"音読み"`
	Kunyomi []string `json:"訓読み"`
}

type MJCharInfo struct {
	MJMojizukeimei          string            `json:"MJ文字図形名"`
	Kosekitouitsumojinumber string            `json:"戸籍統一文字番号"`
	Jukinettouitsumojicode  string            `json:"住基ネット統一文字コード"`
	Nyukanseijicode         string            `json:"入管正字コード"`
	Nyukangaijicode         string            `json:"入管外字コード"`
	Kanjishisaku            Kanjishisaku      `json:"漢字施策"`
	JISX0213                JISX0213          `json:"JISX0213"`
	UCS                     UCS               `json:"UCS"`
	Font                    Font              `json:"IPAmj明朝フォント実装"`
	MJMojizukei             MJMojizukei       `json:"MJ文字図形"`
	Toukitouitsumojinumber  string            `json:"登記統一文字番号"`
	Busyunaikakusuu         []Busyunaikakusuu `json:"部首内画数"`
	Soukakusuu              int               `json:"総画数"`
	Yomi                    Yomi              `json:"読み"`
	Daikanwa                string            `json:"大漢和"`
	Nihongokanjijiten       int               `json:"日本語漢字辞典"`
	Shindaijiten            int               `json:"新大字典"`
	Daijigen                int               `json:"大字源"`
	Daikangorin             int               `json:"大漢語林"`
}
