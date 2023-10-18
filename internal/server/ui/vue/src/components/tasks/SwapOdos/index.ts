import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import OdosSwap from "@/components/tasks/SwapOdos/OdosSwap.vue";
import MenuOdosSwap from "@/components/tasks/SwapOdos/MenuOdosSwap.vue";

export const SwapOdosSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.odosSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: OdosSwap,
  menu: MenuOdosSwap,
  service: {
    name: 'Odos',
    link: 'https://app.odos.xyz',
    img: '/icons/odos.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA, Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
