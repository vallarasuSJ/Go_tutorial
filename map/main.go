package main

import "fmt"

func main() {
	m:=make(map[int]string)
	m[0]="a"
	m[1]="b"
	m[2]="c"
	m[3]="d" 



	//comma ok idiom
	v,ok:=m[0]
	if ok{
		fmt.Println(v)
	}
   //comma ok and statement statement idiom
	if v,ok:=m[0];ok{
		fmt.Println(v)
	}
	

	//exercise 1
	map1:=make(map[string][]int)
	map1["kv"]=[]int{1,3,3,4}
	map1["hi"]=[]int{5,4,6,7,7}
	for i,v1:=range map1{
		fmt.Println(i, " ",v1)
		for k,v2:=range v1{
			fmt.Println(k," ",v2)
		}

	}

	slice:=[]string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some",
	"advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever",
	"you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all",
	"the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve",
	"had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually",
	"communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant",
	"a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to",
	"reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious",
	"natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few",
	"veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach",
	"itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so",
	"it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a",
	"politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown",
	"men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have",
	"feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by",
	"some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on",
	"the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least",
	"the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and",
	"marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of",
	"infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i",
	"forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly",
	"repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out",
	"unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i",
	"come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be",
	"founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain",
	"point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from",
	"the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in",
	"uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no",
	"more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.",
	"only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was",
	"exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which",
	"i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of",
	"successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,",
	"some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were",
	"related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes",
	"ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with",
	"that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the",
	"creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic",
	"readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and",
	"which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out",
	"all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust",
	"floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my",
	"interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my",
	"family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western",
	"city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,",
	"and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of",
	"buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s",
	}

	map2:=make(map[string]int)
	for _,v:=range slice{
		map2[v]++
	}
	for k,v:=range map2{
		fmt.Println(k, " ", v)
	}

	
	
}