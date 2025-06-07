package hat_test

import (
	"testing"

	"github.com/Arsfiqball/codec/widget/hat"
)

func TestProtect(t *testing.T) {
	t.Run("should protect struct", func(t *testing.T) {
		type subSubDocumentT struct {
			ProtectedForAdmin string `protect:"admin"`
			ProtectedForSelf  string `protect:"admin,self"`
		}

		type subDocumentT struct {
			ProtectedForAdmin string          `protect:"admin"`
			ProtectedForSelf  string          `protect:"admin,self"`
			ProtectedDeep     subSubDocumentT `protect:"admin,self"`
		}

		type documentT struct {
			Public            string
			ProtectedForAdmin string `protect:"admin"`
			ProtectedForSelf  string `protect:"admin,self"`
			ProtectedDeep     subDocumentT
		}

		doc := documentT{
			Public:            "abc",
			ProtectedForAdmin: "def",
			ProtectedForSelf:  "ghi",
			ProtectedDeep: subDocumentT{
				ProtectedForAdmin: "jkl",
				ProtectedForSelf:  "mno",
				ProtectedDeep: subSubDocumentT{
					ProtectedForAdmin: "pqr",
					ProtectedForSelf:  "stu",
				},
			},
		}

		if err := hat.Protect(&doc, "self"); err != nil {
			t.Error("should not return error")
		}

		// b, err := hat.MarshalJSON(doc)
		// if err != nil {
		// 	t.Error("should not return error", err)
		// }

		// fmt.Println(string(b))
	})

	t.Run("should protect slice of struct", func(t *testing.T) {
		type subSubDocumentT struct {
			ProtectedForAdmin string `protect:"admin"`
			ProtectedForSelf  string `protect:"admin,self"`
		}

		type subDocumentT struct {
			ProtectedForAdmin string          `protect:"admin"`
			ProtectedForSelf  string          `protect:"admin,self"`
			ProtectedDeep     subSubDocumentT `protect:"admin,self"`
		}

		type documentT struct {
			Public            string
			ProtectedForAdmin string `protect:"admin"`
			ProtectedForSelf  string `protect:"admin,self"`
			ProtectedDeep     []subDocumentT
		}

		docs := []documentT{{
			Public:            "abc",
			ProtectedForAdmin: "def",
			ProtectedForSelf:  "ghi",
			ProtectedDeep: []subDocumentT{
				{
					ProtectedForAdmin: "jkl",
					ProtectedForSelf:  "mno",
					ProtectedDeep: subSubDocumentT{
						ProtectedForAdmin: "pqr",
						ProtectedForSelf:  "stu",
					},
				},
				{
					ProtectedForAdmin: "vwx",
					ProtectedForSelf:  "yza",
					ProtectedDeep: subSubDocumentT{
						ProtectedForAdmin: "bcd",
						ProtectedForSelf:  "efg",
					},
				},
			},
		}}

		if err := hat.Protect(&docs, "self"); err != nil {
			t.Error("should not return error")
		}

		// b, err := hat.MarshalJSON(docs)
		// if err != nil {
		// 	t.Error("should not return error", err)
		// }

		// fmt.Println(string(b))
	})

	t.Run("should work with hat.Unit", func(t *testing.T) {
		type subSubDocumentT struct {
			ProtectedForAdmin hat.Unit[string] `protect:"admin"`
			ProtectedForSelf  hat.Unit[string] `protect:"admin,self"`
		}

		type subDocumentT struct {
			ProtectedForAdmin hat.Unit[string]          `protect:"admin"`
			ProtectedForSelf  hat.Unit[string]          `protect:"admin,self"`
			ProtectedDeep     hat.Unit[subSubDocumentT] `protect:"admin,self"`
		}

		type documentT struct {
			Public            string
			ProtectedForAdmin hat.Unit[string] `protect:"admin"`
			ProtectedForSelf  hat.Unit[string] `protect:"admin,self"`
			ProtectedDeep     hat.Unit[subDocumentT]
		}

		doc := documentT{
			Public:            "abc",
			ProtectedForAdmin: hat.Value("def"),
			ProtectedForSelf:  hat.Value("ghi"),
			ProtectedDeep: hat.Value(subDocumentT{
				ProtectedForAdmin: hat.Value("jkl"),
				ProtectedForSelf:  hat.Value("mno"),
				ProtectedDeep: hat.Value(subSubDocumentT{
					ProtectedForAdmin: hat.Value("pqr"),
					ProtectedForSelf:  hat.Value("stu"),
				}),
			}),
		}

		if err := hat.Protect(&doc, "self"); err != nil {
			t.Error("should not return error", err)
		}

		if doc.ProtectedForAdmin.IsPresent() {
			t.Error("doc.ProtectedForAdmin should not present")
		}

		if doc.ProtectedDeep.Get().ProtectedForAdmin.IsPresent() {
			t.Error("doc.ProtectedDeep.Get().ProtectedForAdmin should not present")
		}

		if doc.ProtectedDeep.Get().ProtectedDeep.Get().ProtectedForAdmin.IsPresent() {
			t.Error("doc.ProtectedDeep.Get().ProtectedDeep.Get().ProtectedForAdmin should not present")
		}

		// TODO: Bug hat.MarshalJSON shows {} for non present value instead of omitting it
	})

	t.Run("should work with pointer", func(t *testing.T) {
		type subSubDocumentT struct {
			ProtectedForAdmin string `protect:"admin"`
			ProtectedForSelf  string `protect:"admin,self"`
		}

		type subDocumentT struct {
			ProtectedForAdmin string           `protect:"admin"`
			ProtectedForSelf  *string          `protect:"admin,self"`
			ProtectedDeep     *subSubDocumentT `protect:"admin,self"`
		}

		type documentT struct {
			Public            string
			ProtectedForAdmin string `protect:"admin"`
			ProtectedForSelf  string `protect:"admin,self"`
			ProtectedDeep     *subDocumentT
		}

		v := "mno"

		doc := documentT{
			Public:            "abc",
			ProtectedForAdmin: "def",
			ProtectedForSelf:  "ghi",
			ProtectedDeep: &subDocumentT{
				ProtectedForAdmin: "jkl",
				ProtectedForSelf:  &v,
				ProtectedDeep: &subSubDocumentT{
					ProtectedForAdmin: "pqr",
					ProtectedForSelf:  "stu",
				},
			},
		}

		if err := hat.Protect(&doc, "self"); err != nil {
			t.Error("should not return error")
		}

		// b, err := hat.MarshalJSON(doc)
		// if err != nil {
		// 	t.Error("should not return error")
		// }

		// fmt.Println(string(b))
	})
}
