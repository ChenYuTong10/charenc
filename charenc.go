package charenc

type UnsupportedEncoding string

func (e UnsupportedEncoding) Error() string { return "unsupported encoding" + string(e) }
