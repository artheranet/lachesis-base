package adapters

import (
	"github.com/artheranet/lachesis/abft/dagidx"
	"github.com/artheranet/lachesis/hash"
	"github.com/artheranet/lachesis/inter/idx"
	"github.com/artheranet/lachesis/vecfc"
)

type VectorSeqToDagIndexSeq struct {
	*vecfc.HighestBeforeSeq
}

type BranchSeq struct {
	vecfc.BranchSeq
}

// Seq is a maximum observed e.Seq in the branch
func (b *BranchSeq) Seq() idx.Event {
	return b.BranchSeq.Seq
}

// MinSeq is a minimum observed e.Seq in the branch
func (b *BranchSeq) MinSeq() idx.Event {
	return b.BranchSeq.MinSeq
}

// Get i's position in the byte-encoded vector clock
func (b VectorSeqToDagIndexSeq) Get(i idx.Validator) dagidx.Seq {
	seq := b.HighestBeforeSeq.Get(i)
	return &BranchSeq{seq}
}

type VectorToDagIndexer struct {
	*vecfc.Index
}

func (v *VectorToDagIndexer) GetMergedHighestBefore(id hash.Event) dagidx.HighestBeforeSeq {
	return VectorSeqToDagIndexSeq{v.Index.GetMergedHighestBefore(id)}
}
