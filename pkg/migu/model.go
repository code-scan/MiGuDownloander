package migu

import (
	"encoding/json"
	"log"
)

type MiGuResult struct {
	Code                   string                 `json:"code"`
	Info                   string                 `json:"info"`
	SongResultData         SongResultData         `json:"songResultData"`
	BestShowResultData     BestShowResultData     `json:"bestShowResultData"`
	BestShowResultToneData BestShowResultToneData `json:"bestShowResultToneData"`
}
type Singers struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Albums struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type ImgItems struct {
	ImgSizeType string `json:"imgSizeType"`
	Img         string `json:"img"`
}
type Tones struct {
	ID          string `json:"id"`
	CopyrightID string `json:"copyrightId"`
	Price       string `json:"price"`
	ExpireDate  string `json:"expireDate"`
}
type RelatedSongs struct {
	ResourceType     string `json:"resourceType"`
	ResourceTypeName string `json:"resourceTypeName"`
	CopyrightID      string `json:"copyrightId"`
	ProductID        string `json:"productId"`
}
type RateFormats struct {
	ResourceType         string `json:"resourceType"`
	FormatType           string `json:"formatType"`
	URL                  string `json:"url,omitempty"`
	Format               string `json:"format"`
	Size                 string `json:"size"`
	FileType             string `json:"fileType,omitempty"`
	Price                string `json:"price"`
	IosURL               string `json:"iosUrl,omitempty"`
	AndroidURL           string `json:"androidUrl,omitempty"`
	AndroidFileType      string `json:"androidFileType,omitempty"`
	IosFileType          string `json:"iosFileType,omitempty"`
	IosSize              string `json:"iosSize,omitempty"`
	AndroidSize          string `json:"androidSize,omitempty"`
	IosFormat            string `json:"iosFormat,omitempty"`
	AndroidFormat        string `json:"androidFormat,omitempty"`
	IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
	AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
}
type NewRateFormats struct {
	ResourceType         string `json:"resourceType"`
	FormatType           string `json:"formatType"`
	URL                  string `json:"url,omitempty"`
	Format               string `json:"format,omitempty"`
	Size                 string `json:"size,omitempty"`
	FileType             string `json:"fileType,omitempty"`
	Price                string `json:"price"`
	IosURL               string `json:"iosUrl,omitempty"`
	AndroidURL           string `json:"androidUrl,omitempty"`
	AndroidFileType      string `json:"androidFileType,omitempty"`
	IosFileType          string `json:"iosFileType,omitempty"`
	IosSize              string `json:"iosSize,omitempty"`
	AndroidSize          string `json:"androidSize,omitempty"`
	IosFormat            string `json:"iosFormat,omitempty"`
	AndroidFormat        string `json:"androidFormat,omitempty"`
	IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
	AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
	AndroidNewFormat     string `json:"androidNewFormat,omitempty"`
	IosBit               int    `json:"iosBit,omitempty"`
	AndroidBit           int    `json:"androidBit,omitempty"`
}
type Z3DCode struct {
	ResourceType    string `json:"resourceType"`
	FormatType      string `json:"formatType"`
	Price           string `json:"price"`
	IosURL          string `json:"iosUrl"`
	AndroidURL      string `json:"androidUrl"`
	AndroidFileType string `json:"androidFileType"`
	IosFileType     string `json:"iosFileType"`
	IosSize         string `json:"iosSize"`
	AndroidSize     string `json:"androidSize"`
	IosFormat       string `json:"iosFormat"`
	AndroidFormat   string `json:"androidFormat"`
	AndroidFileKey  string `json:"androidFileKey"`
	IosFileKey      string `json:"iosFileKey"`
	H5URL           string `json:"h5Url"`
	H5Size          string `json:"h5Size"`
	H5Format        string `json:"h5Format"`
}
type SongResult struct {
	ID               string           `json:"id"`
	ResourceType     string           `json:"resourceType"`
	ContentID        string           `json:"contentId"`
	CopyrightID      string           `json:"copyrightId"`
	Name             string           `json:"name"`
	HighlightStr     []string         `json:"highlightStr"`
	Singers          []Singers        `json:"singers"`
	Albums           []Albums         `json:"albums"`
	Tags             []string         `json:"tags,omitempty"`
	LyricURL         string           `json:"lyricUrl"`
	TrcURL           string           `json:"trcUrl"`
	ImgItems         []ImgItems       `json:"imgItems"`
	TelevisionNames  []string         `json:"televisionNames"`
	Tones            []Tones          `json:"tones,omitempty"`
	RelatedSongs     []RelatedSongs   `json:"relatedSongs"`
	ToneControl      string           `json:"toneControl"`
	RateFormats      []RateFormats    `json:"rateFormats"`
	NewRateFormats   []NewRateFormats `json:"newRateFormats"`
	Z3DCode          Z3DCode          `json:"z3dCode,omitempty"`
	SongType         string           `json:"songType"`
	IsInDAlbum       string           `json:"isInDAlbum"`
	Copyright        string           `json:"copyright"`
	DigitalColumnID  string           `json:"digitalColumnId"`
	Mrcurl           string           `json:"mrcurl"`
	SongDescs        string           `json:"songDescs"`
	SongAliasName    string           `json:"songAliasName"`
	TranslateName    string           `json:"translateName,omitempty"`
	InvalidateDate   string           `json:"invalidateDate"`
	IsInSalesPeriod  string           `json:"isInSalesPeriod"`
	DalbumID         string           `json:"dalbumId"`
	IsInSideDalbum   string           `json:"isInSideDalbum"`
	VipType          string           `json:"vipType"`
	ChargeAuditions  string           `json:"chargeAuditions"`
	ScopeOfcopyright string           `json:"scopeOfcopyright"`
	MovieNames       []string         `json:"movieNames,omitempty"`
	MvList           []interface{}    `json:"mvList,omitempty"`
	MvCopyright      string           `json:"mvCopyright,omitempty"`
}
type SongResultData struct {
	TotalCount  string        `json:"totalCount"`
	Correct     []interface{} `json:"correct"`
	ResultType  string        `json:"resultType"`
	IsFromCache string        `json:"isFromCache"`
	Result      []SongResult  `json:"result"`
	TipStatus   string        `json:"tipStatus"`
}
type SingerPicURL struct {
	ImgSizeType string `json:"imgSizeType"`
	Img         string `json:"img"`
}
type Result struct {
	Mod              string         `json:"mod"`
	ID               string         `json:"id"`
	SingerName       string         `json:"singerName"`
	SingerAliasNames []string       `json:"singerAliasNames"`
	SingerArea       string         `json:"singerArea"`
	AlbumCount       string         `json:"albumCount"`
	SongCount        string         `json:"songCount"`
	MvCount          string         `json:"mvCount"`
	SingerPicURL     []SingerPicURL `json:"singerPicUrl"`
}
type BestShowResultData struct {
	TotalCount string   `json:"totalCount"`
	Result     []Result `json:"result"`
}
type BestShowResultToneData struct {
}

type Music struct {
	ID     string
	Name   string
	Album  string
	Singer string
	URL    string
}

func UnMarshalMiGuResult(result []byte) *MiGuResult {
	var migu MiGuResult
	if err := json.Unmarshal(result, &migu); err != nil {
		log.Println("Unmarshal MiGu Result Error: ", err)
	}
	return &migu

}
