package form

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postData := url.Values{}
	f := New(postData)

	isValid := f.Valid()
	if !isValid {
		t.Error("expected valid; got invalid")
	}
}

func TestForm_Required(t *testing.T) {
	postData := url.Values{}
	f := New(postData)

	f.Required("a", "b", "c")
	if f.Valid() {
		t.Error("expected invalid; got valid")
	}

	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	f = New(postData)
	f.Required("a", "b", "c")
	if !f.Valid() {
		t.Error("expected valid; got invalid")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	f := New(postData)

	got := f.Has("a")
	if got {
		t.Error("expected false; got true")
	}

	postData.Add("a", "a")

	got = f.Has("a")
	if !got {
		t.Error("expected true; got false")
	}
}

func TestForm_MinLength(t *testing.T) {
	postData := url.Values{}
	f := New(postData)

	f.MinLength("a", 10)
	if f.Valid() {
		t.Error("expected false; got true")
	}

	postData.Add("hello", "hello")
	f = New(postData)

	f.MinLength("hello", 5)
	if !f.Valid() {
		t.Error("expected true; got false")
	}

	f.MinLength("hello", 10)
	if f.Valid() {
		t.Error("expected false; got true")
	}
}

func TestForm_MaxLength(t *testing.T) {
	postData := url.Values{}
	f := New(postData)

	f.MaxLength("a", 10)
	if !f.Valid() {
		t.Error("expected true; got false")
	}

	postData.Add("hello", "hello world")
	f = New(postData)

	f.MaxLength("hello", 20)
	if !f.Valid() {
		t.Error("expected true; got false")
	}

	f.MaxLength("hello", 10)
	if f.Valid() {
		t.Error("expected false; got true")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postData := url.Values{}
	f := New(postData)

	f.IsEmail("a")
	if f.Valid() {
		t.Error("expected false; got true")
	}

	postData.Add("email-valid", "abc@example.com")
	postData.Add("email-invalid", "abc")

	f = New(postData)

	f.IsEmail("email-valid")
	if !f.Valid() {
		t.Error("expected true; got false")
	}

	f.IsEmail("email-invalid")
	if f.Valid() {
		t.Error("expected false; got true")
	}
}
