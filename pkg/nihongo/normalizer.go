package nihongo

type Normalizer func(s string) (string, error)

func Normalize(s string, normalizers ...Normalizer) (string, error) {
	var err error
	for i := range normalizers {
		s, err = normalizers[i](s)
		if err != nil {
			return "", err
		}
	}
	return s, nil
}
