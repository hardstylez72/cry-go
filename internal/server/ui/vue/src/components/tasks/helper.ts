import {Amount} from "@/generated/flow";

export const getStatusColor = (status: string) => {
  switch (status) {
    case 'error':
      return 'red'
    case 'completed':
      return 'green'
    case 'not started':
      return 'grey'
    default:
      return 'blue'
  }
}

export const getAmountSend = (am?: Amount): string => {
  if (!am) {
    return 'all coins'
  }

  switch (true) {
    case am.sendAll:
      return 'all coins'
    case am.sendPercent !== undefined:
      return `${String(am.sendPercent)}%`
    case am.sendValue !== undefined:
      return `${String(am.sendAmount)} `
    case am.percRange !== undefined:
      return `${String(am.percRange.min)}-${String(am.percRange.max)}%`
    default:
      return "invalid value"
  }
}

export const onlyInteger = (v: any) => {
  if (!v) {
    return 'Обязательное поле'
  }

  if (Number.isInteger(Number(v))) {
    return true
  }
  return 'Должно быть целое число'
}

export const moreThanZero = (v: any) => {

  if (Number.isNaN(v)) {
    return "Должно быть числом"
  }
  if (String(v) === "0") {
    return "Должно быть больше 0"
  }

  return true
}


export const required = (v: any) => !!v || 'обязательно для заполнения'
