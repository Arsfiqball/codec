package gormrepo

import (
	"github.com/Arsfiqball/codec/internal/value/domain"

	"github.com/Arsfiqball/talkback"
)

var domainTranslations = talkback.SqlTranslations{
	domain.FieldID:       talkback.SqlFieldTranslation{Column: "id", TypeConverter: talkback.SqlConvertString},
	domain.FieldName:     talkback.SqlFieldTranslation{Column: "name", TypeConverter: talkback.SqlConvertString},
	domain.FieldEmail:    talkback.SqlFieldTranslation{Column: "email", TypeConverter: talkback.SqlConvertString},
	domain.FieldPassword: talkback.SqlFieldTranslation{Column: "password", TypeConverter: talkback.SqlConvertString},
}

var domainPreloadable = talkback.SqlPreloadable{}
