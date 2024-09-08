import { createSignal, type Component, createEffect, For } from 'solid-js'
import styles from './App.module.css'

const App: Component = () => {
  const [pokemons, setPokemons] = createSignal<Pokemon[]>([])
  const [page, setPage] = createSignal<number>(0)

  createEffect(() => {
    // 交差观察器（Intersection Observer）文档
    // <https://developer.mozilla.org/zh-CN/docs/Web/API/Intersection_Observer_API>

    // 创建交差观察器
    const intersectionObserver = new IntersectionObserver(
      // 交差时的回调函数
      async (entires) => {
        if (!entires[0].isIntersecting) {
          return
        }

        const data = await queryPokemons(page())
        setPokemons((previous) => [...previous, ...data])
        setPage((previous) => previous + 1)
      },

      // 交差观察器的配置
      { threshold: 1 },
    )

    // 观察被交差对象
    intersectionObserver.observe(document.querySelector('.loading')!)
  })

  return (
    <div id={styles.App}>
      <ul>
        <For each={pokemons()}>
          {(pokemon) => (
            <li>
              <img src={pokemon.picture} alt={pokemon.name} />
              <div>
                <span>{pokemon.id.toString().padStart(4, '0')}</span>
                <span>{pokemon.name}</span>
              </div>
            </li>
          )}
        </For>
      </ul>

      <div class="loading">加载中...</div>
    </div>
  )
}

export default App

interface Pokemon {
  id: string
  name: string
  picture: string
}

/**
 * 请求宝可梦列表 API
 * @param page 页数
 * @returns 宝可梦列表
 */
const queryPokemons = async (page: number): Promise<Pokemon[]> => {
  try {
    const response = await fetch(
      `${import.meta.env.VITE_API_BASE_URL}/pokemon?page=${page}`,
    )

    const data: { results: Pokemon[] } = await response.json()
    return data.results
  } catch (error) {
    console.error(error)
    return []
  }
}
