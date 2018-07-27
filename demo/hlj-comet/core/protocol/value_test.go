package protocol

import(
	"testing"
)
func TestValues_GetSet(t *testing.T) {
	values := Values{}
	values.Set("foo", "bar")

	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar but got, %s", v)
	}
}

func TestValues_MarshalJSON(t *testing.T) {
	values := Values{}
	values.Set("foo", "bar")
	values.Set("bar", "foo")

	bytes, err := json.Marshal(values)
	if err != nil {
		t.Errorf("marshal error: %s", err.Error())
	}

	if string(bytes) != `{"foo":"bar","bar":"foo"}` && string(bytes) != `{"bar":"foo","foo":"bar"}` {
		t.Errorf("not as expected, got: %s", bytes)
	}
}

func TestValues_UnmarshalJSON(t *testing.T) {
	bytes := `{"foo":"bar","bar":"foo"}`
	values := Values{}

	if err := json.Unmarshal([]byte(bytes), &values); err != nil {
		t.Errorf("unmarshal error: %s", err.Error())
	}
	if v := values.Get("foo"); v != "bar" {
		t.Errorf("expected bar, but got %s", v)
	}
}