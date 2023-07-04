import { ref } from 'vue'

import apiClient from '@/apis'
import { UserInfo } from '@/apis/generated'

const useMe = () => {
  const myInfo = ref<UserInfo>()
  apiClient.user.getMe().then((res) => (myInfo.value = res))

  const isLogin = () => {
    return myInfo.value !== undefined
  }

  const getId = () => {
    return myInfo.value?.traPId ?? ''
  }

  return { isLogin, getId }
}

export default useMe
