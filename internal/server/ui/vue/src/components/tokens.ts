import {Token} from "@/generated/flow";


export interface token {
  value: Token
  name: string

  img: string
}

export const tokenProps: Record<Token, token> = {
  ETH: {
    name: 'ETH',
    img: '/icons/coins/ETH.webp',
    value: Token.ETH,
  },
  USDCBridged: {
    name: 'USDC.b',
    img: '/icons/coins/USDbC.webp',
    value: Token.USDCBridged,
  },
  USDC: {
    name: 'USDC',
    img: '/icons/coins/USDC.webp',
    value: Token.USDC,
  },
  USDp: {
    name: 'USD+',
    img: '/icons/coins/USD+.webp',
    value: Token.USDp,
  },
  BNB: {
    name: 'BNB',
    img: '/icons/coins/BNB.webp',
    value: Token.BNB,
  },
  BUSD: {
    name: 'BUSD',
    img: '/icons/coins/BUSD.webp',
    value: Token.BUSD,
  },
  LSD: {
    name: 'LSD',
    img: '/icons/coins/LSD.png',
    value: Token.LSD,
  },
  LUSD: {
    name: 'LUSD',
    img: '/icons/coins/LUSD.png',
    value: Token.LUSD,
  },
  STG: {
    name: 'STG',
    img: '/icons/coins/STG.png',
    value: Token.STG,
  },
  veSTG: {
    name: 'veSTG',
    img: '/icons/coins/STG.png',
    value: Token.veSTG,
  },
  MAV: {
    name: 'MAV',
    img: '/icons/coins/MAV.png',
    value: Token.MAV,
  },
  WETH: {
    name: 'WETH',
    img: '/icons/coins/ETH.webp',
    value: Token.WETH,
  },
  AVAX: {
    name: 'AVAX',
    img: '/icons/coins/AVAX.png',
    value: Token.AVAX,
  },
  MATIC: {
    name: 'MATIC',
    img: '/icons/coins/MATIC.png',
    value: Token.MATIC,
  },
  IZI: {
    name: 'IZI',
    img: '/icons/coins/IZI.png',
    value: Token.IZI,
  },
  MUTE: {
    name: 'MUTE',
    img: '/icons/coins/MUTE.png',
    value: Token.MUTE,
  },
  SPACE: {
    name: 'SPACE',
    img: '/icons/coins/SPACE.png',
    value: Token.SPACE,
  },
  USDT: {
    name: 'USDT',
    img: '/icons/coins/USDT.png',
    value: Token.USDT,
  },
  VC: {
    name: 'VC',
    img: '/icons/coins/VC.png',
    value: Token.VC,
  }
}

export const stableCoins = [
  Token.ETH,
  Token.USDC,
  Token.USDCBridged,
  Token.USDT,
  Token.BUSD,
  Token.USDp,
]
