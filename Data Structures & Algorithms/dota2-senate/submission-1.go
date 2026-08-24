func predictPartyVictory(senate string) string {
	// use a map to record cnt of two parties
	// key is party, value is []int, position list
	// end is a party's value len is 0
	// for loop to check if the other party exist, if yes, remove the next one
	// and re-queue it
	// we don't need to record the banned position
	// because we re-queueu it, also need to record the banned position

	liveMap := make(map[rune][]int)
	n := len(senate)
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			liveMap['R'] = append(liveMap['R'], i)
		} else {
			liveMap['D'] = append(liveMap['D'], i)
		}	
	}
	for len(liveMap['R']) > 0 && len(liveMap['D']) > 0 {
		firstR := liveMap['R'][0]
		firstD := liveMap['D'][0]
		liveMap['R'] = liveMap['R'][1:len(liveMap['R'])]
		liveMap['D'] = liveMap['D'][1:len(liveMap['D'])]
		if firstR < firstD {
			liveMap['R'] = append(liveMap['R'], firstR + n)
		} else {
			liveMap['D'] = append(liveMap['D'], firstD + n)
		}
	}
	if len(liveMap['R']) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}
