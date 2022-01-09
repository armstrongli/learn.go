package main

type Assets struct {
	assets []Asset
}

func (a *Assets) DoStartWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Unlock()
			d.Open()
		}
	}
}

func (a *Assets) DoStopWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Close()
			d.Lock()
		}
	}
}
