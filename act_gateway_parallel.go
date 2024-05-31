package loong

import (
	"context"
)

type parallelGatewayCmd struct {
	gateway
}

func (p *parallelGatewayCmd) Do(ctx context.Context) error {
	p.Exec.GwType = parallel

	return p.join(ctx)
}

func (p *parallelGatewayCmd) Emit(ctx context.Context, emt Emitter) error {
	return p.fork(ctx, emt)
}

func (p *parallelGatewayCmd) join(ctx context.Context) error {
	in := len(p.GetIncomingAssociation())
	if in > 1 && p.ForkID == "" {
		panic("不支持无fork 直接join, 在当前节点之前缺少排他网关")
	}

	if in == 1 && p.ForkID == "" {
		// 新入口
		p.ForkMode = newFork
		return nil
	}

	p.Exec.JoinTag = p.GetId()
	p.Exec.Status = STATUS_FINISH
	if err := p.Store.JoinExec(ctx, &p.Exec); err != nil {
		return err
	}
	if in == 1 && p.ForkID != "" {
		p.ForkMode = forkFork
		return nil
	}

	// in > 1
	total, err := p.Store.LoadForks(ctx, p.Exec.ForkID)
	if err != nil {
		return err
	}

	joined, x, _ := findJoined(total, p.GetId())

	if hasSameOutTag(x) {
		panic("流程图错,相同入口汇聚多次")
	}

	if joined > in {
		panic("join数量大于入口数量,流程错。前方缺少并行网关") // 见鬼！
	}

	if joined == in && in == len(total) {
		// join数量等于入口数量，且每个入口都join了一次，fulljoin
		p.ForkMode = fullJoin
		return nil
	}
	if joined == in && in < len(total) {
		if !hasSameForkTag(x) {
			p.ForkMode = fullJoin
			return nil
		}
		// 已完成join 数量 小于总数，模式b， 部分join
		p.ForkMode = halfJoin
		return nil
	}
	// n<in
	// 多进多出，未完成join
	return nil
}

func (c *parallelGatewayCmd) fork(ctx context.Context, emt Emitter) error {
	out := c.GetOutgoingAssociation()
	outN := len(out)

	c.Exec.ForkTag = c.GetId()

	// fork ... fork 模式 和 newFork 模式
	// 直接fork
	if c.ForkMode == forkFork || c.ForkMode == newFork {
		_, xs := c.Exec.forkOut(out)

		if err := c.Store.ForkExec(ctx, xs); err != nil {
			return err
		}
		return c.EmitExec(ctx, xs, emt)
	}

	if c.ForkMode == fullJoin {
		// 全join, 存在2种情况
		// 1 最顶层的fork - join
		// 2 嵌套的  fork - join
		// 区别为p join ID 是否非空， "" 顶层， 非空为嵌套

		if c.Exec.isTop() {
			// 情况1，已经是顶层了
			// 看出口数量，如果大于1 是fork，等于1是正常结束
			// 注意, join 后可以立即fork
			top := c.Exec.top()
			if outN == 1 {
				// 只有一个出口，正常结束
				// 设置一下outTag
				xs := top.children(out)
				return c.EmitExec(ctx, xs, emt)
			}
			// join 后立刻fork了
			_, xs := top.forkOut(out)
			if err := c.Store.ForkExec(ctx, xs); err != nil {
				return err
			}
			return c.EmitExec(ctx, xs, emt)
		} else {
			// 情况2 嵌套fork - join
			// 此种情况下forkid为parent fork id
			xs := c.Exec.parent().children(out)
			if err := c.Store.ForkExec(ctx, xs); err != nil {
				return err
			}
			return c.EmitExec(ctx, xs, emt)
		}
	}

	if c.ForkMode == halfJoin { // 部分join模式下
		// half 模式下，等同于上一个fork多fork了几条线
		// 所有forkid 都等同于上一个fork
		// 部分join 只允许一个出口的限制似乎可以取消?
		if len(c.GetOutgoingAssociation()) > 1 {
			panic("部分join模式下只允许有一个出口")
		}
		xs := c.Exec.children(out)
		if err := c.Store.ForkExec(ctx, xs); err != nil {
			return err
		}
		return c.EmitExec(ctx, xs, emt)
	}

	panic("未知的for-join模式")
}

func findJoined(tags []Exec, me string) (n int, in []Exec, finished int) {
	for _, x := range tags {
		if x.JoinTag == me {
			n++
			in = append(in, x)
		}
		if x.Status == STATUS_FINISH {
			finished += 1
		}
	}
	return
}

func hasSameOutTag(in []Exec) bool {
	m := make(map[string]int)
	for _, x := range in {
		_, ok := m[x.OutTag]
		if ok {
			return true
		}
		m[x.OutTag] = 1
	}
	return false
}

func hasSameForkTag(in []Exec) bool {
	m := make(map[string]int)
	for _, x := range in {
		_, ok := m[x.ForkTag]
		if ok {
			return true
		}
		m[x.ForkTag] = 1
	}
	return false
}
