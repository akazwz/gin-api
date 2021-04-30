package utils

/*func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}*/

/*func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExist(v)
		if err != nil {
			return err
		}
		if !exist {
			global.LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.LOG.Error("create directory" + v, zap.Any(" error", err))
			}
		}
	}
	return err
}*/
