import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskZkSwap from "@/components/tasks/SWAPS/ZkSwap/Block.vue";
import MenuTaskZkSwap from "@/components/tasks/SWAPS/ZkSwap/Menu.vue";
import TaskMuteIoSwap from "@/components/tasks/SWAPS/muteio/Block.vue";
import MenuTaskMuteioSwap from "@/components/tasks/SWAPS/muteio/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const MuteIOSwapSpec: TaskSpec = {
  deprecated: false,
  component: TaskMuteIoSwap,
  menu: MenuTaskMuteioSwap,
  canBeEstimated: true,
  descFn(task) {
    const p = task.muteioSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.muteioSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  service: {
    name: 'Mute.io',
    img: '/icons/muteio.png',
    link: 'https://app.mute.io/swap',
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
