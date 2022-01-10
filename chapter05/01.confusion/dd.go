package main

type downloadFromNetDisk struct {
	secret   DynamicSecret
	filePath string
}

func (dd *downloadFromNetDisk) DownloadFile() {
	if err := dd.logincheck(); err != nil {
		// todo 重新登录
	}
	dd.downloadFromAliYun(dd.filePath)
}

func (dd *downloadFromNetDisk) logincheck() error {
	dd.checkSecret(dd.secret.GetSecret())
	return nil
}

func (dd *downloadFromNetDisk) downloadFromAliYun(file string) {

}

func (dd *downloadFromNetDisk) checkSecret(secret string) {
	// todo 调用阿里云的验证接口去验证密码是否有效
}
