import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import OdosSwap from "@/components/tasks/SWAPS/Woofi/Block.vue";
import MenuOdosSwap from "@/components/tasks/SWAPS/Woofi/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const SwapWoofiSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.woofiSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: OdosSwap,
  menu: MenuOdosSwap,
  service: {
    name: 'Woofi',
    link: 'https://fi.woo.org/swap/',
    img: '/icons/woofi.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.Base]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    tokenSwapPair(Network.Base, Token.ETH, Token.USDCBridged),
    tokenSwapPair(Network.Base, Token.USDCBridged, Token.ETH),
  ]
}
