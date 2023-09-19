import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, allNetworks, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import Menu from "@/components/tasks/Exchange/WithdrawMenu.vue";
import Block from "@/components/tasks/Exchange/Withdraw.vue";
import MenuOkexDeposit from "@/components/tasks/Exchange/DepositMenu.vue";
import TaskOkexDeposit from "@/components/tasks/Exchange/Deposit.vue";

export const ExchangeWithdrawSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: false,
  menu: Menu,
  component: Block,
  descFn(task) {
    let p = task.withdrawExchangeTask
    return ` (${p?.network} am: ${p?.sendAllCoins ? 'all' : `[${p?.amountMin}:${p?.amountMax}]`} ${p?.token})`
  },
  service: {
    name: 'Вывод',
    img: '',
    link: '',
    op: 'вывод',
  },
  job: TaskJob.Exchange,
  networks: new Set<Network>(allNetworks),
  airdrops: new Set<Airdrop>(Universal),
  profileType: new Set([ProfileType.EVM, ProfileType.StarkNet])
}

export const ExchangeDeposit: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: MenuOkexDeposit,
  component: TaskOkexDeposit,
  descFn(task) {
    let p = task.okexDepositTask
    return ` (${p?.network} ${p?.token})`
  },
  service: {
    name: 'Депозит',
    link: 'https://www.okx.com/',
    img: '',
    op: 'депозит',
  },
  job: TaskJob.Exchange,
  networks: new Set<Network>(allNetworks),
  airdrops: new Set<Airdrop>(Universal),
  profileType: new Set([ProfileType.EVM])
}
