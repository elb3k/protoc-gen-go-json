package example

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestSimple(t *testing.T) {
	// Marshal
	in := &Hello{
		Status: Status_STATUS_PASSED,
		Param: &Hello_Num{
			Num: 1,
		},
	}
	bytes, err := json.Marshal(in)
	require.Nil(t, err)

	// Unmarshal
	var out Hello
	err = json.Unmarshal(bytes, &out)
	require.Nil(t, err)
	require.True(t, proto.Equal(in, &out))
}

func TestComplex(t *testing.T) {
	// Complex type that includes Hello
	type Complex struct {
		Param string
		Hello *Hello
	}

	// Marshal
	in := &Complex{
		Param: "param",
		Hello: &Hello{
			Status: Status_STATUS_FAILED,
			Param: &Hello_Text{
				Text: "text",
			},
		},
	}
	bytes, err := json.Marshal(in)
	require.Nil(t, err)

	// Unmarshal
	var out Complex
	err = json.Unmarshal(bytes, &out)
	require.Nil(t, err)
	// reflect.DeepEqual did not work
	require.True(t, in.Param == out.Param)
	require.True(t, proto.Equal(in.Hello, out.Hello))

}
