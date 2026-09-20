package main

import (
	"fmt"
	"testing"
)

func TestFindSecretWord(t *testing.T) {
	testCases := []struct {
		name   string
		secret string
		words  []string
		want   bool
	}{
		{
			name: "example 1",
		},
		{
			name:   "test case 1",
			secret: "ccoyyo",
			words: []string{
				"wichbx",
				"oahwep",
				"tpulot",
				"eqznzs",
				"vvmplb",
				"eywinm",
				"dqefpt",
				"kmjmxr",
				"ihkovg",
				"trbzyb",
				"xqulhc",
				"bcsbfw",
				"rwzslk",
				"abpjhw",
				"mpubps",
				"viyzbc",
				"kodlta",
				"ckfzjh",
				"phuepp",
				"rokoro",
				"nxcwmo",
				"awvqlr",
				"uooeon",
				"hhfuzz",
				"sajxgr",
				"oxgaix",
				"fnugyu",
				"lkxwru",
				"mhtrvb",
				"xxonmg",
				"tqxlbr",
				"euxtzg",
				"tjwvad",
				"uslult",
				"rtjosi",
				"hsygda",
				"vyuica",
				"mbnagm",
				"uinqur",
				"pikenp",
				"szgupv",
				"qpxmsw",
				"vunxdn",
				"jahhfn",
				"kmbeok",
				"biywow",
				"yvgwho",
				"hwzodo",
				"loffxk",
				"xavzqd",
				"vwzpfe",
				"uairjw",
				"itufkt",
				"kaklud",
				"jjinfa",
				"kqbttl",
				"zocgux",
				"ucwjig",
				"meesxb",
				"uysfyc",
				"kdfvtw",
				"vizxrv",
				"rpbdjh",
				"wynohw",
				"lhqxvx",
				"kaadty",
				"dxxwut",
				"vjtskm",
				"yrdswc",
				"byzjxm",
				"jeomdc",
				"saevda",
				"himevi",
				"ydltnu",
				"wrrpoc",
				"khuopg",
				"ooxarg",
				"vcvfry",
				"thaawc",
				"bssybb",
				"ccoyyo",
				"ajcwbj",
				"arwfnl",
				"nafmtm",
				"xoaumd",
				"vbejda",
				"kaefne",
				"swcrkh",
				"reeyhj",
				"vmcwaf",
				"chxitv",
				"qkwjna",
				"vklpkp",
				"xfnayl",
				"ktgmfn",
				"xrmzzm",
				"fgtuki",
				"zcffuv",
				"srxuus",
				"pydgmq",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, word := range tc.words {
				fmt.Printf("%s: ", word)
				rank := 0
				for _, word2 := range tc.words {
					count := countMatchLetters(word, word2)
					if count > 0 {
						rank += count
					}
					fmt.Printf("%s(%d) ", word2, count)
				}
				fmt.Printf(": %d\n", rank)
			}
		})
	}
}
