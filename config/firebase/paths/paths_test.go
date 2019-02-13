package paths

import (
	"testing"
)

func TestPathConstruction(t *testing.T) {
	ex := "/documents/MonetaryRequests/XUtvJm2jMae6CVwa33MVbYN2iZH2/2019-02/g3aRogfhNIKwcwTxkGwF"
	ex2 := "/documents/GroupRequests/XUtvJm2jMae6CVwa33MVbYN2iZH2/2019-02/g3aRogfhNIKwcwTxkGwF"
	id := "AGd1wTq8VTRdxUlokeVcwOjx7Ce2"

	str := ExtractAndReplaceMethodIdAndDatePath(
		id,
		ex,
	)

	if str != "MonetaryRequests/AGd1wTq8VTRdxUlokeVcwOjx7Ce2/2019-02" {
		t.Error(str)
	}

	str = ExtractMethodIdAndDatePath(ex)

	if str != "MonetaryRequests/XUtvJm2jMae6CVwa33MVbYN2iZH2/2019-02" {
		t.Error(str)
	}

	snf, str := ExtractAndReplaceMethodIdAndDatePathWithSnowflake(
		id,
		ex,
	)

	if str != "MonetaryRequests/AGd1wTq8VTRdxUlokeVcwOjx7Ce2/2019-02" || snf != "g3aRogfhNIKwcwTxkGwF" {
		t.Error(str, snf)
	}

	snf, str = ExtractMethodIdAndDatePathWithSnowflake(ex)

	if str != "MonetaryRequests/XUtvJm2jMae6CVwa33MVbYN2iZH2/2019-02" || snf != "g3aRogfhNIKwcwTxkGwF" {
		t.Error(str)
	}

	str = TransformGroupIntoMonetary(ex2)

	if str != ex {
		t.Error(str)
	}

	str = ExtractAndReplaceMethodIdAndDatePath(id, ex2)

	if str != "GroupRequests/AGd1wTq8VTRdxUlokeVcwOjx7Ce2/2019-02" {
		t.Error(str)
	}

	str = TransformGroupIntoMonetary(str)

	if str != "MonetaryRequests/AGd1wTq8VTRdxUlokeVcwOjx7Ce2/2019-02" {
		t.Error(str)
	}
}
