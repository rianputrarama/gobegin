package main

import "testing"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 29/04/2023;
 * Time : 10:57;
**/
func TestHello(t *testing.T) {
	got := Hello("rose")
	want := "Hello, rose"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
