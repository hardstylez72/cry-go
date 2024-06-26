import {Network, ProcessStatus} from "@/generated/process";
import dayjs from "dayjs";

import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import {TaskType, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/helper";


dayjs.extend(utc)
dayjs.extend(timezone)

export const GetStatusText = (s: string): string => {
    switch (s) {
        case ProcessStatus.StatusError:
            return "Error"
        case ProcessStatus.StatusRunning:
            return "Running"
        case ProcessStatus.StatusReady:
            return "Ready"
        case ProcessStatus.StatusDone:
            return "Done"
        case ProcessStatus.StatusStop:
            return "Stop"
        case ProcessStatus.StatusRetry:
            return "Retry"
        default:
            return "Unknown"
    }
}

export const GetStatusColor = (s: string): string => {
    switch (s) {
        case ProcessStatus.StatusError:
            return "#FF8A80"
        case ProcessStatus.StatusRunning:
            return "#BBDEFB"
        case ProcessStatus.StatusReady:
            return "#FFFFFF00"
        case ProcessStatus.StatusDone:
            return "#C8E6C9"
        case ProcessStatus.StatusRetry:
            return "#BBDEFB"
        case ProcessStatus.StatusStop:
            return "#BDBDBD"
        default:
            return "#BDBDBD"
    }
}


export class Timer {

    private timer: any = 0
    resolve: any


    fn = () => {

    }

    constructor() {
    }


    cb(fn: () => any) {
        this.fn = fn
    }

    add(seconds: number) {
        if (this.timer) {
            clearTimeout(this.timer)
        }

        this.timer = setTimeout(() => {
            this.fn()
        }, seconds)
    }
}

export const Delay = (seconds: number, cb?: (() => void)) => new Promise((resolve, reject) => {
    const timer = setTimeout(() => {
        if (cb) {
            cb()
        }
        resolve({})
    }, seconds)
})


export const formatTime = (d: Date): string => {
    const tz = dayjs.tz.guess()
    return dayjs(d).utc().tz(tz, false).format('YYYY-MM-DD HH:mm:ss')
}

export const getTime = (addMinutes?: number): string => {
    const tz = dayjs.tz.guess()
    const d = new Date()
    if (addMinutes) {
        const h = addMinutes / 60
        return dayjs(d).utc().add(h, 'h').add(addMinutes, 'm').tz(tz, false).format('HH:mm')
    }
    return dayjs(d).utc().tz(tz, false).format('HH:mm')
}

export const getDate = (addDays?: number): string => {
    const tz = dayjs.tz.guess()
    const d = new Date()
    if (addDays) {
        return dayjs(d).utc().add(addDays, 'd').tz(tz, false).format('YYYY-MM-DD')
    }
    return dayjs(d).utc().tz(tz, false).format('YYYY-MM-DD')
}

export const ts = (date: string, time: string): Date => {
    const tz = dayjs.tz.guess()
    return dayjs(date + " " + time).utc().tz(tz, false).toDate()
}

export const humanDuration = (s?: string | Number): string => {
    let totalSeconds = Number(s)
    const hours = Math.floor(totalSeconds / 3600);
    totalSeconds %= 3600;
    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    let result = ''
    if (hours !== 0) {
        result += ` ${hours} hour `
    }
    if (minutes !== 0) {
        result += ` ${minutes} min `
    }
    if (seconds !== 0) {
        result += ` ${seconds} sec `
    }

    if ((hours === 0) && (minutes === 0) && (seconds === 0)) {
        result = '0 sec'
    }

    return result
}


export interface TaskSettings {
    slippage?: string
    swapRateRatio?: string
    swapUseExchangeRate?: string
}

export const castTaskSettingsMap = (a: Object): Map<TaskType, TaskSettings> => {

    const result = new Map<TaskType, TaskSettings>()

    if (!a) {
        return result
    }

    for (const name of Object.getOwnPropertyNames(a)) {
        //@ts-ignore
        result.set(name, {...a[name]})
    }

    return result
}

export const castTaskSettingsMapObj = (a: Map<TaskType, TaskSettings>): Object => {

    let out = {}
    a.forEach((k, v) => {
        out[v] = k
    })

    return out
}

export const percent = (v: any) => {
    const value = Number(v).valueOf()
    if (value === 0) {
        return 'required'
    }

    if (v > 100) {
        return 'value can not be more than 100'
    }

    if (v < 0) {
        return 'value can not be negative'
    }
    
    return true
}

export const onlyInteger = (v: number) => {
    if (!v) {
        return 'required'
    }

    if (Number.isInteger(Number(v))) {
        return true
    }
    return 'number must be integer'
}

export const ratio = (v: number) => {

    if (Number.isNaN(v)) {
        return 'must be number'
    }

    if (!v) {
        return 'required'
    }

    if (v <= 0) {
        return 'must be > 0'
    }

    if (v > 1) {
        return 'must be < 1'
    }
    return true
}


export const shuffleArray = <T = any>(array: T[]) => {
    for (var i = array.length - 1; i > 0; i--) {
        var j = Math.floor(Math.random() * (i + 1));
        var temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}

export interface SwapPair {
    network: Network
    name: string
    from: Token
    to: Token
}

export const tokenSwapPair = (n: Network, t1: Token, t2: Token): SwapPair => {
    return {
        network: n,
        name: `${t1} -> ${t2}`,
        from: t1,
        to: t2
    }
}

export interface BridgePair {
    name: string
    from: Token
    to: Token

    fromNetwork: Network
    toNetwork: Network
}

export const tokenBridgePair = (n1: Network, n2: Network, t1: Token, t2: Token): BridgePair => {
    return {
        name: `${t1} -> ${t2}`,
        from: t1,
        to: t2,
        fromNetwork: n1,
        toNetwork: n2,
    }
}

export interface network {
    value: Network
    name: string
    img: string
}

export const networkProps: Record<Network, network> = {
    DFK: {
        name: 'DFK',
        img: '/icons/networks/dfk.svg',
        value: Network.DFK,
    },
    Shimmer: {
        name: 'Shimmer',
        img: '/icons/networks/shimmer.svg',
        value: Network.Shimmer,
    },
    Celo: {
        name: 'Celo',
        img: '/icons/networks/celo.svg',
        value: Network.Celo,
    },
    Klaytn: {
        name: 'Klaytn',
        img: '/icons/networks/klaytn.svg',
        value: Network.Klaytn,
    },
    Loot: {
        name: 'Loot',
        img: '/icons/networks/loot.svg',
        value: Network.Loot,
    },
    Fuse: {
        name: 'Fuse',
        img: '/icons/networks/fuse.svg',
        value: Network.Fuse,
    },
    Conflux: {
        name: 'Conflux',
        img: '/icons/networks/cfx.svg',
        value: Network.Conflux,
    },
    Core: {
        name: 'CORE',
        img: '/icons/networks/core.png',
        value: Network.Core,
    },
    Etherium: {
        name: 'Ethereum',
        img: '/icons/networks/ethereum.svg',
        value: Network.Etherium,
    },
    ZKSYNCERA: {
        name: 'ZkSync Era',
        img: '/icons/networks/zksync.svg',
        value: Network.ZKSYNCERA,
    },
    ARBITRUM: {
        name: 'Arbitrum',
        img: '/icons/networks/arbitrum.svg',
        value: Network.ARBITRUM,
    },
    OPTIMISM: {
        name: 'Optimism',
        img: '/icons/networks/optimism.svg',
        value: Network.OPTIMISM,
    },
    AVALANCHE: {
        name: 'Avalanche',
        img: '/icons/networks/avalanche.svg',
        value: Network.AVALANCHE,
    },
    GOERLIETH: {
        name: 'GoerliETH',
        img: '/icons/networks/goerli.svg',
        value: Network.GOERLIETH,
    },
    BinanaceBNB: {
        name: 'BNB',
        img: '/icons/networks/bsc.svg',
        value: Network.BinanaceBNB,
    },
    POLIGON: {
        name: 'Polygon',
        img: '/icons/networks/polygon.svg',
        value: Network.POLIGON,
    },
    ZKSYNCERATESTNET: {
        name: 'ZkSync Era TestNet',
        img: '/icons/networks/zksync.svg',
        value: Network.ZKSYNCERATESTNET,
    },
    ZKSYNCLITE: {
        name: 'ZkSync lite',
        img: '/icons/networks/zksync.svg',
        value: Network.ZKSYNCLITE,
    },
    StarkNet: {
        name: 'StarkNet',
        img: '/icons/starknet.svg',
        value: Network.StarkNet,
    },
    Meter: {
        name: 'Meter',
        img: '/icons/networks/meter.svg',
        value: Network.Meter,
    },
    Tenet: {
        name: 'Tenet',
        img: '/icons/networks/tenet.svg',
        value: Network.Tenet,
    },
    Canto: {
        name: 'Canto',
        img: '/icons/networks/canto.svg',
        value: Network.Canto,
    },
    ArbitrumNova: {
        name: 'Arbitrum Nova',
        img: '/icons/networks/arb-nova.svg',
        value: Network.ArbitrumNova,
    },
    PolygonZKEVM: {
        name: 'Polygon zkEVM',
        img: '/icons/networks/polygon-zkevm.svg',
        value: Network.PolygonZKEVM,
    },
    Fantom: {
        name: 'Fantom',
        img: '/icons/networks/fantom.svg',
        value: Network.Fantom,
    },
    Base: {
        name: 'Base',
        img: '/icons/networks/base.svg',
        value: Network.Base,
    },
    opBNB: {
        name: 'opBNB',
        img: '/icons/networks/opbnb.svg',
        value: Network.opBNB,
    },
    Linea: {
        name: 'Linea',
        img: '/icons/networks/linea.ico',
        value: Network.Linea,
    },
    Zora: {
        name: 'Zora',
        img: '/icons/networks/zora.png',
        value: Network.Zora,
    },
}


export const profileTitle = (item: any): string => {
    if (item.type === ProfileType.EVM) {
        return `${item.num} (${item.type})`
    }
    return `${item.num} (${item.type} ${item.subType})`
}


export const isMobile = () => {
    return window.innerWidth < 1280
}

export const copyToClipboard = async (text) => {
    if (navigator.clipboard) {
        await navigator.clipboard.writeText(text)
        return //codes below wont be executed
    }
    const textArea = document.createElement("textarea")
    textArea.value = text

    document.body.appendChild(textArea)

    textArea.focus()
    textArea.select()

    document.execCommand('copy')

    document.body.removeChild(textArea)
}
