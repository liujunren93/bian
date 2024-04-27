package entitys

type RequestParameter interface {
	SetTimeStamp(timestamp int64)
}

type RequestParam map[string]any

func (r RequestParam) SetTimeStamp(timestamp int64) {
	r["timeStamp"] = timestamp
}
