package user

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayUserUserinfoShareResponse struct {
  api.AlipayResponse
  Avatar                  string                      `json:"avatar"`                     // 用户头像
  UserId                  string                      `json:"user_id"`                    // 用户的userId
  UserTypeValue           string                      `json:"user_type_value"`            // 用户类型（1/2） 1代表公司账户2代表个人账户
  UserStatus              string                      `json:"user_status"`                // 用户状态（Q/T/B/W）。 Q代表快速注册用户 T代表已认证用户 B代表被冻结账户 W代表已注册，未激活的账户
  FirmName                string                      `json:"firm_name"`                  // 公司名称（用户类型是公司类型时公司名称才有此字段）。
  RealName                string                      `json:"real_name"`                  // 用户的真实姓名。
  Email                   string                      `json:"email"`                      // 用户支付宝账号绑定的邮箱地址
  CertNo                  string                      `json:"cert_no"`                    // 证件号码
  Gender                  string                      `json:"gender"`                     // 性别（F：女性；M：男性）
  Phone                   string                      `json:"phone"`                      // 电话号码。
  Mobile                  string                      `json:"mobile"`                     // 手机号码。
  IsCertified             string                      `json:"is_certified"`               // 是否通过实名认证。T是通过 F是没有实名认证
  IsBankAuth              string                      `json:"is_bank_auth"`               // T为是银行卡认证，F为非银行卡认证。
  IsIdAuth                string                      `json:"is_id_auth"`                 // T为是身份证认证，F为非身份证认证。
  IsMobileAuth            string                      `json:"is_mobile_auth"`             // T为是手机认证，F为非手机认证。
  IsLicenceAuth           string                      `json:"is_licence_auth"`            // T为通过营业执照认证，F为没有通过
  CertTypeValue           string                      `json:"cert_type_value"`            // 0:身份证 1:护照 2:军官证 3:士兵证 4:回乡证 5:临时身份证 6:户口簿 7:警官证 8:台胞证 9:营业执照 10其它证件
  DeliverPhone            string                      `json:"deliver_phone"`              // 收货地址的联系人固定电话。
  DeliverMobile           string                      `json:"deliver_mobile"`             // 收货地址的联系人移动电话。
  DeliverFullname         string                      `json:"deliver_fullname"`           // 收货人全称
  Province                string                      `json:"province"`                   // 省份名称。
  City                    string                      `json:"city"`                       // 市名称。
  Area                    string                      `json:"area"`                       // 区县名称。
  Address                 string                      `json:"address"`                    // 详细地址。
  Zip                     string                      `json:"zip"`                        // 邮政编码。
  DeliverProvince         string                      `json:"deliver_province"`           // 收货人所在省份
  DeliverCity             string                      `json:"deliver_city"`               // 收货人所在城市
  DeliverArea             string                      `json:"deliver_area"`               // 收货人所在区县
  DeliverAddressList      []AlipayUserDeliverAddress  `json:"deliver_address_list"`       // 收货人地址列表
  DefaultDeliverAddress   string                      `json:"default_deliver_address"`    // 是否默认收货地址，暂时不返回值
  AddressCode             string                      `json:"address_code"`               // 区域编码，暂时不返回值
  IsStudentCertified      string                      `json:"is_student_certified"`       // 是否是学生
  IsCertifyGradeA         string                      `json:"is_certify_grade_a"`         // T：表示A类实名认证；F：表示非A类实名认证
  AlipayUserId            string                      `json:"alipay_user_id"`             // 支付宝用户ID
  Birthday                string                      `json:"birthday"`                   // 用户生日
  NickName                string                      `json:"nick_name"`                  // 用户昵称
  FamilyName              string                      `json:"family_name"`                // 姓氏，取的是realName中的首个字符，对非中文、中文复姓支持较差。
  ReducedBirthday         string                      `json:"reduced_birthday"`           // 生日的月和日，MMdd格式
  IsBalanceFrozen         string                      `json:"is_balance_frozen"`          // T--被冻结；F--未冻结
  BalanceFreezeType       string                      `json:"balance_freeze_type"`        // 【注意】当is_balance_frozen为“F”时，改字段不会返回. CTU ---- CTU冻结，允许用户开启 ALIBABA ---- ALIBABA冻结，允许用户开启 SERVER ---- 后台冻结，允许用户开启 USER ---- 用户冻结 CTU_N---- CTU冻结，不允许用户开启 ALIBABA_N ---- ALIBABA冻结，不允许用户开启 SERVER_N ---- 后台冻结，不允许用户开启 UNKNOWN ---- 降级、或查询超时
}
