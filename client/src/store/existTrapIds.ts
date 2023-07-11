import { defineStore } from 'pinia'
import { ref } from 'vue'

import { traPId } from '@/apis/generated'
import useIconUrl from '@/composables/useIconUrl'

export const useExistTrapIds = defineStore('existTrapIds', () => {
  const exists = ref<traPId[]>([])
  const notExists = ref<traPId[]>([])

  const { genIconUrl } = useIconUrl()

  const isExist = (id: traPId) => {
    if (exists.value.includes(id)) {
      return true
    }

    if (notExists.value.includes(id)) {
      return false
    }

    let returnValue = false
    fetch(genIconUrl(id))
      .then((res) => {
        if (res.status == 200) {
          exists.value.push(id)
          returnValue = true
        } else {
          notExists.value.push(id)
          returnValue = false
        }
      })
      .catch(() => {
        notExists.value.push(id)
        returnValue = false
      })

    return returnValue
  }

  return { isExist, exists, notExists }
})
