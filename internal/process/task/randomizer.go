package task

import (
	"math/rand"
	"time"

	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type PermuteArg struct {
	In               []RandomTask
	Limit            map[v1.TaskType]int
	StartToken       *v1.Token
	FinishToken      *v1.Token
	StartFromNetwork v1.Network
	MaxCount         int
	MaxItemLenght    int
	Debug            bool
}

func RandomN[T any](in []T, n int) []T {

	if len(in) == 0 {
		return []T{}
	}

	for len(in) < n {
		in = append(in, in...)
	}

	return in[:n]
}

func AddOptionalTask[T any](cases [][]T, item T, limit int) [][]T {

	for j, _ := range cases {
		for i := 0; i < limit; i++ {
			s := rand.New(rand.NewSource(time.Now().UnixNano()))
			add := s.Intn(2) == 1
			if add {
				cases[j] = InsertRandom(cases[j], item)
			}
		}
	}

	return cases
}

func InsertRandom[T any](arr []T, target T) []T {
	if len(arr) == 0 {
		return []T{target}
	}

	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := s.Intn(len(arr) + 1)

	if len(arr) == index { // nil or empty slice or after last element
		arr = append(arr, target)
		return arr
	}
	arr = append(arr[:index+1], arr[index:]...) // index < len(a)
	arr[index] = target
	return arr

}
func Insert[T any](arr []T, target T, index int) []T {
	if len(arr) == 0 {
		return []T{target}
	}

	if len(arr) == index { // nil or empty slice or after last element
		arr = append(arr, target)
		return arr
	}
	arr = append(arr[:index+1], arr[index:]...) // index < len(a)
	arr[index] = target
	return arr

}

func Permute(arg PermuteArg) ([][]RandomTask, error) {

	out := make([][]RandomTask, 0)

	services := map[v1.TaskType][]RandomTask{}
	for _, el := range arg.In {
		services[el.Type] = append(services[el.Type], el)
	}

	duplicates := map[string]bool{}

	for _, tasks := range services {
		for _, task := range tasks {

			if task.FromToken() != nil && arg.StartToken != nil && task.FromToken().String() != arg.StartToken.String() {
				continue
			}
			limit := decLimit(arg.Limit, task.Type)
			rec(arg, services, limit, []RandomTask{task}, &out, arg.MaxItemLenght, arg.MaxCount, duplicates)
		}
	}

	return out, nil
}

func decLimit(limit map[v1.TaskType]int, target v1.TaskType) map[v1.TaskType]int {
	l := maps.Clone(limit)
	l[target]--
	if l[target] == 0 {
		delete(l, target)
	}
	return l
}

func overLimit(limit map[v1.TaskType]int, target v1.TaskType) bool {
	v, ok := limit[target]
	if !ok {
		return true
	}

	if v == 0 {
		return true
	}

	return false
}

func rec(arg PermuteArg, s map[v1.TaskType][]RandomTask, limit map[v1.TaskType]int, cur []RandomTask, out *[][]RandomTask, maxCurLen, maxTotalLen int, duplicates map[string]bool) {

	if len(*out) >= maxTotalLen {
		return
	}

	if len(cur) == maxCurLen {

		key := ""

		if arg.FinishToken != nil && cur[len(cur)-1].ToToken() != nil {
			if cur[len(cur)-1].ToToken().String() != arg.FinishToken.String() {
				return
			}
		}

		for _, task := range cur {
			key += task.InKey() + task.Type.String() + task.OutKey()
		}

		_, exist := duplicates[key]
		if exist {
			return
		}
		duplicates[key] = true

		*out = append(*out, cur)
		return
	}

	if len(limit) == 0 {
		return
	}

	last := cur[len(cur)-1]

	for taskType, tasks := range s {
		if overLimit(limit, taskType) {
			continue
		}

		for _, task := range tasks {
			if task.InKey() == last.OutKey() {
				nlimit := decLimit(limit, task.Type)
				cloune := slices.Clone(cur)
				cloune = append(cloune, task)
				rec(arg, s, nlimit, cloune, out, maxCurLen, maxTotalLen, duplicates)
			}
		}
	}
}

type RandomTask struct {
	Type v1.TaskType
	Kind v1.TaskKind
	P    any
	//*v1.RPswapItem
	//*v1.RPsimple
}

func (r RandomTask) InKey() string {

	switch r.P.(type) {
	case *v1.RPswapItem:
		p := r.P.(*v1.RPswapItem)
		return p.Network.String() + p.From.String()
	case *v1.RPsimple:
		p := r.P.(*v1.RPsimple)
		return p.Network.String()
	default:
		return ""
	}
}

func (r RandomTask) OutKey() string {

	switch r.P.(type) {
	case *v1.RPswapItem:
		p := r.P.(*v1.RPswapItem)
		return p.Network.String() + p.To.String()
	case *v1.RPsimple:
		p := r.P.(*v1.RPsimple)
		return p.Network.String()
	default:
		return ""
	}
}

func (r RandomTask) Stable() bool {

	switch r.P.(type) {
	case *v1.RPswapItem:
		p := r.P.(*v1.RPswapItem)
		_, fromStable := stableCoinMap[p.GetFrom()]
		_, toStable := stableCoinMap[p.GetTo()]
		return fromStable && toStable
	default:
		return false
	}
}

func (r RandomTask) FromToken() *v1.Token {

	switch r.P.(type) {
	case *v1.RPswapItem:
		p := r.P.(*v1.RPswapItem)
		from := p.From
		return &from
	default:
		return nil
	}
}

func (r RandomTask) ToToken() *v1.Token {

	switch r.P.(type) {
	case *v1.RPswapItem:
		p := r.P.(*v1.RPswapItem)
		from := p.To
		return &from
	default:
		return nil
	}
}

func ExtractItems(f []*v1.RandomTask, amNative, amToken *v1.Amount, optional bool) []RandomTask {

	out := []RandomTask{}

	for _, task := range f {
		if optional != task.Optional {
			continue
		}

		switch task.P.(type) {

		case *v1.RandomTask_Swap:
			p := task.P.(*v1.RandomTask_Swap)
			for _, item := range p.Swap.GetItems() {

				if bozdo.NativeTokenMap[item.Network].String() == item.From.String() {
					item.Am = amNative
				} else {
					item.Am = amToken
				}

				out = append(out, RandomTask{
					Type: task.TaskType,
					Kind: v1.TaskKind_TKSwap,
					P:    item,
				})
			}
		case *v1.RandomTask_Simple:
			p := task.P.(*v1.RandomTask_Simple)
			out = append(out, RandomTask{
				Type: task.TaskType,
				Kind: v1.TaskKind_TKSimple,
				P:    p.Simple,
			})
		}
	}

	return out
}

var stableCoinMap = map[v1.Token]bool{
	v1.Token_USDp:        true,
	v1.Token_USDC:        true,
	v1.Token_USDCBridged: true,
	v1.Token_USDT:        true,
	v1.Token_BUSD:        true,
	v1.Token_ETH:         true,
}

func RandomizeRandomBlock(b *v1.FlowBlock_Rand, ignoreStartToken, ignoreFinishToken bool, minDelay, maxDelay int64) ([][]*v1.Task, error) {

	limitMap := map[v1.TaskType]int{}
	for _, el := range b.Rand.Tasks {
		limitMap[el.TaskType]++
	}

	amNative := &v1.Amount{
		Kind: &v1.Amount_PercRange{
			PercRange: &v1.PercentRange{
				Min: b.Rand.NativeTokenMinPercent,
				Max: b.Rand.NativeTokenMaxPercent,
			},
		},
	}

	amToken := &v1.Amount{
		Kind: &v1.Amount_PercRange{
			PercRange: &v1.PercentRange{
				Min: b.Rand.NonnativeTokenMinPercent,
				Max: b.Rand.NonnativeTokenMaxPercent,
			},
		},
	}

	optionalTasks := ExtractItems(b.Rand.Tasks, amNative, amToken, true)
	tasks := ExtractItems(b.Rand.Tasks, amNative, amToken, false)

	arg := PermuteArg{
		In:               tasks,
		Limit:            limitMap,
		StartFromNetwork: b.Rand.StartNetwork,
		MaxCount:         10_000,
		MaxItemLenght:    int(b.Rand.TaskCount),
		Debug:            false,
	}

	if !ignoreStartToken {
		arg.StartToken = &b.Rand.StartToken
	}

	if !ignoreFinishToken {
		arg.FinishToken = &b.Rand.FinishToken
	}

	cases, err := Permute(arg)
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(cases), func(i, j int) {
		cases[i], cases[j] = cases[j], cases[i]
	})

	for _, randomTask := range optionalTasks {
		cases = AddOptionalTask(cases, randomTask, 1)
	}

	out, err := CastTasks(cases)
	if err != nil {
		return nil, err
	}

	out = InsertDelays(out, minDelay, maxDelay)

	return out, nil
}

func InsertDelays(out [][]*v1.Task, minDelay, maxDelay int64) [][]*v1.Task {
	for j := range out {
		for i := 0; i < len(out[j]); i++ {
			if i%2 != 1 {
				out[j] = Insert(out[j], NewDelayTask(minDelay, maxDelay), i)
			}
		}
	}
	return out
}

func NewDelayTask(min, max int64) *v1.Task {

	return &v1.Task{
		TaskType: v1.TaskType_Delay,
		Task: &v1.Task_DelayTask{
			DelayTask: &v1.DelayTask{
				Duration:       0,
				WaitFor:        nil,
				Random:         true,
				MinRandom:      &min,
				MaxRandom:      &max,
				RandomDuration: nil,
			},
		},
	}
}

func CastTasks(cases [][]RandomTask) ([][]*v1.Task, error) {
	out := [][]*v1.Task{}

	for _, c := range cases {

		tasks := []*v1.Task{}
		for weight, randomTask := range c {
			item, err := castItem(randomTask, int64(weight))
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, item)
		}

		out = append(out, tasks)
	}

	return out, nil
}

func TokenCombo(flows [][]*v1.Task) []*v1.TokenArr {

	m := map[v1.Token]map[v1.Token]bool{}

	for _, flow := range flows {

		var input *TaskInput
		var err error
		start := 0
		for start < len(flows) {
			firstTask := flow[start]

			input, err = taskInput(firstTask)
			if err == nil && !input.NoInput {
				break
			}
			start++
		}

		end := len(flow) - 1

		var output *TaskOutput
		for end > 0 {
			lastTask := flow[end]
			output, err = taskOutput(lastTask)
			if err == nil && !output.NoOutput {
				break
			}

			end--
		}

		_, exist := m[input.Token]
		if !exist {
			m[input.Token] = map[v1.Token]bool{}
		}
		m[input.Token][output.Token] = true
	}

	out := []*v1.TokenArr{}

	for from, toMap := range m {
		tos := []v1.Token{}
		for to, _ := range toMap {
			tos = append(tos, to)
		}

		out = append(out, &v1.TokenArr{
			From: from,
			To:   tos,
		})
	}

	return out
}

func taskInput(task *v1.Task) (*TaskInput, error) {
	spec, exist := SpecMap[task.TaskType]
	if !exist {
		return nil, errors.New("no spec found")
	}

	if spec.Input == nil {
		return nil, errors.New("no input fn found")
	}
	return spec.Input(task)
}

func taskOutput(task *v1.Task) (*TaskOutput, error) {
	spec, exist := SpecMap[task.TaskType]
	if !exist {
		return nil, errors.New("no spec found")
	}
	if spec.Output == nil {
		return nil, errors.New("no output fn found")
	}

	return spec.Output(task)
}

func castItem(in RandomTask, weight int64) (*v1.Task, error) {

	out := &v1.Task{
		Weight:   weight,
		TaskType: in.Type,
		Task:     nil,
	}

	spec := SpecMap[in.Type]
	if spec.CastR != nil {
		return spec.CastR(in)
	}

	return out, errors.New("gg")
}
