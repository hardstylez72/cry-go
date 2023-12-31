import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MenuTaskTestNetBridge from "@/components/tasks/BRIDGE/TestNet/Menu.vue";
import TaskTestNetBridgeSwap from "@/components/tasks/BRIDGE/TestNet/Block.vue";

export const TestNetSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: false,
  menu: MenuTaskTestNetBridge,
  component: TaskTestNetBridgeSwap,
  descFn(task) {
    return `[${task.testNetBridgeSwapTask?.minAmount}:${task.testNetBridgeSwapTask?.maxAmount}]`
  },
  bridge: (task: Task) => {
    const p = task.testNetBridgeSwapTask
    return {
      from: Token.ETH, to: Token.ETH, fromNetwork: p?.network, toNetwork: Network.GOERLIETH
    }
  },
  service: {
    name: 'TestNet-bridge',
    img: '/icons/testnetbridge.svg',
    link: 'https://testnetbridge.com/',
    op: 'bridge',

  },
  job: TaskJob.Bridge,
  networks: new Set<Network>([Network.ARBITRUM]),
  airdrops: new Set<Airdrop>([Airdrop.LayerZero]),
  profileType: new Set([ProfileType.EVM])
}
