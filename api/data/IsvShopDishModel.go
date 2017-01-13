package data

type IsvShopDishModel struct {
  ShopId                        string  `json:"shop_id"`                            // 外部商户ID/或者门店ID
  Platform                      string  `json:"platform"`                           // 平台类型（二维火，mike)
  OuterDishId                   string  `json:"outer_dish_id"`                      // 菜品id
  Name                          string  `json:"name"`                               // 菜品名称
  Price                         float64 `json:"price"`                              // 菜品价格
  Quantity                      int     `json:"quantity"`                           // 菜品库存
  Pict                          string  `json:"pict"`                               // 多图路径分割
  Content                       string  `json:"content"`                            // 菜品的描述内容
  SortCol                       []int   `json:"sort_col"`                           // 菜品热度值（分数越高表示热度越高）
  Unit                          string  `json:"unit"`                               // 菜品显示的单位（份/斤/杯）
  DishTypeId                    string  `json:"dish_type_id"`                       // 菜品分类ID(商家自定义)
  DishTypeName                  string  `json:"dish_type_name"`                     // 商家定义菜品的分类名称
  SoldCntSevenD                 int     `json:"sold_cnt_seven_d"`                   // shopid最近7天销量（份）
  SoldCntThirtyD                int     `json:"sold_cnt_thirty_d"`                  // shopid最近30天销量（份）
  SoldUsercntThirtyD            int     `json:"sold_usercnt_thirty_d"`              // shopid最近30天消费支付宝用户数
  SoldReusercntThirtyD          int     `json:"sold_reusercnt_thirty_d"`            // shopid最近30天购买2次及以上的支付宝用户数
  MerchantSoldCntSevenD         int     `json:"merchant_sold_cnt_seven_d"`          // merchant最近7天销量（份）
  MerchantSoldCntThirtyD        int     `json:"merchant_sold_cnt_thirty_d"`         // merchant最近30天销量（份）
  MerchantSoldUsercntThirtyD    int     `json:"merchant_sold_usercnt_thirty_d"`     // merchant最近30天消费支付宝用户数
  MerchantSoldReusercntThirtyD  int     `json:"merchant_sold_reusercnt_thirty_d"`   // merchant最近30天购买2次及以上的支付宝用户数
  GoodLevel                     string  `json:"good_level"`                         // 菜品热度等级（0/0.5/1/1.5/2/2.5/3/3.5/4/4.5/5）该字段是对sort_col做离散化。
}
