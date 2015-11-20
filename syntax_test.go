package util

type SomeType1 struct {
	isInitialized bool
	field01       string
}

func (self *SomeType1) InitFunc() interface{} {
	return func(field01 string) {
		if !self.isInitialized {
			// 初始化
			self.field01 = field01
		}
		self.isInitialized = true
	}
}
