package contents

import "testing"

func TestCountWords(t *testing.T) {
	contents := loadContents()

	topWords := countWords(contents)
	if len(topWords) == 0 {
		t.Error("countWords failed")
	}

	t.Logf("%+v", topWords)
}

func TestRankWords(t *testing.T) {
	contents := loadContents()

	topWords := countWords(contents)
	if len(topWords) == 0 {
		t.Error("countWords failed")
	}

	top10Words := rankWords(topWords)
	if len(topWords) == 0 {
		t.Error("rankWords failed")
	}

	t.Logf("%+v", top10Words)
}

func loadContents() string {
	contents := `The quick brown fox jumped over the lazy dogs The quick brown fox jumped over 
	the lazy dogs The quick brown fox jumped over the lazy dogs The quick brown fox 
	jumped The quick brown fox jumped over the lazy dogs The quick brown fox jumped 
	over the lazy dogs The quick brown fox jumped over the lazy dogs The quick brown 
	fox jumped The quick brown fox jumped over the lazy dogs The quick brown 
	fox jumped over the lazy dogs The quick brown fox jumped over the lazy dogs 
	The quick brown fox jumped A wonderful serenity has taken possession of my 
	entire soul, like these sweet mornings of spring which I enjoy with my whole heart. 
	I am alone, and feel the charm of existence in this spot, which was created for the 
	bliss of souls like mine. I am so happy, my dear friend, so absorbed in the exquisite 
	sense of mere tranquil existence, that I neglect my talents. I should be incapable of 
	drawing a single stroke at the present moment; and yet I feel that I never was a greater 
	artist than now. When, while the lovely valley teems with vapour around me, and the 
	meridian sun strikes the upper surface of the impenetrable foliage of my trees, 
	and but a few stray gleams steal into the inner sanctuary, I throw myself down 
	among the tall grass by the trickling stream; and, as I lie close to the earth, 
	a thousand unknown plants are noticed by me: when I hear the buzz of the little 
	world among the stalks, and grow familiar with the countless indescribable forms 
	of the insects and flies, then I feel the presence of the Almighty, who formed us in his own image`
	return contents
}
