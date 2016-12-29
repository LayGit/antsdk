package user

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayUserInfoShareResponse struct {
  api.AlipayResponse
  UserId                          string                      `json:"user_id"`                            // 用户的userId
  Avatar                          string                      `json:"avatar"`                             // 用户头像
  Phone                           string                      `json:"phone"`                              // 电话号码。
  UserType                        string                      `json:"user_type"`                          // 用户类型（1/2） 1代表公司账户2代表个人账户
  UserStatus                      string                      `json:"user_status"`                        // 用户状态（Q/T/B/W）。 Q代表快速注册用户 T代表已认证用户 B代表被冻结账户 W代表已注册，未激活的账户
  Email                           string                      `json:"email"`                              // 用户支付宝账号绑定的邮箱地址
  UserName                        string                      `json:"user_name"`                          // 若用户是个人用户，则是用户的真实姓名；若是企业用户，则是企业名称。
  Mobile                          string                      `json:"mobile"`                             // 手机号码。
  IsCertified                     string                      `json:"is_certified"`                       // 是否通过实名认证。T是通过 F是没有实名认证
  CertType                        string                      `json:"cert_type"`                          // 0:身份证 1:护照 2:军官证 3:士兵证 4:回乡证 5:临时身份证 6:户口簿 7:警官证 8:台胞证 9:营业执照 10其它证件
  Province                        string                      `json:"province"`                           // 省份名称。
  City                            string                      `json:"city"`                               // 市名称。
  Area                            string                      `json:"area"`                               // 区县名称。
  Address                         string                      `json:"address"`                            // 详细地址。
  Zip                             string                      `json:"zip"`                                // 邮政编码。
  NickName                        string                      `json:"nick_name"`                          // 用户昵称
  IsBalanceFrozen                 string                      `json:"is_balance_frozen"`                  // 余额账户是否被冻结。 T--被冻结；F--未冻结
  IsStudentCertified              string                      `json:"is_student_certified"`               // 是否是学生
  CertNo                          string                      `json:"cert_no"`                            // 证件号码，结合证件类型使用.
  Gender                          string                      `json:"gender"`                             // 性别（F：女性；M：男性）
  PersonBirthday                  string                      `json:"person_birthday"`                    // 个人用户生日。
  Profession                      string                      `json:"profession"`                         // 职业
  PersonCertExpiryDate            string                      `json:"person_cert_expiry_date"`            // 证件有效期（用户类型是个人的时候才有此字段）。
  PersonPictures                  []AlipayUserPicture         `json:"person_pictures"`                    // 个人证件图片（用户类型是个人的时候才有此字段）。
  LicenseNo                       string                      `json:"license_no"`                         // 企业执照号码（用户类型是公司类型时才有此字段）。
  BusinessScope                   string                      `json:"business_scope"`                     // 经营/业务范围（用户类型是公司类型时才有此字段）。
  LicenseExpiryDate               string                      `json:"license_expiry_date"`                // 营业执照有效期，yyyyMMdd或长期，（用户类型是公司类型时才有此字段）。
  OrganizationCode                string                      `json:"organization_code"`                  // 组织机构代码（用户类型是公司类型时才有此字段）。
  FirmPictures                    []AlipayUserPicture         `json:"firm_pictures"`                      // 企业相关证件图片，包含图片地址（目前需要调用alipay.user.certify.image.fetch转换一下）及类型（用户类型是公司类型时才有此字段）。
  FirmType                        string                      `json:"firm_type"`                          // 公司类型，包括（用户类型是公司类型时才有此字段）： CO(公司) INST(事业单位), COMM(社会团体), NGO(民办非企业组织), STATEORGAN(党政国家机关)
  FirmLegalPersonName             string                      `json:"firm_legal_person_name"`             // 企业法人名称（用户类型是公司类型时才有此字段）。
  FirmLegalPersonCertNo           string                      `json:"firm_legal_person_cert_no"`          // 法人证件号码（用户类型是公司类型时才有此字段）。
  FirmLegalPersonCertType         string                      `json:"firm_legal_person_cert_type"`        // 企业法人证件类型, 返回值参考cert_type字段（用户类型是公司类型时才有此字段）。
  FirmLegalPersonCertExpiryDate   string                      `json:"firm_legal_person_cert_expiry_date"` // 企业法人证件有效期（用户类型是公司类型时才有此字段）。
  FirmLegalPersonPictures         []AlipayUserPicture         `json:"firm_legal_person_pictures"`         // 企业法人证件图片（用户类型是公司类型时才有此字段）。
  FirmAgentPersonName             string                      `json:"firm_agent_person_name"`             // 企业代理人姓名（用户类型是公司类型时才有此字段）。
  FirmAgentPersonCertNo           string                      `json:"firm_agent_person_cert_no"`          // 企业代理人证件号码（用户类型是公司类型时才有此字段）。
  FirmAgentPersonCertType         string                      `json:"firm_agent_person_cert_type"`        // 企业代理人证件类型, 返回值参考cert_type字段（用户类型是公司类型时才有此字段）。
  FirmAgentPersonCertExpiryDate   string                      `json:"firm_agent_person_cert_expiry_date"` // 企业代理人证件有效期（用户类型是公司类型时才有此字段）。
  DeliverAddresses                []AlipayUserDeliverAddress  `json:"deliver_addresses"`                  // 收货地址列表
}

type AlipayUserPicture struct {
  PictureURL  string `json:"picture_url"`   // 用于调用alipay.user.certify.image.fetch接口，获取图片资源
  PictureType string `json:"picture_type"`  // 图片类型，包括身份证正反面、营业执照等
}
