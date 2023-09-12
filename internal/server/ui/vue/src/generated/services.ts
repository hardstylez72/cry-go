//@ts-nocheck
// @formatter:off
import axios from 'axios'
import dayjs from "dayjs";
import {mapActions,} from "pinia";
import {useUserStore} from "@/plugins/pinia";

// @ts-ignore
import CryptoJS, {Hex} from "crypto-js";

import {serviceOptions as rso, ProfileService} from './profile'
import {serviceOptions as so1, HelperService} from './helper'
import {serviceOptions as so2, WithdrawerService} from './withdraw'
import {serviceOptions as so3, FlowService} from './flow'
import {serviceOptions as so4, ProcessService} from './process'
import {serviceOptions as so5, SettingsService} from './settings'
import {serviceOptions as so6, Swap1InchService} from './swap1inch'
import {serviceOptions as so7, OrbiterService} from './orbiter'
import {serviceOptions as so8, PublicService} from './public'
import {serviceOptions as so9, IssueService} from './issue'


export const instance = axios.create({
  baseURL: import.meta.env.DEV ? 'http://localhost:8083/' : '/',
  timeout: 60000,
  withCredentials: true,
});


const store = mapActions(useUserStore, ['redirectToGeneralPage'])

instance.interceptors.response.use(res => {
  return res
}, err => {
  console.error(err)

  if (err.response?.status === 401) {
    store.redirectToGeneralPage()
  } else if (err.response?.status === 403) {

  }

  throw new Error(err.response?.data?.message)
})

export const loginGoogleHref = () => {
  return import.meta.env.DEV ? "http://localhost:8083/api/google/login" : '/api/google/login'
}

rso.axios = instance
so1.axios = instance
so2.axios = instance
so3.axios = instance
so4.axios = instance
so5.axios = instance
so6.axios = instance
so7.axios = instance
so8.axios = instance
so9.axios = instance


export const profileService = new ProfileService()
export const helperService = new HelperService()
export const withdrawerService = new WithdrawerService()
export const flowService = new FlowService()
export const processService = new ProcessService()
export const settingsService = new SettingsService()
export const swap1inchService = new Swap1InchService()
export const orbiterService = new OrbiterService()
export const publicService = new PublicService()
export const issueService = new IssueService()





