import api from './index'

export interface CouponVerifyResponse {
  valid: boolean
  code: string
  discount: number
  message: string
}

export const couponApi = {
  verify(code: string, amount: number) {
    return api.post<CouponVerifyResponse>('/coupons/verify', { code, amount })
  },
}
