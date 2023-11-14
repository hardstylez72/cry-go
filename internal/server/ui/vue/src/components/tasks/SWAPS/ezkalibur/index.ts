import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskEzkaliburSwap from "@/components/tasks/SWAPS/ezkalibur/Block.vue";
import MenuTaskEzkaliburSwap from "@/components/tasks/SWAPS/ezkalibur/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const EzkaliburSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.ezkaliburSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: TaskEzkaliburSwap,
  menu: MenuTaskEzkaliburSwap,
  service: {
    name: 'Ezkalibur',
    link: 'https://dapp.ezkalibur.com/',
    img: '/icons/ezkalibur.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),
  ]
}
