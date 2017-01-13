package marketing

type MergeVerifyConfig struct {
  GeneralMergeSingle bool `json:"general_merge_single"` // 全场和单品叠加标识
  GeneralMergeOthers bool `json:"general_merge_others"` // 可以与任意其他券叠加
}
