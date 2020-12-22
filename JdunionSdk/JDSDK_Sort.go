package JdunionSdk

type Onestruct struct {
	Key   string
	Value string
}
type Persons []Onestruct

// Len()方法和Swap()方法不用变化
// 获取此 slice 的长度
func (p Persons) Len() int { return len(p) }

// 交换数据
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Persons) Less(i, j int) bool {
	return p[i].Key < p[j].Key
}
