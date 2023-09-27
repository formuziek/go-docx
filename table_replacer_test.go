package docx

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestTableReplacer_Replace(t *testing.T) {
	replaceMap := [][]TablePlaceholder{
		{
			{
				Key:   "Name",
				Value: "John Doe",
			},
			{
				Key:   "Value",
				Value: "42",
			},
		},
		{
			{
				Key:   "Name",
				Value: "Jane Doe",
			},
			{
				Key:   "Value",
				Value: "43",
			},
		},
	}

	doc, err := Open("./test/table_template.docx")
	if err != nil {
		t.Error(err)
		return
	}

	tableReplacer := NewTableReplacer(doc.files[DocumentXml])
	err = tableReplacer.Replace("t1", replaceMap)
	if err != nil {
		t.Error("replacing failed", err)
		return
	}

	doc.SetFile(DocumentXml, tableReplacer.document)

	err = doc.WriteToFile("./test/out.docx")
	if err != nil {
		t.Error("unable to write", err)
		return
	}

	document, err := Open("./test/out.docx")
	if err != nil {
		t.Error("failed to open docx")
		return
	}

	documentXml := document.files[DocumentXml]

	err = xml.Unmarshal(documentXml, new(interface{}))
	if err != nil {
		t.Error("failed to unmarshal xml, replacing failed")
		return
	}

	// cleanup
	_ = os.Remove("./test/out.docx")
}
