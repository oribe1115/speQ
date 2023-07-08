const useIconUrl = () => {
  const genIconUrl = (traPId: string) => {
    return 'https://q.trap.jp/api/v3/public/icon/' + traPId
  }

  return { genIconUrl }
}

export default useIconUrl
