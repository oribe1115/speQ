import { AppClient } from '@/apis/generated'

const apiClient = new AppClient({
  BASE: import.meta.env.PROD ? '/api' : 'http://localhost:4010/api'
})

export default apiClient
