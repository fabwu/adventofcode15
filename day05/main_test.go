package main

import (
"testing"
)

func TestIsNiceString1(t *testing.T) {
	expected := true
	actual := isOldNiceString("ugknbfddgicrmopn")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsNiceString2(t *testing.T) {
	expected := true
	actual := isOldNiceString("aaa")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsNiceString3(t *testing.T) {
	expected := false
	actual := isOldNiceString("jchzalrnumimnmhp")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsNiceString4(t *testing.T) {
	expected := false
	actual := isOldNiceString("haegwjzuvuyypxyu")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsNiceString5(t *testing.T) {
	expected := false
	actual := isOldNiceString("dvszwmarrgswjxmb")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCountVowels(t *testing.T) {
	expected := 5
	actual := countVowels("kajsiescvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasNoDoubleLetter(t *testing.T) {
	expected := false
	actual := hasDoubleLetter("asfjcvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasDoubleLetter(t *testing.T) {
	expected := true
	actual := hasDoubleLetter("kaajsiescvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasNoBadString(t *testing.T) {
	expected := false
	actual := hasBadString("kaajsiescvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasBadStringAB(t *testing.T) {
	expected := true
	actual := hasBadString("kaajsabiescvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasBadStringCD(t *testing.T) {
	expected := true
	actual := hasBadString("kaajsbiecdvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasBadStringPQ(t *testing.T) {
	expected := true
	actual := hasBadString("kaajsiescvopqlku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasBadStringXY(t *testing.T) {
	expected := true
	actual := hasBadString("kaajsaixyescvolku")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasPairXY(t *testing.T) {
	expected := true
	actual := hasPair("xyxy")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasPairAA(t *testing.T) {
	expected := true
	actual := hasPair("aabcdefgaa")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasOverlappingPair(t *testing.T) {
	expected := false
	actual := hasPair("aaa")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasSandwich1(t *testing.T) {
	expected := true
	actual := hasSandwich("xyx")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasSandwich2(t *testing.T) {
	expected := true
	actual := hasSandwich("abcdefeghi")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasSandwich3(t *testing.T) {
	expected := true
	actual := hasSandwich("aaa")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasNoSandwich(t *testing.T) {
	expected := false
	actual := hasSandwich("uurcxstgmygtbstg")
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}