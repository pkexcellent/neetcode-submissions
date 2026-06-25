type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    	// we can use a index table to encode the array of string,
	var encoded string
	var contentString string
	var lengthString string
	for _, s := range strs {
		lengthString = lengthString + fmt.Sprintf("%d,", len(s))
		contentString = contentString + s
	}
	encoded = lengthString + "|" + contentString
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
    var parts = strings.SplitN(encoded, "|", 2)
	var metadata, content = parts[0], parts[1]
	var lengths = strings.Split(metadata, ",")
	var decoded []string
	var idxStart = 0
	for _, l := range lengths {
		if l != "" {
			n, err := strconv.Atoi(l)
			if err != nil {
				break
			}
			idxEnd := idxStart + n
			if idxEnd > len(content) {
				return nil
			}
			decoded = append(decoded, content[idxStart:idxEnd])
			idxStart = idxEnd
		}
	}
	return decoded
}

