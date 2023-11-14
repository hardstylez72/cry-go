import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import TaskDeployStarkNetAccount from "@/components/tasks/OTHER/StarknetDeployAccount/Block.vue";
import DeployStarkNetAccount from "@/components/tasks/OTHER/StarknetDeployAccount/Menu.vue";

export const DeploySpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.deployStarkNetAccountTask
    return ''
  },
  component: TaskDeployStarkNetAccount,
  menu: DeployStarkNetAccount,
  service: {
    name: 'StarkNet acc deploy',
    link: 'https://starkgate.starknet.io/',
    img: '/icons/starknet.svg',
    op: 'deploy account',
  },
  job: TaskJob.Other,
  networks: new Set<Network>([Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet])
}
