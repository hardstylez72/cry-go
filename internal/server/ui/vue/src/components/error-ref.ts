export const errorRef = (text: string): string => {
  let base = "("
  switch (true) {
    case text.search('Withdrawal address is not allowlisted for verification exemption') !== -1:
      base = addReason(base, "address is not in whitelist on okex exchange for withdrawal")
      break
    case text.search('transaction failed') !== -1:
      base = addReason(base, "try to increase gas multiplayer value in /settings")
      break
    case text.search('is higher than max') !== -1:
      base = addReason(base, "try to increase max gas value in /settings")
      break
    case text.search('Too little received') !== -1:
      base = addReason(base, 'заданный slippage ниже рыночного в пуле')
      break
    case text.search('insufficient funds') !== -1: {
      base = addReason(base, 'недостаточно средств на балансе для оплаты газа')
    }
  }

  base += ")"

  if (base == "()") {
    return ""
  }
  base = "Возможные причины: " + base
  return base
}

const addReason = (base, reason: string): string => {
  return base + reason + ", "
}
