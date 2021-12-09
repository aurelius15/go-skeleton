package reflection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringFieldByName(t *testing.T) {
	type testStruct struct {
		PublicPort  string
		privatePort string
		SessionID   int
	}

	type input struct {
		obj       interface{}
		filedName string
	}

	type output struct {
		value string
		err   error
	}

	type test struct {
		input    input
		expected output
	}

	testObj := testStruct{
		PublicPort:  "8080",
		privatePort: "80",
		SessionID:   0,
	}

	tests := []test{
		{
			input: input{
				obj:       testObj,
				filedName: "PublicPort",
			},
			expected: output{
				value: "8080",
				err:   nil,
			},
		},
		{
			input: input{
				obj:       testObj,
				filedName: "privatePort",
			},
			expected: output{
				value: "80",
				err:   nil,
			},
		},
		{
			input: input{
				obj:       testObj,
				filedName: "SessionID",
			},
			expected: output{
				value: "",
				err:   ErrFiledValueIsNotString,
			},
		},
		{
			input: input{
				obj:       testObj,
				filedName: "fieldThatDoesNotExist",
			},
			expected: output{
				value: "",
				err:   ErrFiledValueIsNotString,
			},
		},
	}

	for _, testCase := range tests {
		value, err := StringFieldByName(testCase.input.obj, testCase.input.filedName)
		assert.Equal(t, value, testCase.expected.value)
		assert.Equal(t, err, testCase.expected.err)
	}
}

func TestFieldByTag(t *testing.T) {
	type input struct {
		obj      interface{}
		tagName  string
		tagValue string
	}

	type output struct {
		value reflect.Value
		err   error
	}

	type test struct {
		input    input
		expected output
		msg      string
	}

	objTestCase2 := struct {
		UserName string `json:"user_name"`
	}{UserName: "test"}

	objTestCase3 := struct {
		UserName string `arg:"-c,separate"`
	}{UserName: "test"}

	objTestCase4 := struct {
		UserName string
	}{UserName: "test"}

	tests := []test{
		{
			input: input{
				obj:      nil,
				tagName:  "",
				tagValue: "",
			},
			expected: output{
				value: reflect.Value{},
				err:   ErrTypeNotStructure,
			},
			msg: "Nil instead of structure",
		},
		{
			input: input{
				obj:      objTestCase2,
				tagName:  "json",
				tagValue: "user_name",
			},
			expected: output{
				value: reflect.ValueOf(objTestCase2).FieldByName("UserName"),
				err:   nil,
			},
			msg: "Search by json tag",
		},
		{
			input: input{
				obj:      objTestCase3,
				tagName:  "arg",
				tagValue: "separate",
			},
			expected: output{
				value: reflect.ValueOf(objTestCase3).FieldByName("UserName"),
				err:   nil,
			},
			msg: "Tag value has comma",
		},
		{
			input: input{
				obj:      objTestCase4,
				tagName:  "arg",
				tagValue: "separate",
			},
			expected: output{
				value: reflect.Value{},
				err:   ErrFieldWithRequestedTagNotExist,
			},
			msg: "Tag value has comma",
		},
	}

	for _, testCase := range tests {
		value, err := FieldByTag(testCase.input.obj, testCase.input.tagName, testCase.input.tagValue)

		if !value.IsValid() && !testCase.expected.value.IsValid() {
			assert.Equal(t, value, testCase.expected.value, testCase.msg)
		} else {
			assert.Equal(t, value.Interface(), testCase.expected.value.Interface(), testCase.msg)
		}

		assert.Equal(t, err, testCase.expected.err, testCase.msg)
	}
}
