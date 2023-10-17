package main

import "encoding/xml"

type EpubContainerXml struct {
	XMLName   xml.Name `xml:"container"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Version   string   `xml:"version,attr"`
	Rootfiles struct {
		Text     string   `xml:",chardata"`
		Rootfile RootFile `xml:"rootfile"`
	} `xml:"rootfiles"`
}

type RootFile struct {
	Text      string `xml:",chardata"`
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

type EpubPackXml struct {
	XMLName          xml.Name `xml:"package" json:"-"`
	Text             string   `xml:",chardata" json:"-"`
	Xmlns            string   `xml:"xmlns,attr" json:"-"`
	UniqueIdentifier string   `xml:"unique-identifier,attr" json:"uniqueIdentifier"`
	Version          string   `xml:"version,attr" json:"version"`
	Metadata         struct {
		Text    string `xml:",chardata" json:"-"`
		Dc      string `xml:"dc,attr" json:"dc"`
		Dcterms string `xml:"dcterms,attr" json:"dcterms"`
		Opf     string `xml:"opf,attr" json:"opf"`
		Xsi     string `xml:"xsi,attr" json:"xsi"`
		Date    string `xml:"date" json:"date"`
		Meta    []struct {
			Text    string `xml:",chardata" json:"-"`
			Name    string `xml:"name,attr" json:"name"`
			Content string `xml:"content,attr" json:"content"`
		} `xml:"meta" json:"meta"`
		Title      string `xml:"title" json:"title"`
		Language   string `xml:"language" json:"language"`
		Identifier []struct {
			Text   string `xml:",chardata" json:"-"`
			Scheme string `xml:"scheme,attr" json:"scheme"`
			ID     string `xml:"id,attr" json:"id"`
		} `xml:"identifier" json:"identifier"`
		Creator struct {
			Text   string `xml:",chardata" json:"-"`
			FileAs string `xml:"file-as,attr" json:"fileAs"`
			Role   string `xml:"role,attr" json:"role"`
		} `xml:"creator" json:"creator"`
		Contributor struct {
			Text string `xml:",chardata" json:"-"`
			Role string `xml:"role,attr" json:"role"`
		} `xml:"contributor" json:"contributor"`
	} `xml:"metadata" json:"metadata"`
	Manifest struct {
		Text string `xml:",chardata" json:"-"`
		Item []struct {
			Text      string `xml:",chardata" json:"-"`
			Href      string `xml:"href,attr" json:"href"`
			ID        string `xml:"id,attr" json:"id"`
			MediaType string `xml:"media-type,attr" json:"mediaType"`
		} `xml:"item" json:"item"`
	} `xml:"manifest" json:"manifest"`
	Spine struct {
		Text    string `xml:",chardata" json:"-"`
		Toc     string `xml:"toc,attr" json:"toc"`
		Itemref []struct {
			Text  string `xml:",chardata" json:"-"`
			Idref string `xml:"idref,attr" json:"idref"`
		} `xml:"itemref" json:"itemref"`
	} `xml:"spine" json:"spine"`
}

type BookIndexData struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	FileList []string `json:"fileList"`
}
